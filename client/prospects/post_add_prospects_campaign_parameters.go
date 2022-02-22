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

// NewPostAddProspectsCampaignParams creates a new PostAddProspectsCampaignParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAddProspectsCampaignParams() *PostAddProspectsCampaignParams {
	return &PostAddProspectsCampaignParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAddProspectsCampaignParamsWithTimeout creates a new PostAddProspectsCampaignParams object
// with the ability to set a timeout on a request.
func NewPostAddProspectsCampaignParamsWithTimeout(timeout time.Duration) *PostAddProspectsCampaignParams {
	return &PostAddProspectsCampaignParams{
		timeout: timeout,
	}
}

// NewPostAddProspectsCampaignParamsWithContext creates a new PostAddProspectsCampaignParams object
// with the ability to set a context for a request.
func NewPostAddProspectsCampaignParamsWithContext(ctx context.Context) *PostAddProspectsCampaignParams {
	return &PostAddProspectsCampaignParams{
		Context: ctx,
	}
}

// NewPostAddProspectsCampaignParamsWithHTTPClient creates a new PostAddProspectsCampaignParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAddProspectsCampaignParamsWithHTTPClient(client *http.Client) *PostAddProspectsCampaignParams {
	return &PostAddProspectsCampaignParams{
		HTTPClient: client,
	}
}

/* PostAddProspectsCampaignParams contains all the parameters to send to the API endpoint
   for the post add prospects campaign operation.

   Typically these are written to a http.Request.
*/
type PostAddProspectsCampaignParams struct {

	// Body.
	Body *models.CreateProspect

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post add prospects campaign params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAddProspectsCampaignParams) WithDefaults() *PostAddProspectsCampaignParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post add prospects campaign params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAddProspectsCampaignParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post add prospects campaign params
func (o *PostAddProspectsCampaignParams) WithTimeout(timeout time.Duration) *PostAddProspectsCampaignParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post add prospects campaign params
func (o *PostAddProspectsCampaignParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post add prospects campaign params
func (o *PostAddProspectsCampaignParams) WithContext(ctx context.Context) *PostAddProspectsCampaignParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post add prospects campaign params
func (o *PostAddProspectsCampaignParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post add prospects campaign params
func (o *PostAddProspectsCampaignParams) WithHTTPClient(client *http.Client) *PostAddProspectsCampaignParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post add prospects campaign params
func (o *PostAddProspectsCampaignParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post add prospects campaign params
func (o *PostAddProspectsCampaignParams) WithBody(body *models.CreateProspect) *PostAddProspectsCampaignParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post add prospects campaign params
func (o *PostAddProspectsCampaignParams) SetBody(body *models.CreateProspect) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostAddProspectsCampaignParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
