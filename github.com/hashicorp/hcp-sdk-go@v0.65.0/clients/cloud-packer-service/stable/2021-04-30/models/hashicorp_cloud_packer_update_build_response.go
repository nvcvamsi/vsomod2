// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudPackerUpdateBuildResponse hashicorp cloud packer update build response
//
// swagger:model hashicorp.cloud.packer.UpdateBuildResponse
type HashicorpCloudPackerUpdateBuildResponse struct {

	// Information about the build you updated.
	Build *HashicorpCloudPackerBuild `json:"build,omitempty"`
}

// Validate validates this hashicorp cloud packer update build response
func (m *HashicorpCloudPackerUpdateBuildResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuild(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerUpdateBuildResponse) validateBuild(formats strfmt.Registry) error {
	if swag.IsZero(m.Build) { // not required
		return nil
	}

	if m.Build != nil {
		if err := m.Build.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("build")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("build")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud packer update build response based on the context it is used
func (m *HashicorpCloudPackerUpdateBuildResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBuild(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerUpdateBuildResponse) contextValidateBuild(ctx context.Context, formats strfmt.Registry) error {

	if m.Build != nil {
		if err := m.Build.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("build")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("build")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPackerUpdateBuildResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPackerUpdateBuildResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPackerUpdateBuildResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
