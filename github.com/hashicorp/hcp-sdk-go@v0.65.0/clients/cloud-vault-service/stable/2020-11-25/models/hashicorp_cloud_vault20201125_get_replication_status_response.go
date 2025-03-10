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

// HashicorpCloudVault20201125GetReplicationStatusResponse hashicorp cloud vault 20201125 get replication status response
//
// swagger:model hashicorp.cloud.vault_20201125.GetReplicationStatusResponse
type HashicorpCloudVault20201125GetReplicationStatusResponse struct {

	// connection status
	ConnectionStatus *HashicorpCloudVault20201125GetReplicationStatusResponseConnectionStatus `json:"connection_status,omitempty"`

	// sync progress
	SyncProgress *HashicorpCloudVault20201125GetReplicationStatusResponseSyncProgress `json:"sync_progress,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 get replication status response
func (m *HashicorpCloudVault20201125GetReplicationStatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyncProgress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125GetReplicationStatusResponse) validateConnectionStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectionStatus) { // not required
		return nil
	}

	if m.ConnectionStatus != nil {
		if err := m.ConnectionStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection_status")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125GetReplicationStatusResponse) validateSyncProgress(formats strfmt.Registry) error {
	if swag.IsZero(m.SyncProgress) { // not required
		return nil
	}

	if m.SyncProgress != nil {
		if err := m.SyncProgress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sync_progress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sync_progress")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud vault 20201125 get replication status response based on the context it is used
func (m *HashicorpCloudVault20201125GetReplicationStatusResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnectionStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSyncProgress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125GetReplicationStatusResponse) contextValidateConnectionStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.ConnectionStatus != nil {
		if err := m.ConnectionStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection_status")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125GetReplicationStatusResponse) contextValidateSyncProgress(ctx context.Context, formats strfmt.Registry) error {

	if m.SyncProgress != nil {
		if err := m.SyncProgress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sync_progress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sync_progress")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125GetReplicationStatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125GetReplicationStatusResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125GetReplicationStatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
