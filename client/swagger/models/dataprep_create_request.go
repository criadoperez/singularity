// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DataprepCreateRequest dataprep create request
//
// swagger:model dataprep.CreateRequest
type DataprepCreateRequest struct {

	// Whether to delete the source files after export
	DeleteAfterExport *bool `json:"deleteAfterExport,omitempty"`

	// Maximum size of the CAR files to be created
	MaxSize *string `json:"maxSize,omitempty"`

	// Name of the preparation
	// Required: true
	Name *string `json:"name"`

	// Name of Output storage systems to be used for the output
	OutputStorages []string `json:"outputStorages"`

	// Target piece size of the CAR files used for piece commitment calculation
	PieceSize string `json:"pieceSize,omitempty"`

	// Name of Source storage systems to be used for the source
	SourceStorages []string `json:"sourceStorages"`
}

// Validate validates this dataprep create request
func (m *DataprepCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataprepCreateRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this dataprep create request based on context it is used
func (m *DataprepCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataprepCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataprepCreateRequest) UnmarshalBinary(b []byte) error {
	var res DataprepCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
