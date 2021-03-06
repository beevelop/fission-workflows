// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/apiserver/apiserver.proto

/*
Package apiserver is a generated protocol buffer package.

It is generated from these files:
	pkg/apiserver/apiserver.proto

It has these top-level messages:
	WorkflowIdentifier
	SearchWorkflowResponse
	WorkflowInvocationIdentifier
	WorkflowInvocationList
	Health
*/
package apiserver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import fission_workflows_types "github.com/fission/fission-workflows/pkg/types"
import fission_workflows_version "github.com/fission/fission-workflows/pkg/version"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WorkflowIdentifier struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *WorkflowIdentifier) Reset()                    { *m = WorkflowIdentifier{} }
func (m *WorkflowIdentifier) String() string            { return proto.CompactTextString(m) }
func (*WorkflowIdentifier) ProtoMessage()               {}
func (*WorkflowIdentifier) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WorkflowIdentifier) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SearchWorkflowResponse struct {
	Workflows []string `protobuf:"bytes,1,rep,name=workflows" json:"workflows,omitempty"`
}

func (m *SearchWorkflowResponse) Reset()                    { *m = SearchWorkflowResponse{} }
func (m *SearchWorkflowResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchWorkflowResponse) ProtoMessage()               {}
func (*SearchWorkflowResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SearchWorkflowResponse) GetWorkflows() []string {
	if m != nil {
		return m.Workflows
	}
	return nil
}

type WorkflowInvocationIdentifier struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *WorkflowInvocationIdentifier) Reset()                    { *m = WorkflowInvocationIdentifier{} }
func (m *WorkflowInvocationIdentifier) String() string            { return proto.CompactTextString(m) }
func (*WorkflowInvocationIdentifier) ProtoMessage()               {}
func (*WorkflowInvocationIdentifier) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *WorkflowInvocationIdentifier) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type WorkflowInvocationList struct {
	Invocations []string `protobuf:"bytes,1,rep,name=invocations" json:"invocations,omitempty"`
}

func (m *WorkflowInvocationList) Reset()                    { *m = WorkflowInvocationList{} }
func (m *WorkflowInvocationList) String() string            { return proto.CompactTextString(m) }
func (*WorkflowInvocationList) ProtoMessage()               {}
func (*WorkflowInvocationList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *WorkflowInvocationList) GetInvocations() []string {
	if m != nil {
		return m.Invocations
	}
	return nil
}

type Health struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *Health) Reset()                    { *m = Health{} }
func (m *Health) String() string            { return proto.CompactTextString(m) }
func (*Health) ProtoMessage()               {}
func (*Health) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Health) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*WorkflowIdentifier)(nil), "fission.workflows.apiserver.WorkflowIdentifier")
	proto.RegisterType((*SearchWorkflowResponse)(nil), "fission.workflows.apiserver.SearchWorkflowResponse")
	proto.RegisterType((*WorkflowInvocationIdentifier)(nil), "fission.workflows.apiserver.WorkflowInvocationIdentifier")
	proto.RegisterType((*WorkflowInvocationList)(nil), "fission.workflows.apiserver.WorkflowInvocationList")
	proto.RegisterType((*Health)(nil), "fission.workflows.apiserver.Health")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for WorkflowAPI service

type WorkflowAPIClient interface {
	Create(ctx context.Context, in *fission_workflows_types.WorkflowSpec, opts ...grpc.CallOption) (*WorkflowIdentifier, error)
	List(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*SearchWorkflowResponse, error)
	Get(ctx context.Context, in *WorkflowIdentifier, opts ...grpc.CallOption) (*fission_workflows_types.Workflow, error)
	Delete(ctx context.Context, in *WorkflowIdentifier, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	Validate(ctx context.Context, in *fission_workflows_types.WorkflowSpec, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type workflowAPIClient struct {
	cc *grpc.ClientConn
}

func NewWorkflowAPIClient(cc *grpc.ClientConn) WorkflowAPIClient {
	return &workflowAPIClient{cc}
}

func (c *workflowAPIClient) Create(ctx context.Context, in *fission_workflows_types.WorkflowSpec, opts ...grpc.CallOption) (*WorkflowIdentifier, error) {
	out := new(WorkflowIdentifier)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowAPI/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowAPIClient) List(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*SearchWorkflowResponse, error) {
	out := new(SearchWorkflowResponse)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowAPI/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowAPIClient) Get(ctx context.Context, in *WorkflowIdentifier, opts ...grpc.CallOption) (*fission_workflows_types.Workflow, error) {
	out := new(fission_workflows_types.Workflow)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowAPI/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowAPIClient) Delete(ctx context.Context, in *WorkflowIdentifier, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowAPI/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowAPIClient) Validate(ctx context.Context, in *fission_workflows_types.WorkflowSpec, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowAPI/Validate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WorkflowAPI service

type WorkflowAPIServer interface {
	Create(context.Context, *fission_workflows_types.WorkflowSpec) (*WorkflowIdentifier, error)
	List(context.Context, *google_protobuf1.Empty) (*SearchWorkflowResponse, error)
	Get(context.Context, *WorkflowIdentifier) (*fission_workflows_types.Workflow, error)
	Delete(context.Context, *WorkflowIdentifier) (*google_protobuf1.Empty, error)
	Validate(context.Context, *fission_workflows_types.WorkflowSpec) (*google_protobuf1.Empty, error)
}

func RegisterWorkflowAPIServer(s *grpc.Server, srv WorkflowAPIServer) {
	s.RegisterService(&_WorkflowAPI_serviceDesc, srv)
}

func _WorkflowAPI_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fission_workflows_types.WorkflowSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowAPIServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowAPI/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowAPIServer).Create(ctx, req.(*fission_workflows_types.WorkflowSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowAPI_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowAPIServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowAPI/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowAPIServer).List(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowAPIServer).Get(ctx, req.(*WorkflowIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowAPI_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowAPIServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowAPI/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowAPIServer).Delete(ctx, req.(*WorkflowIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowAPI_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fission_workflows_types.WorkflowSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowAPIServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowAPI/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowAPIServer).Validate(ctx, req.(*fission_workflows_types.WorkflowSpec))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkflowAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fission.workflows.apiserver.WorkflowAPI",
	HandlerType: (*WorkflowAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _WorkflowAPI_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _WorkflowAPI_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _WorkflowAPI_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _WorkflowAPI_Delete_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _WorkflowAPI_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/apiserver/apiserver.proto",
}

// Client API for WorkflowInvocationAPI service

type WorkflowInvocationAPIClient interface {
	// Create a new workflow invocation
	//
	// In case the invocation specification is missing fields or contains invalid fields, a HTTP 400 is returned.
	Invoke(ctx context.Context, in *fission_workflows_types.WorkflowInvocationSpec, opts ...grpc.CallOption) (*WorkflowInvocationIdentifier, error)
	InvokeSync(ctx context.Context, in *fission_workflows_types.WorkflowInvocationSpec, opts ...grpc.CallOption) (*fission_workflows_types.WorkflowInvocation, error)
	// Cancel a workflow invocation
	//
	// This action is irreverisble. A canceled invocation cannot be resumed or restarted.
	// In case that an invocation already is canceled, has failed or has completed, nothing happens.
	// In case that an invocation does not exist a HTTP 404 error status is returned.
	Cancel(ctx context.Context, in *WorkflowInvocationIdentifier, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	List(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*WorkflowInvocationList, error)
	// Get the specification and status of a workflow invocation
	//
	// Get returns three different aspects of the workflow invocation, namely the spec (specification), status and logs.
	// To lighten the request load, consider using a more specific request.
	Get(ctx context.Context, in *WorkflowInvocationIdentifier, opts ...grpc.CallOption) (*fission_workflows_types.WorkflowInvocation, error)
	Validate(ctx context.Context, in *fission_workflows_types.WorkflowInvocationSpec, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type workflowInvocationAPIClient struct {
	cc *grpc.ClientConn
}

func NewWorkflowInvocationAPIClient(cc *grpc.ClientConn) WorkflowInvocationAPIClient {
	return &workflowInvocationAPIClient{cc}
}

func (c *workflowInvocationAPIClient) Invoke(ctx context.Context, in *fission_workflows_types.WorkflowInvocationSpec, opts ...grpc.CallOption) (*WorkflowInvocationIdentifier, error) {
	out := new(WorkflowInvocationIdentifier)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowInvocationAPI/Invoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowInvocationAPIClient) InvokeSync(ctx context.Context, in *fission_workflows_types.WorkflowInvocationSpec, opts ...grpc.CallOption) (*fission_workflows_types.WorkflowInvocation, error) {
	out := new(fission_workflows_types.WorkflowInvocation)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowInvocationAPI/InvokeSync", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowInvocationAPIClient) Cancel(ctx context.Context, in *WorkflowInvocationIdentifier, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowInvocationAPI/Cancel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowInvocationAPIClient) List(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*WorkflowInvocationList, error) {
	out := new(WorkflowInvocationList)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowInvocationAPI/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowInvocationAPIClient) Get(ctx context.Context, in *WorkflowInvocationIdentifier, opts ...grpc.CallOption) (*fission_workflows_types.WorkflowInvocation, error) {
	out := new(fission_workflows_types.WorkflowInvocation)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowInvocationAPI/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowInvocationAPIClient) Validate(ctx context.Context, in *fission_workflows_types.WorkflowInvocationSpec, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.WorkflowInvocationAPI/Validate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WorkflowInvocationAPI service

type WorkflowInvocationAPIServer interface {
	// Create a new workflow invocation
	//
	// In case the invocation specification is missing fields or contains invalid fields, a HTTP 400 is returned.
	Invoke(context.Context, *fission_workflows_types.WorkflowInvocationSpec) (*WorkflowInvocationIdentifier, error)
	InvokeSync(context.Context, *fission_workflows_types.WorkflowInvocationSpec) (*fission_workflows_types.WorkflowInvocation, error)
	// Cancel a workflow invocation
	//
	// This action is irreverisble. A canceled invocation cannot be resumed or restarted.
	// In case that an invocation already is canceled, has failed or has completed, nothing happens.
	// In case that an invocation does not exist a HTTP 404 error status is returned.
	Cancel(context.Context, *WorkflowInvocationIdentifier) (*google_protobuf1.Empty, error)
	List(context.Context, *google_protobuf1.Empty) (*WorkflowInvocationList, error)
	// Get the specification and status of a workflow invocation
	//
	// Get returns three different aspects of the workflow invocation, namely the spec (specification), status and logs.
	// To lighten the request load, consider using a more specific request.
	Get(context.Context, *WorkflowInvocationIdentifier) (*fission_workflows_types.WorkflowInvocation, error)
	Validate(context.Context, *fission_workflows_types.WorkflowInvocationSpec) (*google_protobuf1.Empty, error)
}

func RegisterWorkflowInvocationAPIServer(s *grpc.Server, srv WorkflowInvocationAPIServer) {
	s.RegisterService(&_WorkflowInvocationAPI_serviceDesc, srv)
}

func _WorkflowInvocationAPI_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fission_workflows_types.WorkflowInvocationSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowInvocationAPI/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).Invoke(ctx, req.(*fission_workflows_types.WorkflowInvocationSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowInvocationAPI_InvokeSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fission_workflows_types.WorkflowInvocationSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).InvokeSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowInvocationAPI/InvokeSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).InvokeSync(ctx, req.(*fission_workflows_types.WorkflowInvocationSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowInvocationAPI_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowInvocationIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowInvocationAPI/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).Cancel(ctx, req.(*WorkflowInvocationIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowInvocationAPI_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowInvocationAPI/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).List(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowInvocationAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowInvocationIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowInvocationAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).Get(ctx, req.(*WorkflowInvocationIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowInvocationAPI_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(fission_workflows_types.WorkflowInvocationSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowInvocationAPIServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.WorkflowInvocationAPI/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowInvocationAPIServer).Validate(ctx, req.(*fission_workflows_types.WorkflowInvocationSpec))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkflowInvocationAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fission.workflows.apiserver.WorkflowInvocationAPI",
	HandlerType: (*WorkflowInvocationAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _WorkflowInvocationAPI_Invoke_Handler,
		},
		{
			MethodName: "InvokeSync",
			Handler:    _WorkflowInvocationAPI_InvokeSync_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _WorkflowInvocationAPI_Cancel_Handler,
		},
		{
			MethodName: "List",
			Handler:    _WorkflowInvocationAPI_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _WorkflowInvocationAPI_Get_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _WorkflowInvocationAPI_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/apiserver/apiserver.proto",
}

// Client API for AdminAPI service

type AdminAPIClient interface {
	Status(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*Health, error)
	Version(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*fission_workflows_version.Info, error)
}

type adminAPIClient struct {
	cc *grpc.ClientConn
}

func NewAdminAPIClient(cc *grpc.ClientConn) AdminAPIClient {
	return &adminAPIClient{cc}
}

func (c *adminAPIClient) Status(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*Health, error) {
	out := new(Health)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.AdminAPI/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminAPIClient) Version(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*fission_workflows_version.Info, error) {
	out := new(fission_workflows_version.Info)
	err := grpc.Invoke(ctx, "/fission.workflows.apiserver.AdminAPI/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AdminAPI service

type AdminAPIServer interface {
	Status(context.Context, *google_protobuf1.Empty) (*Health, error)
	Version(context.Context, *google_protobuf1.Empty) (*fission_workflows_version.Info, error)
}

func RegisterAdminAPIServer(s *grpc.Server, srv AdminAPIServer) {
	s.RegisterService(&_AdminAPI_serviceDesc, srv)
}

func _AdminAPI_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAPIServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.AdminAPI/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAPIServer).Status(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminAPI_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAPIServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fission.workflows.apiserver.AdminAPI/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAPIServer).Version(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdminAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fission.workflows.apiserver.AdminAPI",
	HandlerType: (*AdminAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _AdminAPI_Status_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _AdminAPI_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/apiserver/apiserver.proto",
}

func init() { proto.RegisterFile("pkg/apiserver/apiserver.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 648 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xcd, 0x6e, 0xd3, 0x4e,
	0x10, 0xc0, 0xe5, 0xb6, 0xf2, 0xbf, 0x99, 0xe8, 0x5f, 0x95, 0xa1, 0x84, 0x92, 0xb6, 0x6a, 0x58,
	0x40, 0x2a, 0x41, 0x78, 0xa5, 0x56, 0x42, 0x22, 0x07, 0xa4, 0x52, 0x10, 0x54, 0xe2, 0x80, 0x1a,
	0xd4, 0x4a, 0xbd, 0xb9, 0xce, 0x3a, 0x59, 0xd5, 0xf5, 0x1a, 0x7b, 0x93, 0x2a, 0x20, 0x2e, 0x5c,
	0xb8, 0x70, 0xe3, 0xc8, 0x99, 0x47, 0xe0, 0xc0, 0x73, 0xf0, 0x0a, 0x3c, 0x08, 0xca, 0x7a, 0xfd,
	0x51, 0x1c, 0x87, 0x9a, 0x5e, 0xe2, 0x78, 0x76, 0x66, 0x7e, 0xf3, 0xb9, 0x86, 0x8d, 0xe0, 0xb4,
	0x4f, 0xed, 0x80, 0x47, 0x2c, 0x1c, 0xb1, 0x30, 0xfb, 0x67, 0x05, 0xa1, 0x90, 0x02, 0xd7, 0x5c,
	0x1e, 0x45, 0x5c, 0xf8, 0xd6, 0xb9, 0x08, 0x4f, 0x5d, 0x4f, 0x9c, 0x47, 0x56, 0xaa, 0xd2, 0xec,
	0xf4, 0xb9, 0x1c, 0x0c, 0x4f, 0x2c, 0x47, 0x9c, 0x51, 0xad, 0x97, 0x3c, 0x1f, 0xa6, 0xfa, 0x74,
	0x02, 0x90, 0xe3, 0x80, 0x45, 0xf1, 0x6f, 0xec, 0xb8, 0xf9, 0xe4, 0xd2, 0xb6, 0x23, 0x16, 0xaa,
	0x53, 0xfd, 0xd4, 0xf6, 0x6b, 0x7d, 0x21, 0xfa, 0x1e, 0xa3, 0xea, 0xed, 0x64, 0xe8, 0x52, 0x76,
	0x16, 0xc8, 0xb1, 0x3e, 0x5c, 0xd7, 0x87, 0x76, 0xc0, 0xa9, 0xed, 0xfb, 0x42, 0xda, 0x92, 0x0b,
	0x5f, 0xa3, 0xc9, 0x5d, 0xc0, 0x23, 0x4d, 0xd8, 0xef, 0x31, 0x5f, 0x72, 0x97, 0xb3, 0x10, 0x97,
	0x60, 0x8e, 0xf7, 0x56, 0x8d, 0x96, 0xb1, 0x55, 0x3b, 0x98, 0xe3, 0x3d, 0xf2, 0x08, 0x1a, 0x5d,
	0x66, 0x87, 0xce, 0x20, 0xd1, 0x3d, 0x60, 0x51, 0x20, 0xfc, 0x88, 0xe1, 0x3a, 0xd4, 0xd2, 0x08,
	0x57, 0x8d, 0xd6, 0xfc, 0x56, 0xed, 0x20, 0x13, 0x10, 0x0b, 0xd6, 0x53, 0xef, 0xfe, 0x48, 0x38,
	0x0a, 0x3d, 0x83, 0xd3, 0x81, 0x46, 0x51, 0xff, 0x15, 0x8f, 0x24, 0xb6, 0xa0, 0xce, 0x53, 0x49,
	0x42, 0xca, 0x8b, 0x48, 0x0b, 0xcc, 0x97, 0xcc, 0xf6, 0xe4, 0x00, 0x1b, 0x60, 0x46, 0xd2, 0x96,
	0xc3, 0x48, 0x7b, 0xd6, 0x6f, 0xdb, 0xdf, 0x17, 0xa0, 0x9e, 0xb8, 0xdf, 0x7d, 0xbd, 0x8f, 0x23,
	0x30, 0xf7, 0x42, 0x66, 0x4b, 0x86, 0xf7, 0xac, 0x62, 0x6b, 0xe3, 0x06, 0x25, 0xfa, 0xdd, 0x80,
	0x39, 0x4d, 0x6a, 0xcd, 0x98, 0x00, 0xab, 0x58, 0x47, 0xb2, 0xf2, 0xf1, 0xe7, 0xaf, 0x2f, 0x73,
	0x4b, 0xa4, 0x46, 0x13, 0x83, 0x8e, 0xd1, 0x46, 0x17, 0x16, 0x54, 0x4e, 0x0d, 0x2b, 0x6e, 0x8d,
	0x95, 0xf4, 0xcd, 0x7a, 0x3e, 0xe9, 0x5b, 0x73, 0x67, 0x26, 0x66, 0x7a, 0x23, 0xc8, 0x35, 0x85,
	0xaa, 0x63, 0x86, 0xc2, 0xb7, 0x30, 0xff, 0x82, 0x49, 0xac, 0x1a, 0x75, 0xf3, 0xf6, 0x5f, 0xab,
	0x41, 0x1a, 0x8a, 0xb6, 0x8c, 0x4b, 0x29, 0x8d, 0xbe, 0xe7, 0xbd, 0x0f, 0xc8, 0xc1, 0x7c, 0xc6,
	0x3c, 0x26, 0x59, 0x75, 0x6a, 0x49, 0x35, 0x12, 0x54, 0xfb, 0x4f, 0xd4, 0x00, 0x16, 0x0f, 0x6d,
	0x8f, 0xf7, 0x2a, 0xf4, 0xaf, 0x0c, 0xb1, 0xa1, 0x10, 0x37, 0x09, 0x66, 0x88, 0x91, 0x76, 0xdd,
	0x31, 0xda, 0xdb, 0xdf, 0x4c, 0xb8, 0x51, 0x1c, 0xcb, 0xc9, 0x04, 0x7d, 0x36, 0xc0, 0x9c, 0x48,
	0x4e, 0xa7, 0xe7, 0x7b, 0x31, 0x84, 0xcc, 0x54, 0x05, 0xf3, 0xf8, 0x72, 0x05, 0x9a, 0xb2, 0x36,
	0x49, 0x49, 0x48, 0x9d, 0x66, 0x0b, 0x30, 0x19, 0xac, 0xaf, 0x06, 0x40, 0x1c, 0x4e, 0x77, 0xec,
	0x3b, 0xd5, 0x43, 0x7a, 0x50, 0xc1, 0x80, 0x50, 0x15, 0xc4, 0x7d, 0xb2, 0x9c, 0x0b, 0x82, 0x46,
	0x63, 0xdf, 0xe9, 0x18, 0xed, 0x63, 0xc4, 0x82, 0x18, 0x87, 0x60, 0xee, 0xd9, 0xbe, 0xc3, 0x3c,
	0xfc, 0xf7, 0xd4, 0x4b, 0x5b, 0xb8, 0xaa, 0xa2, 0xc1, 0xf6, 0x05, 0xac, 0x9e, 0x93, 0xab, 0x6c,
	0xdb, 0xf4, 0xeb, 0x88, 0x5c, 0x57, 0xb8, 0xff, 0x31, 0xdf, 0x01, 0xfc, 0x64, 0xc4, 0x0b, 0x77,
	0x85, 0xf4, 0x2a, 0x75, 0x40, 0xe7, 0x8c, 0xc5, 0x9c, 0x65, 0x6e, 0x37, 0x2a, 0x4f, 0x41, 0x59,
	0x89, 0x37, 0x15, 0xee, 0x16, 0x59, 0xc9, 0xe3, 0xf2, 0x7b, 0xf2, 0xc3, 0x80, 0xc5, 0xdd, 0xde,
	0x19, 0x57, 0xab, 0x71, 0x04, 0x66, 0x57, 0x5d, 0xbb, 0xa5, 0x85, 0xbf, 0x33, 0xb3, 0x4c, 0xf1,
	0x5d, 0x4e, 0x96, 0x15, 0x14, 0x70, 0x91, 0x0e, 0x94, 0xe0, 0x1d, 0xbe, 0x81, 0xff, 0x0e, 0xe3,
	0xaf, 0x5f, 0xa9, 0xe7, 0xcd, 0x29, 0x9e, 0x93, 0x2f, 0xe6, 0xbe, 0xef, 0x8a, 0x9c, 0x57, 0x2d,
	0x7e, 0x5a, 0x3f, 0xae, 0xa5, 0xec, 0x13, 0x53, 0xf9, 0xdb, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff,
	0x1b, 0x97, 0x68, 0x06, 0x10, 0x08, 0x00, 0x00,
}
