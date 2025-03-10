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

// HashicorpCloudVault20201125IsVaultPluginRegisteredRequest hashicorp cloud vault 20201125 is vault plugin registered request
//
// swagger:model hashicorp.cloud.vault_20201125.IsVaultPluginRegisteredRequest
type HashicorpCloudVault20201125IsVaultPluginRegisteredRequest struct {

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`

	// location
	Location *HashicorpCloudInternalLocationLocation `json:"location,omitempty"`

	// plugin name
	PluginName string `json:"plugin_name,omitempty"`

	// plugin type
	PluginType string `json:"plugin_type,omitempty"`

	// plugin version
	PluginVersion string `json:"plugin_version,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 is vault plugin registered request
func (m *HashicorpCloudVault20201125IsVaultPluginRegisteredRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125IsVaultPluginRegisteredRequest) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud vault 20201125 is vault plugin registered request based on the context it is used
func (m *HashicorpCloudVault20201125IsVaultPluginRegisteredRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125IsVaultPluginRegisteredRequest) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125IsVaultPluginRegisteredRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125IsVaultPluginRegisteredRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125IsVaultPluginRegisteredRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
