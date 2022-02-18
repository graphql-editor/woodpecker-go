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

// Prospect prospect
//
// swagger:model Prospect
type Prospect struct {
	CreateProspect

	// last contacted
	LastContacted string `json:"last_contacted,omitempty"`

	// last replied
	LastReplied string `json:"last_replied,omitempty"`

	// status
	// Enum: [ACTIVE BLACKLIST AUTOREPLIED TO-CHECK TO-REVIEW BOUNCED INVALID REPLIED]
	Status string `json:"status,omitempty"`

	// updated
	Updated string `json:"updated,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Prospect) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CreateProspect
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CreateProspect = aO0

	// AO1
	var dataAO1 struct {
		LastContacted string `json:"last_contacted,omitempty"`

		LastReplied string `json:"last_replied,omitempty"`

		Status string `json:"status,omitempty"`

		Updated string `json:"updated,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.LastContacted = dataAO1.LastContacted

	m.LastReplied = dataAO1.LastReplied

	m.Status = dataAO1.Status

	m.Updated = dataAO1.Updated

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Prospect) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CreateProspect)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		LastContacted string `json:"last_contacted,omitempty"`

		LastReplied string `json:"last_replied,omitempty"`

		Status string `json:"status,omitempty"`

		Updated string `json:"updated,omitempty"`
	}

	dataAO1.LastContacted = m.LastContacted

	dataAO1.LastReplied = m.LastReplied

	dataAO1.Status = m.Status

	dataAO1.Updated = m.Updated

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this prospect
func (m *Prospect) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CreateProspect
	if err := m.CreateProspect.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var prospectTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","BLACKLIST","AUTOREPLIED","TO-CHECK","TO-REVIEW","BOUNCED","INVALID","REPLIED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		prospectTypeStatusPropEnum = append(prospectTypeStatusPropEnum, v)
	}
}

// property enum
func (m *Prospect) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, prospectTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Prospect) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this prospect based on the context it is used
func (m *Prospect) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CreateProspect
	if err := m.CreateProspect.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Prospect) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Prospect) UnmarshalBinary(b []byte) error {
	var res Prospect
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
