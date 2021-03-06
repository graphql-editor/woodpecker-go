// Code generated by go-swagger; DO NOT EDIT.

package campaigns

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
	"github.com/go-openapi/swag"
)

// NewGetCampaignListParams creates a new GetCampaignListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCampaignListParams() *GetCampaignListParams {
	return &GetCampaignListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCampaignListParamsWithTimeout creates a new GetCampaignListParams object
// with the ability to set a timeout on a request.
func NewGetCampaignListParamsWithTimeout(timeout time.Duration) *GetCampaignListParams {
	return &GetCampaignListParams{
		timeout: timeout,
	}
}

// NewGetCampaignListParamsWithContext creates a new GetCampaignListParams object
// with the ability to set a context for a request.
func NewGetCampaignListParamsWithContext(ctx context.Context) *GetCampaignListParams {
	return &GetCampaignListParams{
		Context: ctx,
	}
}

// NewGetCampaignListParamsWithHTTPClient creates a new GetCampaignListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCampaignListParamsWithHTTPClient(client *http.Client) *GetCampaignListParams {
	return &GetCampaignListParams{
		HTTPClient: client,
	}
}

/* GetCampaignListParams contains all the parameters to send to the API endpoint
   for the get campaign list operation.

   Typically these are written to a http.Request.
*/
type GetCampaignListParams struct {

	/* ID.

	   Campaign ID
	*/
	ID []int32

	/* Status.

	   Campaign status
	*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get campaign list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCampaignListParams) WithDefaults() *GetCampaignListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get campaign list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCampaignListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get campaign list params
func (o *GetCampaignListParams) WithTimeout(timeout time.Duration) *GetCampaignListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get campaign list params
func (o *GetCampaignListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get campaign list params
func (o *GetCampaignListParams) WithContext(ctx context.Context) *GetCampaignListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get campaign list params
func (o *GetCampaignListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get campaign list params
func (o *GetCampaignListParams) WithHTTPClient(client *http.Client) *GetCampaignListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get campaign list params
func (o *GetCampaignListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get campaign list params
func (o *GetCampaignListParams) WithID(id []int32) *GetCampaignListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get campaign list params
func (o *GetCampaignListParams) SetID(id []int32) {
	o.ID = id
}

// WithStatus adds the status to the get campaign list params
func (o *GetCampaignListParams) WithStatus(status *string) *GetCampaignListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get campaign list params
func (o *GetCampaignListParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetCampaignListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// binding items for id
		joinedID := o.bindParamID(reg)

		// query array param id
		if err := r.SetQueryParam("id", joinedID...); err != nil {
			return err
		}
	}

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetCampaignList binds the parameter id
func (o *GetCampaignListParams) bindParamID(formats strfmt.Registry) []string {
	iDIR := o.ID

	var iDIC []string
	for _, iDIIR := range iDIR { // explode []int32

		iDIIV := swag.FormatInt32(iDIIR) // int32 as string
		iDIC = append(iDIC, iDIIV)
	}

	// items.CollectionFormat: ""
	iDIS := swag.JoinByFormat(iDIC, "")

	return iDIS
}
