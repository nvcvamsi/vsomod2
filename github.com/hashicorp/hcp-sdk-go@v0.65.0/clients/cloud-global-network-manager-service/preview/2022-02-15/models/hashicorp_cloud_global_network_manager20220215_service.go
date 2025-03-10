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

// HashicorpCloudGlobalNetworkManager20220215Service hashicorp cloud global network manager 20220215 service
//
// swagger:model hashicorp.cloud.global_network_manager_20220215.Service
type HashicorpCloudGlobalNetworkManager20220215Service struct {

	// checks critical
	ChecksCritical int32 `json:"checks_critical,omitempty"`

	// checks passing
	ChecksPassing int32 `json:"checks_passing,omitempty"`

	// checks warning
	ChecksWarning int32 `json:"checks_warning,omitempty"`

	// The unique identifier of the Consul cluster the service is registered with
	ClusterID string `json:"cluster_id,omitempty"`

	// V2 resource id for parent cluster
	ClusterResourceID string `json:"cluster_resource_id,omitempty"`

	// V2 resource name for parent cluster
	ClusterResourceName string `json:"cluster_resource_name,omitempty"`

	// True if the service is “connect native”, meaning that it does not use an Envoy proxy.
	ConnectNative bool `json:"connect_native,omitempty"`

	// True if the service is connected to a gateway.
	ConnectedWithGateway bool `json:"connected_with_gateway,omitempty"`

	// True if the service is connected to an Envoy mesh proxy.
	ConnectedWithProxy bool `json:"connected_with_proxy,omitempty"`

	// Union of all the external sources on the service instances.
	ExternalSources []string `json:"external_sources"`

	// Only applies for `gateway` type services
	GatewayConfig *HashicorpCloudGlobalNetworkManager20220215ServiceGatewayConfig `json:"gateway_config,omitempty"`

	// instance count
	InstanceCount int32 `json:"instance_count,omitempty"`

	// Kind of service
	// Possible values: "", "typical", "connect-proxy", "mesh-gateway", "terminating-gateway", "ingress-gateway"
	// "destination", "api-gateway"
	Kind string `json:"kind,omitempty"`

	// Represents how external service requests are handled with regards to mutual TLS protocol
	// Possible values: "", "strict", "permissive"
	MtlsMode string `json:"mtls_mode,omitempty"`

	// Name of the service.
	Name string `json:"name,omitempty"`

	// Namespace where the service is deployed.
	Namespace string `json:"namespace,omitempty"`

	// Partition where the service is deployed.
	Partition string `json:"partition,omitempty"`

	// A name defined at the admin partition level to help determine if two services are the same.
	SamenessGroup *HashicorpCloudGlobalNetworkManager20220215ServiceSamenessGroup `json:"sameness_group,omitempty"`

	// Each instance of a service can have its own tags. At the service level, we return a set of all tags of all service instances.
	Tags []string `json:"tags"`

	// Determines whether service traffic is forced to pass through the sidecar proxy.
	TransparentProxy bool `json:"transparent_proxy,omitempty"`
}

// Validate validates this hashicorp cloud global network manager 20220215 service
func (m *HashicorpCloudGlobalNetworkManager20220215Service) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGatewayConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSamenessGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215Service) validateGatewayConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.GatewayConfig) { // not required
		return nil
	}

	if m.GatewayConfig != nil {
		if err := m.GatewayConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gateway_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gateway_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215Service) validateSamenessGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.SamenessGroup) { // not required
		return nil
	}

	if m.SamenessGroup != nil {
		if err := m.SamenessGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sameness_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sameness_group")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud global network manager 20220215 service based on the context it is used
func (m *HashicorpCloudGlobalNetworkManager20220215Service) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGatewayConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSamenessGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215Service) contextValidateGatewayConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.GatewayConfig != nil {
		if err := m.GatewayConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gateway_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gateway_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215Service) contextValidateSamenessGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.SamenessGroup != nil {
		if err := m.SamenessGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sameness_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sameness_group")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215Service) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215Service) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudGlobalNetworkManager20220215Service
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
