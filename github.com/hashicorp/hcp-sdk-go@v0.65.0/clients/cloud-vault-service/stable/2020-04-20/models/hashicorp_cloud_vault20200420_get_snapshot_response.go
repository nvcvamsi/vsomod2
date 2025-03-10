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

// HashicorpCloudVault20200420GetSnapshotResponse GetSnapshotResponse is a response for retrieving a snapshot's information.
//
// swagger:model hashicorp.cloud.vault_20200420.GetSnapshotResponse
type HashicorpCloudVault20200420GetSnapshotResponse struct {

	// snapshot is the requested snapshot.
	Snapshot *HashicorpCloudVault20200420Snapshot `json:"snapshot,omitempty"`
}

// Validate validates this hashicorp cloud vault 20200420 get snapshot response
func (m *HashicorpCloudVault20200420GetSnapshotResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20200420GetSnapshotResponse) validateSnapshot(formats strfmt.Registry) error {
	if swag.IsZero(m.Snapshot) { // not required
		return nil
	}

	if m.Snapshot != nil {
		if err := m.Snapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshot")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud vault 20200420 get snapshot response based on the context it is used
func (m *HashicorpCloudVault20200420GetSnapshotResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSnapshot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20200420GetSnapshotResponse) contextValidateSnapshot(ctx context.Context, formats strfmt.Registry) error {

	if m.Snapshot != nil {
		if err := m.Snapshot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshot")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20200420GetSnapshotResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20200420GetSnapshotResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20200420GetSnapshotResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
