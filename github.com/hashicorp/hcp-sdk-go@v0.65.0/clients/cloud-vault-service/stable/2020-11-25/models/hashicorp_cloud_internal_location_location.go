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

// HashicorpCloudInternalLocationLocation Location represents a target for an operation in HCP.
//
// swagger:model hashicorp.cloud.internal.location.Location
type HashicorpCloudInternalLocationLocation struct {

	// organization_id is the id of the organization.
	OrganizationID string `json:"organization_id,omitempty"`

	// project_id is the projects id.
	ProjectID string `json:"project_id,omitempty"`

	// region is the region that the resource is located in. It is
	// optional if the object being referenced is a global object.
	Region *HashicorpCloudInternalLocationRegion `json:"region,omitempty"`
}

// Validate validates this hashicorp cloud internal location location
func (m *HashicorpCloudInternalLocationLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudInternalLocationLocation) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(m.Region) { // not required
		return nil
	}

	if m.Region != nil {
		if err := m.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud internal location location based on the context it is used
func (m *HashicorpCloudInternalLocationLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudInternalLocationLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if m.Region != nil {
		if err := m.Region.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudInternalLocationLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudInternalLocationLocation) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudInternalLocationLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
