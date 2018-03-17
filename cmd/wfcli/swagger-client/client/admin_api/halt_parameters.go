// Code generated by go-swagger; DO NOT EDIT.

package admin_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewHaltParams creates a new HaltParams object
// with the default values initialized.
func NewHaltParams() *HaltParams {

	return &HaltParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHaltParamsWithTimeout creates a new HaltParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHaltParamsWithTimeout(timeout time.Duration) *HaltParams {

	return &HaltParams{

		timeout: timeout,
	}
}

// NewHaltParamsWithContext creates a new HaltParams object
// with the default values initialized, and the ability to set a context for a request
func NewHaltParamsWithContext(ctx context.Context) *HaltParams {

	return &HaltParams{

		Context: ctx,
	}
}

// NewHaltParamsWithHTTPClient creates a new HaltParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHaltParamsWithHTTPClient(client *http.Client) *HaltParams {

	return &HaltParams{
		HTTPClient: client,
	}
}

/*HaltParams contains all the parameters to send to the API endpoint
for the halt operation typically these are written to a http.Request
*/
type HaltParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the halt params
func (o *HaltParams) WithTimeout(timeout time.Duration) *HaltParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the halt params
func (o *HaltParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the halt params
func (o *HaltParams) WithContext(ctx context.Context) *HaltParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the halt params
func (o *HaltParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the halt params
func (o *HaltParams) WithHTTPClient(client *http.Client) *HaltParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the halt params
func (o *HaltParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *HaltParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}