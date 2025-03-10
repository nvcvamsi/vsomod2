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
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudVault20200420FetchAuditLogRequest hashicorp cloud vault 20200420 fetch audit log request
//
// swagger:model hashicorp.cloud.vault_20200420.FetchAuditLogRequest
type HashicorpCloudVault20200420FetchAuditLogRequest struct {

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`

	// interval end
	// Format: date-time
	IntervalEnd strfmt.DateTime `json:"interval_end,omitempty"`

	// interval start
	// Format: date-time
	IntervalStart strfmt.DateTime `json:"interval_start,omitempty"`

	// location
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`
}

// Validate validates this hashicorp cloud vault 20200420 fetch audit log request
func (m *HashicorpCloudVault20200420FetchAuditLogRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntervalEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntervalStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20200420FetchAuditLogRequest) validateIntervalEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.IntervalEnd) { // not required
		return nil
	}

	if err := validate.FormatOf("interval_end", "body", "date-time", m.IntervalEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudVault20200420FetchAuditLogRequest) validateIntervalStart(formats strfmt.Registry) error {
	if swag.IsZero(m.IntervalStart) { // not required
		return nil
	}

	if err := validate.FormatOf("interval_start", "body", "date-time", m.IntervalStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudVault20200420FetchAuditLogRequest) validateLocation(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud vault 20200420 fetch audit log request based on the context it is used
func (m *HashicorpCloudVault20200420FetchAuditLogRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20200420FetchAuditLogRequest) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudVault20200420FetchAuditLogRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20200420FetchAuditLogRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20200420FetchAuditLogRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
