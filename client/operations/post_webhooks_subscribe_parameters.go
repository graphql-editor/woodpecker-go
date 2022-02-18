// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/graphql-editor/woodpecker/models"
)

// NewPostWebhooksSubscribeParams creates a new PostWebhooksSubscribeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWebhooksSubscribeParams() *PostWebhooksSubscribeParams {
	return &PostWebhooksSubscribeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWebhooksSubscribeParamsWithTimeout creates a new PostWebhooksSubscribeParams object
// with the ability to set a timeout on a request.
func NewPostWebhooksSubscribeParamsWithTimeout(timeout time.Duration) *PostWebhooksSubscribeParams {
	return &PostWebhooksSubscribeParams{
		timeout: timeout,
	}
}

// NewPostWebhooksSubscribeParamsWithContext creates a new PostWebhooksSubscribeParams object
// with the ability to set a context for a request.
func NewPostWebhooksSubscribeParamsWithContext(ctx context.Context) *PostWebhooksSubscribeParams {
	return &PostWebhooksSubscribeParams{
		Context: ctx,
	}
}

// NewPostWebhooksSubscribeParamsWithHTTPClient creates a new PostWebhooksSubscribeParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWebhooksSubscribeParamsWithHTTPClient(client *http.Client) *PostWebhooksSubscribeParams {
	return &PostWebhooksSubscribeParams{
		HTTPClient: client,
	}
}

/* PostWebhooksSubscribeParams contains all the parameters to send to the API endpoint
   for the post webhooks subscribe operation.

   Typically these are written to a http.Request.
*/
type PostWebhooksSubscribeParams struct {

	// Body.
	Body *models.WebhookSubscribe

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post webhooks subscribe params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWebhooksSubscribeParams) WithDefaults() *PostWebhooksSubscribeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post webhooks subscribe params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWebhooksSubscribeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post webhooks subscribe params
func (o *PostWebhooksSubscribeParams) WithTimeout(timeout time.Duration) *PostWebhooksSubscribeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post webhooks subscribe params
func (o *PostWebhooksSubscribeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post webhooks subscribe params
func (o *PostWebhooksSubscribeParams) WithContext(ctx context.Context) *PostWebhooksSubscribeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post webhooks subscribe params
func (o *PostWebhooksSubscribeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post webhooks subscribe params
func (o *PostWebhooksSubscribeParams) WithHTTPClient(client *http.Client) *PostWebhooksSubscribeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post webhooks subscribe params
func (o *PostWebhooksSubscribeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post webhooks subscribe params
func (o *PostWebhooksSubscribeParams) WithBody(body *models.WebhookSubscribe) *PostWebhooksSubscribeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post webhooks subscribe params
func (o *PostWebhooksSubscribeParams) SetBody(body *models.WebhookSubscribe) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostWebhooksSubscribeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
