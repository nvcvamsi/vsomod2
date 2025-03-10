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

// HashicorpCloudNetwork20200907CreateTGWAttachmentResponse hashicorp cloud network 20200907 create t g w attachment response
//
// swagger:model hashicorp.cloud.network_20200907.CreateTGWAttachmentResponse
type HashicorpCloudNetwork20200907CreateTGWAttachmentResponse struct {

	// operation is the operation that represents the pending creation of the TGW attachment.
	Operation *cloud.HashicorpCloudOperationOperation `json:"operation,omitempty"`

	// tgw_attachment is the TGW attachment that was just created.
	TgwAttachment *HashicorpCloudNetwork20200907TGWAttachment `json:"tgw_attachment,omitempty"`
}

// Validate validates this hashicorp cloud network 20200907 create t g w attachment response
func (m *HashicorpCloudNetwork20200907CreateTGWAttachmentResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTgwAttachment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907CreateTGWAttachmentResponse) validateOperation(formats strfmt.Registry) error {
	if swag.IsZero(m.Operation) { // not required
		return nil
	}

	if m.Operation != nil {
		if err := m.Operation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907CreateTGWAttachmentResponse) validateTgwAttachment(formats strfmt.Registry) error {
	if swag.IsZero(m.TgwAttachment) { // not required
		return nil
	}

	if m.TgwAttachment != nil {
		if err := m.TgwAttachment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tgw_attachment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tgw_attachment")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud network 20200907 create t g w attachment response based on the context it is used
func (m *HashicorpCloudNetwork20200907CreateTGWAttachmentResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOperation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTgwAttachment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudNetwork20200907CreateTGWAttachmentResponse) contextValidateOperation(ctx context.Context, formats strfmt.Registry) error {

	if m.Operation != nil {
		if err := m.Operation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudNetwork20200907CreateTGWAttachmentResponse) contextValidateTgwAttachment(ctx context.Context, formats strfmt.Registry) error {

	if m.TgwAttachment != nil {
		if err := m.TgwAttachment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tgw_attachment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tgw_attachment")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907CreateTGWAttachmentResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudNetwork20200907CreateTGWAttachmentResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudNetwork20200907CreateTGWAttachmentResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
