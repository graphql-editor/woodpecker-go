// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateGlobalProspect create global prospect
//
// swagger:model CreateGlobalProspect
type CreateGlobalProspect struct {

	// prospects
	Prospects []*CreateGlobalProspectProspectsItems0 `json:"prospects"`

	// update
	Update string `json:"update,omitempty"`
}

// Validate validates this create global prospect
func (m *CreateGlobalProspect) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProspects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateGlobalProspect) validateProspects(formats strfmt.Registry) error {
	if swag.IsZero(m.Prospects) { // not required
		return nil
	}

	for i := 0; i < len(m.Prospects); i++ {
		if swag.IsZero(m.Prospects[i]) { // not required
			continue
		}

		if m.Prospects[i] != nil {
			if err := m.Prospects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prospects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("prospects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create global prospect based on the context it is used
func (m *CreateGlobalProspect) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProspects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateGlobalProspect) contextValidateProspects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Prospects); i++ {

		if m.Prospects[i] != nil {
			if err := m.Prospects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prospects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("prospects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateGlobalProspect) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateGlobalProspect) UnmarshalBinary(b []byte) error {
	var res CreateGlobalProspect
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CreateGlobalProspectProspectsItems0 create global prospect prospects items0
//
// swagger:model CreateGlobalProspectProspectsItems0
type CreateGlobalProspectProspectsItems0 struct {

	// address
	Address string `json:"address,omitempty"`

	// city
	City string `json:"city,omitempty"`

	// company
	Company string `json:"company,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// first name
	FirstName string `json:"first_name,omitempty"`

	// id
	ID float64 `json:"id,omitempty"`

	// last name
	LastName string `json:"last_name,omitempty"`

	// phone
	Phone string `json:"phone,omitempty"`

	// snipet1
	Snipet1 string `json:"snipet1,omitempty"`

	// snipet2
	Snipet2 string `json:"snipet2,omitempty"`

	// snipet3
	Snipet3 string `json:"snipet3,omitempty"`

	// snipet4
	Snipet4 string `json:"snipet4,omitempty"`

	// tags
	Tags string `json:"tags,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// website
	Website string `json:"website,omitempty"`
}

// Validate validates this create global prospect prospects items0
func (m *CreateGlobalProspectProspectsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create global prospect prospects items0 based on context it is used
func (m *CreateGlobalProspectProspectsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateGlobalProspectProspectsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateGlobalProspectProspectsItems0) UnmarshalBinary(b []byte) error {
	var res CreateGlobalProspectProspectsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
