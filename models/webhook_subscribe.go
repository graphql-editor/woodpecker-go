// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WebhookSubscribe webhook subscribe
//
// swagger:model WebhookSubscribe
type WebhookSubscribe struct {

	// event
	// Enum: [prospect_replied link_clicked email_opened prospect_bounced prospect_invalid prospect_interested prospect_maybe_later prospect_not_interested prospect_autoreplied followup_after_autoreply]
	Event string `json:"event,omitempty"`

	// target url
	TargetURL string `json:"target_url,omitempty"`
}

// Validate validates this webhook subscribe
func (m *WebhookSubscribe) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var webhookSubscribeTypeEventPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["prospect_replied","link_clicked","email_opened","prospect_bounced","prospect_invalid","prospect_interested","prospect_maybe_later","prospect_not_interested","prospect_autoreplied","followup_after_autoreply"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		webhookSubscribeTypeEventPropEnum = append(webhookSubscribeTypeEventPropEnum, v)
	}
}

const (

	// WebhookSubscribeEventProspectReplied captures enum value "prospect_replied"
	WebhookSubscribeEventProspectReplied string = "prospect_replied"

	// WebhookSubscribeEventLinkClicked captures enum value "link_clicked"
	WebhookSubscribeEventLinkClicked string = "link_clicked"

	// WebhookSubscribeEventEmailOpened captures enum value "email_opened"
	WebhookSubscribeEventEmailOpened string = "email_opened"

	// WebhookSubscribeEventProspectBounced captures enum value "prospect_bounced"
	WebhookSubscribeEventProspectBounced string = "prospect_bounced"

	// WebhookSubscribeEventProspectInvalid captures enum value "prospect_invalid"
	WebhookSubscribeEventProspectInvalid string = "prospect_invalid"

	// WebhookSubscribeEventProspectInterested captures enum value "prospect_interested"
	WebhookSubscribeEventProspectInterested string = "prospect_interested"

	// WebhookSubscribeEventProspectMaybeLater captures enum value "prospect_maybe_later"
	WebhookSubscribeEventProspectMaybeLater string = "prospect_maybe_later"

	// WebhookSubscribeEventProspectNotInterested captures enum value "prospect_not_interested"
	WebhookSubscribeEventProspectNotInterested string = "prospect_not_interested"

	// WebhookSubscribeEventProspectAutoreplied captures enum value "prospect_autoreplied"
	WebhookSubscribeEventProspectAutoreplied string = "prospect_autoreplied"

	// WebhookSubscribeEventFollowupAfterAutoreply captures enum value "followup_after_autoreply"
	WebhookSubscribeEventFollowupAfterAutoreply string = "followup_after_autoreply"
)

// prop value enum
func (m *WebhookSubscribe) validateEventEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, webhookSubscribeTypeEventPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WebhookSubscribe) validateEvent(formats strfmt.Registry) error {
	if swag.IsZero(m.Event) { // not required
		return nil
	}

	// value enum
	if err := m.validateEventEnum("event", "body", m.Event); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this webhook subscribe based on context it is used
func (m *WebhookSubscribe) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebhookSubscribe) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebhookSubscribe) UnmarshalBinary(b []byte) error {
	var res WebhookSubscribe
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
