// Code generated by go-swagger; DO NOT EDIT.

package prospects

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

	"github.com/graphql-editor/woodpecker-go/models"
)

// NewPostAddProspectsListParams creates a new PostAddProspectsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAddProspectsListParams() *PostAddProspectsListParams {
	return &PostAddProspectsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAddProspectsListParamsWithTimeout creates a new PostAddProspectsListParams object
// with the ability to set a timeout on a request.
func NewPostAddProspectsListParamsWithTimeout(timeout time.Duration) *PostAddProspectsListParams {
	return &PostAddProspectsListParams{
		timeout: timeout,
	}
}

// NewPostAddProspectsListParamsWithContext creates a new PostAddProspectsListParams object
// with the ability to set a context for a request.
func NewPostAddProspectsListParamsWithContext(ctx context.Context) *PostAddProspectsListParams {
	return &PostAddProspectsListParams{
		Context: ctx,
	}
}

// NewPostAddProspectsListParamsWithHTTPClient creates a new PostAddProspectsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAddProspectsListParamsWithHTTPClient(client *http.Client) *PostAddProspectsListParams {
	return &PostAddProspectsListParams{
		HTTPClient: client,
	}
}

/* PostAddProspectsListParams contains all the parameters to send to the API endpoint
   for the post add prospects list operation.

   Typically these are written to a http.Request.
*/
type PostAddProspectsListParams struct {

	// Body.
	Body *models.CreateGlobalProspect

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post add prospects list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAddProspectsListParams) WithDefaults() *PostAddProspectsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post add prospects list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAddProspectsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post add prospects list params
func (o *PostAddProspectsListParams) WithTimeout(timeout time.Duration) *PostAddProspectsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post add prospects list params
func (o *PostAddProspectsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post add prospects list params
func (o *PostAddProspectsListParams) WithContext(ctx context.Context) *PostAddProspectsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post add prospects list params
func (o *PostAddProspectsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post add prospects list params
func (o *PostAddProspectsListParams) WithHTTPClient(client *http.Client) *PostAddProspectsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post add prospects list params
func (o *PostAddProspectsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post add prospects list params
func (o *PostAddProspectsListParams) WithBody(body *models.CreateGlobalProspect) *PostAddProspectsListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post add prospects list params
func (o *PostAddProspectsListParams) SetBody(body *models.CreateGlobalProspect) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostAddProspectsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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