// Code generated by go-swagger; DO NOT EDIT.

package service_principals_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
)

// NewServicePrincipalsServiceExchangeWorkloadIdentityTokenParams creates a new ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServicePrincipalsServiceExchangeWorkloadIdentityTokenParams() *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams {
	return &ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServicePrincipalsServiceExchangeWorkloadIdentityTokenParamsWithTimeout creates a new ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams object
// with the ability to set a timeout on a request.
func NewServicePrincipalsServiceExchangeWorkloadIdentityTokenParamsWithTimeout(timeout time.Duration) *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams {
	return &ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams{
		timeout: timeout,
	}
}

// NewServicePrincipalsServiceExchangeWorkloadIdentityTokenParamsWithContext creates a new ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams object
// with the ability to set a context for a request.
func NewServicePrincipalsServiceExchangeWorkloadIdentityTokenParamsWithContext(ctx context.Context) *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams {
	return &ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams{
		Context: ctx,
	}
}

// NewServicePrincipalsServiceExchangeWorkloadIdentityTokenParamsWithHTTPClient creates a new ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewServicePrincipalsServiceExchangeWorkloadIdentityTokenParamsWithHTTPClient(client *http.Client) *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams {
	return &ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams{
		HTTPClient: client,
	}
}

/*
ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams contains all the parameters to send to the API endpoint

	for the service principals service exchange workload identity token operation.

	Typically these are written to a http.Request.
*/
type ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams struct {

	/* Body.

	     token is the external identity token. This represents the identity of a
	workload and for token exchange to occur must be valid given the provider's
	configuration and conditional_access statatement.
	*/
	Body *models.HashicorpCloudIamExchangeWorkloadIdentityTokenRequestToken

	/* ResourceName.

	     resource_name is the resource_name of the workload identity provider to
	exchange against.
	*/
	ResourceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service principals service exchange workload identity token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams) WithDefaults() *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service principals service exchange workload identity token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service principals service exchange workload identity token params
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams) WithTimeout(timeout time.Duration) *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service principals service exchange workload identity token params
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service principals service exchange workload identity token params
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams) WithContext(ctx context.Context) *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service principals service exchange workload identity token params
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service principals service exchange workload identity token params
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams) WithHTTPClient(client *http.Client) *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service principals service exchange workload identity token params
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the service principals service exchange workload identity token params
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams) WithBody(body *models.HashicorpCloudIamExchangeWorkloadIdentityTokenRequestToken) *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the service principals service exchange workload identity token params
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams) SetBody(body *models.HashicorpCloudIamExchangeWorkloadIdentityTokenRequestToken) {
	o.Body = body
}

// WithResourceName adds the resourceName to the service principals service exchange workload identity token params
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams) WithResourceName(resourceName string) *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams {
	o.SetResourceName(resourceName)
	return o
}

// SetResourceName adds the resourceName to the service principals service exchange workload identity token params
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams) SetResourceName(resourceName string) {
	o.ResourceName = resourceName
}

// WriteToRequest writes these params to a swagger request
func (o *ServicePrincipalsServiceExchangeWorkloadIdentityTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param resource_name
	if err := r.SetPathParam("resource_name", o.ResourceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
