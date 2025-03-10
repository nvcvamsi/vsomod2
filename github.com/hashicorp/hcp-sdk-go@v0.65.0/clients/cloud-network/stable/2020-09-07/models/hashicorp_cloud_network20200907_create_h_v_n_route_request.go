// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudNetwork20200907CreateHVNRouteRequest hashicorp cloud network 20200907 create h v n route request
//
// swagger:model hashicorp.cloud.network_20200907.CreateHVNRouteRequest
type HashicorpCloudNetwork20200907CreateHVNRouteRequest struct {

	// destination is a destination CIDR for this HVN Route
	Destination string `json:"destination,omitempty"`

	// hvn is the HVN where a route being created
	Hvn *cloud.HashicorpCloudLocationLink `json:"hvn,omitempty"`

	// id is a user generated id for the route
	ID string `json:"id,omitempty"`

	// target is a target for this HVN Route
	Target *HashicorpCloudNetwork20200907HVNRouteTarget `json:"target,omitempty"`
}

// Validate validates this hashicorp cloud network 20200907 create h v n route request
func (m *HashicorpCloudNetwork20200907CreateHVNRouteRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHvn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907CreateHVNRouteRequest) validateHvn(formats strfmt.Registry) error {
	if swag.IsZero(m.Hvn) { // not required
		return nil
	}

	if m.Hvn != nil {
		if err := m.Hvn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hvn")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hvn")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907CreateHVNRouteRequest) validateTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud network 20200907 create h v n route request based on the context it is used
func (m *HashicorpCloudNetwork20200907CreateHVNRouteRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHvn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907CreateHVNRouteRequest) contextValidateHvn(ctx context.Context, formats strfmt.Registry) error {

	if m.Hvn != nil {
		if err := m.Hvn.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hvn")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hvn")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907CreateHVNRouteRequest) contextValidateTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.Target != nil {
		if err := m.Target.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907CreateHVNRouteRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907CreateHVNRouteRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudNetwork20200907CreateHVNRouteRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
