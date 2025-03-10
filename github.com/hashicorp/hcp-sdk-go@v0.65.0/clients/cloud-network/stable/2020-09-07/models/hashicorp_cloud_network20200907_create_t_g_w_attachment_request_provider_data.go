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

// HashicorpCloudNetwork20200907CreateTGWAttachmentRequestProviderData CreateTGWAttachmentRequestProviderData is the provider specific data to create a TGW attachment. It is to be distinguished from TGWAttachment.ProviderData despite having overlapping fields.
//
// swagger:model hashicorp.cloud.network_20200907.CreateTGWAttachmentRequestProviderData
type HashicorpCloudNetwork20200907CreateTGWAttachmentRequestProviderData struct {

	// aws data
	AwsData *HashicorpCloudNetwork20200907AWSCreateRequestTGWData `json:"aws_data,omitempty"`
}

// Validate validates this hashicorp cloud network 20200907 create t g w attachment request provider data
func (m *HashicorpCloudNetwork20200907CreateTGWAttachmentRequestProviderData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907CreateTGWAttachmentRequestProviderData) validateAwsData(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsData) { // not required
		return nil
	}

	if m.AwsData != nil {
		if err := m.AwsData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws_data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud network 20200907 create t g w attachment request provider data based on the context it is used
func (m *HashicorpCloudNetwork20200907CreateTGWAttachmentRequestProviderData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907CreateTGWAttachmentRequestProviderData) contextValidateAwsData(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsData != nil {
		if err := m.AwsData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws_data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907CreateTGWAttachmentRequestProviderData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907CreateTGWAttachmentRequestProviderData) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudNetwork20200907CreateTGWAttachmentRequestProviderData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
