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

// NewCreateKoofrOtherStorageParams creates a new CreateKoofrOtherStorageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateKoofrOtherStorageParams() *CreateKoofrOtherStorageParams {
	return &CreateKoofrOtherStorageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateKoofrOtherStorageParamsWithTimeout creates a new CreateKoofrOtherStorageParams object
// with the ability to set a timeout on a request.
func NewCreateKoofrOtherStorageParamsWithTimeout(timeout time.Duration) *CreateKoofrOtherStorageParams {
	return &CreateKoofrOtherStorageParams{
		timeout: timeout,
	}
}

// NewCreateKoofrOtherStorageParamsWithContext creates a new CreateKoofrOtherStorageParams object
// with the ability to set a context for a request.
func NewCreateKoofrOtherStorageParamsWithContext(ctx context.Context) *CreateKoofrOtherStorageParams {
	return &CreateKoofrOtherStorageParams{
		Context: ctx,
	}
}

// NewCreateKoofrOtherStorageParamsWithHTTPClient creates a new CreateKoofrOtherStorageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateKoofrOtherStorageParamsWithHTTPClient(client *http.Client) *CreateKoofrOtherStorageParams {
	return &CreateKoofrOtherStorageParams{
		HTTPClient: client,
	}
}

/*
CreateKoofrOtherStorageParams contains all the parameters to send to the API endpoint

	for the create koofr other storage operation.

	Typically these are written to a http.Request.
*/
type CreateKoofrOtherStorageParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateKoofrOtherStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create koofr other storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateKoofrOtherStorageParams) WithDefaults() *CreateKoofrOtherStorageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create koofr other storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateKoofrOtherStorageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create koofr other storage params
func (o *CreateKoofrOtherStorageParams) WithTimeout(timeout time.Duration) *CreateKoofrOtherStorageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create koofr other storage params
func (o *CreateKoofrOtherStorageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create koofr other storage params
func (o *CreateKoofrOtherStorageParams) WithContext(ctx context.Context) *CreateKoofrOtherStorageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create koofr other storage params
func (o *CreateKoofrOtherStorageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create koofr other storage params
func (o *CreateKoofrOtherStorageParams) WithHTTPClient(client *http.Client) *CreateKoofrOtherStorageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create koofr other storage params
func (o *CreateKoofrOtherStorageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create koofr other storage params
func (o *CreateKoofrOtherStorageParams) WithRequest(request *models.StorageCreateKoofrOtherStorageRequest) *CreateKoofrOtherStorageParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create koofr other storage params
func (o *CreateKoofrOtherStorageParams) SetRequest(request *models.StorageCreateKoofrOtherStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateKoofrOtherStorageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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