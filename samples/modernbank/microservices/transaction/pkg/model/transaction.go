// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Transaction transaction
// swagger:model Transaction
type Transaction struct {
	NewTransaction

	// id
	// Required: true
	ID *string `json:"id"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Transaction) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 NewTransaction
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.NewTransaction = aO0

	// AO1
	var dataAO1 struct {
		ID *string `json:"id"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ID = dataAO1.ID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Transaction) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.NewTransaction)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		ID *string `json:"id"`
	}

	dataAO1.ID = m.ID

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this transaction
func (m *Transaction) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with NewTransaction
	if err := m.NewTransaction.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Transaction) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Transaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Transaction) UnmarshalBinary(b []byte) error {
	var res Transaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
