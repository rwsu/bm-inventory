// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Cidr cidr
//
// swagger:model cidr
type Cidr struct {

	// ip address
	IPAddress string `json:"ip-address,omitempty"`

	// mask
	Mask int64 `json:"mask,omitempty"`
}

// Validate validates this cidr
func (m *Cidr) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Cidr) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cidr) UnmarshalBinary(b []byte) error {
	var res Cidr
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
