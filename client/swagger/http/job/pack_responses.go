// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// PackReader is a Reader for the Pack structure.
type PackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPackBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPackInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /job/{id}/pack] Pack", response, response.Code())
	}
}

// NewPackOK creates a PackOK with default headers values
func NewPackOK() *PackOK {
	return &PackOK{}
}

/*
PackOK describes a response with status code 200, with default header values.

OK
*/
type PackOK struct {
	Payload *models.ModelCar
}

// IsSuccess returns true when this pack o k response has a 2xx status code
func (o *PackOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pack o k response has a 3xx status code
func (o *PackOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pack o k response has a 4xx status code
func (o *PackOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pack o k response has a 5xx status code
func (o *PackOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pack o k response a status code equal to that given
func (o *PackOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pack o k response
func (o *PackOK) Code() int {
	return 200
}

func (o *PackOK) Error() string {
	return fmt.Sprintf("[POST /job/{id}/pack][%d] packOK  %+v", 200, o.Payload)
}

func (o *PackOK) String() string {
	return fmt.Sprintf("[POST /job/{id}/pack][%d] packOK  %+v", 200, o.Payload)
}

func (o *PackOK) GetPayload() *models.ModelCar {
	return o.Payload
}

func (o *PackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelCar)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackBadRequest creates a PackBadRequest with default headers values
func NewPackBadRequest() *PackBadRequest {
	return &PackBadRequest{}
}

/*
PackBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PackBadRequest struct {
	Payload string
}

// IsSuccess returns true when this pack bad request response has a 2xx status code
func (o *PackBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pack bad request response has a 3xx status code
func (o *PackBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pack bad request response has a 4xx status code
func (o *PackBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pack bad request response has a 5xx status code
func (o *PackBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pack bad request response a status code equal to that given
func (o *PackBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pack bad request response
func (o *PackBadRequest) Code() int {
	return 400
}

func (o *PackBadRequest) Error() string {
	return fmt.Sprintf("[POST /job/{id}/pack][%d] packBadRequest  %+v", 400, o.Payload)
}

func (o *PackBadRequest) String() string {
	return fmt.Sprintf("[POST /job/{id}/pack][%d] packBadRequest  %+v", 400, o.Payload)
}

func (o *PackBadRequest) GetPayload() string {
	return o.Payload
}

func (o *PackBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackInternalServerError creates a PackInternalServerError with default headers values
func NewPackInternalServerError() *PackInternalServerError {
	return &PackInternalServerError{}
}

/*
PackInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PackInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this pack internal server error response has a 2xx status code
func (o *PackInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pack internal server error response has a 3xx status code
func (o *PackInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pack internal server error response has a 4xx status code
func (o *PackInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pack internal server error response has a 5xx status code
func (o *PackInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pack internal server error response a status code equal to that given
func (o *PackInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pack internal server error response
func (o *PackInternalServerError) Code() int {
	return 500
}

func (o *PackInternalServerError) Error() string {
	return fmt.Sprintf("[POST /job/{id}/pack][%d] packInternalServerError  %+v", 500, o.Payload)
}

func (o *PackInternalServerError) String() string {
	return fmt.Sprintf("[POST /job/{id}/pack][%d] packInternalServerError  %+v", 500, o.Payload)
}

func (o *PackInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *PackInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}