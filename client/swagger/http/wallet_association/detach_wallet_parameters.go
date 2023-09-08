// Code generated by go-swagger; DO NOT EDIT.

package wallet_association

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
)

// NewDetachWalletParams creates a new DetachWalletParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDetachWalletParams() *DetachWalletParams {
	return &DetachWalletParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDetachWalletParamsWithTimeout creates a new DetachWalletParams object
// with the ability to set a timeout on a request.
func NewDetachWalletParamsWithTimeout(timeout time.Duration) *DetachWalletParams {
	return &DetachWalletParams{
		timeout: timeout,
	}
}

// NewDetachWalletParamsWithContext creates a new DetachWalletParams object
// with the ability to set a context for a request.
func NewDetachWalletParamsWithContext(ctx context.Context) *DetachWalletParams {
	return &DetachWalletParams{
		Context: ctx,
	}
}

// NewDetachWalletParamsWithHTTPClient creates a new DetachWalletParams object
// with the ability to set a custom HTTPClient for a request.
func NewDetachWalletParamsWithHTTPClient(client *http.Client) *DetachWalletParams {
	return &DetachWalletParams{
		HTTPClient: client,
	}
}

/*
DetachWalletParams contains all the parameters to send to the API endpoint

	for the detach wallet operation.

	Typically these are written to a http.Request.
*/
type DetachWalletParams struct {

	/* ID.

	   Preparation ID or name
	*/
	ID string

	/* Wallet.

	   Wallet Address
	*/
	Wallet string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the detach wallet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DetachWalletParams) WithDefaults() *DetachWalletParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the detach wallet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DetachWalletParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the detach wallet params
func (o *DetachWalletParams) WithTimeout(timeout time.Duration) *DetachWalletParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the detach wallet params
func (o *DetachWalletParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the detach wallet params
func (o *DetachWalletParams) WithContext(ctx context.Context) *DetachWalletParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the detach wallet params
func (o *DetachWalletParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the detach wallet params
func (o *DetachWalletParams) WithHTTPClient(client *http.Client) *DetachWalletParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the detach wallet params
func (o *DetachWalletParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the detach wallet params
func (o *DetachWalletParams) WithID(id string) *DetachWalletParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the detach wallet params
func (o *DetachWalletParams) SetID(id string) {
	o.ID = id
}

// WithWallet adds the wallet to the detach wallet params
func (o *DetachWalletParams) WithWallet(wallet string) *DetachWalletParams {
	o.SetWallet(wallet)
	return o
}

// SetWallet adds the wallet to the detach wallet params
func (o *DetachWalletParams) SetWallet(wallet string) {
	o.Wallet = wallet
}

// WriteToRequest writes these params to a swagger request
func (o *DetachWalletParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param wallet
	if err := r.SetPathParam("wallet", o.Wallet); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}