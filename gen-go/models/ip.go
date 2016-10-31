package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// IP IP
// swagger:model IP
type IP struct {

	// lat
	// Required: true
	Lat *float64 `json:"lat"`

	// lon
	// Required: true
	Lon *float64 `json:"lon"`
}

// Validate validates this IP
func (m *IP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLat(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLon(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IP) validateLat(formats strfmt.Registry) error {

	if err := validate.Required("lat", "body", m.Lat); err != nil {
		return err
	}

	return nil
}

func (m *IP) validateLon(formats strfmt.Registry) error {

	if err := validate.Required("lon", "body", m.Lon); err != nil {
		return err
	}

	return nil
}
