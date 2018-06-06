package workflow

import (
	"context"
	"fmt"
	"reflect"

	"github.com/fission/fission-workflows/pkg/api"
	"github.com/fission/fission-workflows/pkg/controller"
	"github.com/fission/fission-workflows/pkg/fes"
	"github.com/fission/fission-workflows/pkg/types"
	"github.com/fission/fission-workflows/pkg/types/aggregates"
	"github.com/fission/fission-workflows/pkg/util/gopool"
	"github.com/fission/fission-workflows/pkg/util/labels"
	"github.com/fission/fission-workflows/pkg/util/pubsub"
	"github.com/sirupsen/logrus"
)

const (
	NotificationBuffer    = 100
	evalQueueSize         = 50
	maxParallelExecutions = 100
)

var wfLog = logrus.WithField("component", "controller.wf")

// WorkflowController is the controller concerned with the lifecycle of workflows. It handles responsibilities, such as
// parsing of workflows.
type Controller struct {
	wfCache    fes.CacheReader
	api        *api.Workflow
	sub        *pubsub.Subscription
	cancelFn   context.CancelFunc
	evalQueue  chan string
	evalStates controller.EvalStore
	evalPolicy controller.Rule
}

func NewController(wfCache fes.CacheReader, wfAPI *api.Workflow) *Controller {
	ctr := &Controller{
		wfCache:    wfCache,
		api:        wfAPI,
		evalQueue:  make(chan string, evalQueueSize),
		evalStates: controller.EvalStore{},
	}
	ctr.evalPolicy = defaultPolicy(ctr)
	return ctr
}

func (c *Controller) Init(sctx context.Context) error {
	ctx, cancelFn := context.WithCancel(sctx)
	c.cancelFn = cancelFn

	// Subscribe to invocation creations and task events.
	selector := labels.In(fes.PubSubLabelAggregateType, "workflow")
	if invokePub, ok := c.wfCache.(pubsub.Publisher); ok {
		c.sub = invokePub.Subscribe(pubsub.SubscriptionOptions{
			Buffer:   NotificationBuffer,
			Selector: selector,
		})

		// Workflow Notification listener
		go func(ctx context.Context) {
			for {
				select {
				case notification := <-c.sub.Ch:
					c.handleMsg(notification)
				case <-ctx.Done():
					wfLog.WithField("ctx.err", ctx.Err()).Debug("Notification listener closed.")
					return
				}
			}
		}(ctx)
	}

	// process evaluation queue
	pool := gopool.New(maxParallelExecutions)
	go func(ctx context.Context) {
		for {
			select {
			case eval := <-c.evalQueue:
				pool.Submit(ctx, func() {
					c.Evaluate(eval)
				})
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	return nil
}

func (c *Controller) handleMsg(msg pubsub.Msg) error {
	wfLog.WithField("labels", msg.Labels()).Debug("Handling invocation notification.")
	switch n := msg.(type) {
	case *fes.Notification:
		c.Notify(n)
	default:
		wfLog.WithField("notification", n).Warn("Ignoring unknown notification type")
	}
	return nil
}

func (c *Controller) Tick(tick uint64) error {
	// Assume that all workflows are in evalStates
	//now := time.Now()
	// TODO short loop: eval cache
	// TODO longer loop: cache
	//for _, a := range c.wfCache.List() {
	//	if locked {
	//		continue
	//	}
	//
	//	wfEntity, err := c.wfCache.GetAggregate(a)
	//	if err != nil {
	//		return fmt.Errorf("failed to retrieve: %v", err)
	//	}
	//
	//	wf, ok := wfEntity.(*aggregates.Workflow)
	//	if !ok {
	//		wfLog.WithField("wfEntity", wfEntity).WithField("type", reflect.TypeOf(wfEntity)).
	//			Error("Unexpected type in wfCache")
	//		panic(fmt.Sprintf("unexpected type '%v' in wfCache", reflect.TypeOf(wfEntity)))
	//	}
	//
	//	c.submitEval(wf.id())
	//}
	return nil
}

func (c *Controller) Notify(msg *fes.Notification) error {
	wf, ok := msg.Payload.(*aggregates.Workflow)
	if !ok {
		return fmt.Errorf("received notification of invalid type '%s'. Expected '*aggregates.Workflow'", reflect.TypeOf(msg.Payload))
	}

	c.submitEval(wf.ID())
	return nil
}

func (c *Controller) Evaluate(workflowID string) {
	// Fetch and attempt to claim the evaluation
	evalState := c.evalStates.LoadOrStore(workflowID)
	select {
	case <-evalState.Lock():
		defer evalState.Free()
	default:
		// TODO provide option to wait for a lock
		wfLog.Debugf("Failed to obtain access to workflow %s", workflowID)
		return
	}
	wfLog.Debugf("evaluating workflow %s", workflowID)

	// Fetch the workflow relevant to the invocation
	wf := aggregates.NewWorkflow(workflowID)
	err := c.wfCache.Get(wf)
	// TODO move to rule
	if err != nil && wf.Workflow == nil {
		logrus.Errorf("controller failed to get workflow '%s': %v", workflowID, err)
		return
	}

	// Evaluate invocation
	record := controller.NewEvalRecord() // TODO implement rulepath + cause

	ec := NewEvalContext(evalState, wf.Workflow)

	action := c.evalPolicy.Eval(ec)
	record.Action = action

	// Execute action
	err = action.Apply()
	if err != nil {
		wfLog.Errorf("Action '%T' failed: %v", action, err)
		record.Error = err
	}

	// Record this evaluation
	evalState.Record(record)
}

func (c *Controller) Close() error {
	wfLog.Info("Closing workflow controller...")
	if invokePub, ok := c.wfCache.(pubsub.Publisher); ok {
		err := invokePub.Unsubscribe(c.sub)
		if err != nil {
			return err
		}
	}

	c.cancelFn()
	return nil
}

func (c *Controller) submitEval(ids ...string) bool {
	for _, id := range ids {
		select {
		case c.evalQueue <- id:
			return true
			// ok
		default:
			wfLog.Warnf("Eval queue is full; dropping eval task for '%v'", id)
			return false
		}
	}
	return true
}

func defaultPolicy(ctr *Controller) controller.Rule {
	return &controller.RuleEvalUntilAction{
		Rules: []controller.Rule{
			&RuleSkipIfReady{},
			&RuleRemoveIfDeleted{
				evalCache: ctr.evalStates,
			},
			&RuleEnsureParsed{
				WfAPI: ctr.api,
			},
		},
	}
}

//
// Workflow-specific actions
//

type ActionParseWorkflow struct {
	WfAPI *api.Workflow
	Wf    *types.Workflow
}

func (ac *ActionParseWorkflow) Apply() error {
	_, err := ac.WfAPI.Parse(ac.Wf)
	return err
}
