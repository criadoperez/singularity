// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// CreateStorjExistingStorageReader is a Reader for the CreateStorjExistingStorage structure.
type CreateStorjExistingStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateStorjExistingStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateStorjExistingStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateStorjExistingStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateStorjExistingStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/storj/existing] CreateStorjExistingStorage", response, response.Code())
	}
}

// NewCreateStorjExistingStorageOK creates a CreateStorjExistingStorageOK with default headers values
func NewCreateStorjExistingStorageOK() *CreateStorjExistingStorageOK {
	return &CreateStorjExistingStorageOK{}
}

/*
CreateStorjExistingStorageOK describes a response with status code 200, with default header values.

OK
*/
type CreateStorjExistingStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create storj existing storage o k response has a 2xx status code
func (o *CreateStorjExistingStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create storj existing storage o k response has a 3xx status code
func (o *CreateStorjExistingStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create storj existing storage o k response has a 4xx status code
func (o *CreateStorjExistingStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create storj existing storage o k response has a 5xx status code
func (o *CreateStorjExistingStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create storj existing storage o k response a status code equal to that given
func (o *CreateStorjExistingStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create storj existing storage o k response
func (o *CreateStorjExistingStorageOK) Code() int {
	return 200
}

func (o *CreateStorjExistingStorageOK) Error() string {
	return fmt.Sprintf("[POST /storage/storj/existing][%d] createStorjExistingStorageOK  %+v", 200, o.Payload)
}

func (o *CreateStorjExistingStorageOK) String() string {
	return fmt.Sprintf("[POST /storage/storj/existing][%d] createStorjExistingStorageOK  %+v", 200, o.Payload)
}

func (o *CreateStorjExistingStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreateStorjExistingStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStorjExistingStorageBadRequest creates a CreateStorjExistingStorageBadRequest with default headers values
func NewCreateStorjExistingStorageBadRequest() *CreateStorjExistingStorageBadRequest {
	return &CreateStorjExistingStorageBadRequest{}
}

/*
CreateStorjExistingStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateStorjExistingStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create storj existing storage bad request response has a 2xx status code
func (o *CreateStorjExistingStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create storj existing storage bad request response has a 3xx status code
func (o *CreateStorjExistingStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create storj existing storage bad request response has a 4xx status code
func (o *CreateStorjExistingStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create storj existing storage bad request response has a 5xx status code
func (o *CreateStorjExistingStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create storj existing storage bad request response a status code equal to that given
func (o *CreateStorjExistingStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create storj existing storage bad request response
func (o *CreateStorjExistingStorageBadRequest) Code() int {
	return 400
}

func (o *CreateStorjExistingStorageBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/storj/existing][%d] createStorjExistingStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateStorjExistingStorageBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/storj/existing][%d] createStorjExistingStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateStorjExistingStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateStorjExistingStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStorjExistingStorageInternalServerError creates a CreateStorjExistingStorageInternalServerError with default headers values
func NewCreateStorjExistingStorageInternalServerError() *CreateStorjExistingStorageInternalServerError {
	return &CreateStorjExistingStorageInternalServerError{}
}

/*
CreateStorjExistingStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateStorjExistingStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create storj existing storage internal server error response has a 2xx status code
func (o *CreateStorjExistingStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create storj existing storage internal server error response has a 3xx status code
func (o *CreateStorjExistingStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create storj existing storage internal server error response has a 4xx status code
func (o *CreateStorjExistingStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create storj existing storage internal server error response has a 5xx status code
func (o *CreateStorjExistingStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create storj existing storage internal server error response a status code equal to that given
func (o *CreateStorjExistingStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create storj existing storage internal server error response
func (o *CreateStorjExistingStorageInternalServerError) Code() int {
	return 500
}

func (o *CreateStorjExistingStorageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/storj/existing][%d] createStorjExistingStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateStorjExistingStorageInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/storj/existing][%d] createStorjExistingStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateStorjExistingStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateStorjExistingStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}