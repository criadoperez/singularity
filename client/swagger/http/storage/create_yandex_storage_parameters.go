// Code generated by go-swagger; DO NOT EDIT.

package storage

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

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// NewCreateYandexStorageParams creates a new CreateYandexStorageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateYandexStorageParams() *CreateYandexStorageParams {
	return &CreateYandexStorageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateYandexStorageParamsWithTimeout creates a new CreateYandexStorageParams object
// with the ability to set a timeout on a request.
func NewCreateYandexStorageParamsWithTimeout(timeout time.Duration) *CreateYandexStorageParams {
	return &CreateYandexStorageParams{
		timeout: timeout,
	}
}

// NewCreateYandexStorageParamsWithContext creates a new CreateYandexStorageParams object
// with the ability to set a context for a request.
func NewCreateYandexStorageParamsWithContext(ctx context.Context) *CreateYandexStorageParams {
	return &CreateYandexStorageParams{
		Context: ctx,
	}
}

// NewCreateYandexStorageParamsWithHTTPClient creates a new CreateYandexStorageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateYandexStorageParamsWithHTTPClient(client *http.Client) *CreateYandexStorageParams {
	return &CreateYandexStorageParams{
		HTTPClient: client,
	}
}

/*
CreateYandexStorageParams contains all the parameters to send to the API endpoint

	for the create yandex storage operation.

	Typically these are written to a http.Request.
*/
type CreateYandexStorageParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateYandexStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create yandex storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateYandexStorageParams) WithDefaults() *CreateYandexStorageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create yandex storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateYandexStorageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create yandex storage params
func (o *CreateYandexStorageParams) WithTimeout(timeout time.Duration) *CreateYandexStorageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create yandex storage params
func (o *CreateYandexStorageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create yandex storage params
func (o *CreateYandexStorageParams) WithContext(ctx context.Context) *CreateYandexStorageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create yandex storage params
func (o *CreateYandexStorageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create yandex storage params
func (o *CreateYandexStorageParams) WithHTTPClient(client *http.Client) *CreateYandexStorageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create yandex storage params
func (o *CreateYandexStorageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create yandex storage params
func (o *CreateYandexStorageParams) WithRequest(request *models.StorageCreateYandexStorageRequest) *CreateYandexStorageParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create yandex storage params
func (o *CreateYandexStorageParams) SetRequest(request *models.StorageCreateYandexStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateYandexStorageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}