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

// CreateS3LiaraStorageReader is a Reader for the CreateS3LiaraStorage structure.
type CreateS3LiaraStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateS3LiaraStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateS3LiaraStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateS3LiaraStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateS3LiaraStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/s3/liara] CreateS3LiaraStorage", response, response.Code())
	}
}

// NewCreateS3LiaraStorageOK creates a CreateS3LiaraStorageOK with default headers values
func NewCreateS3LiaraStorageOK() *CreateS3LiaraStorageOK {
	return &CreateS3LiaraStorageOK{}
}

/*
CreateS3LiaraStorageOK describes a response with status code 200, with default header values.

OK
*/
type CreateS3LiaraStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create s3 liara storage o k response has a 2xx status code
func (o *CreateS3LiaraStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create s3 liara storage o k response has a 3xx status code
func (o *CreateS3LiaraStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 liara storage o k response has a 4xx status code
func (o *CreateS3LiaraStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create s3 liara storage o k response has a 5xx status code
func (o *CreateS3LiaraStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create s3 liara storage o k response a status code equal to that given
func (o *CreateS3LiaraStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create s3 liara storage o k response
func (o *CreateS3LiaraStorageOK) Code() int {
	return 200
}

func (o *CreateS3LiaraStorageOK) Error() string {
	return fmt.Sprintf("[POST /storage/s3/liara][%d] createS3LiaraStorageOK  %+v", 200, o.Payload)
}

func (o *CreateS3LiaraStorageOK) String() string {
	return fmt.Sprintf("[POST /storage/s3/liara][%d] createS3LiaraStorageOK  %+v", 200, o.Payload)
}

func (o *CreateS3LiaraStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreateS3LiaraStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateS3LiaraStorageBadRequest creates a CreateS3LiaraStorageBadRequest with default headers values
func NewCreateS3LiaraStorageBadRequest() *CreateS3LiaraStorageBadRequest {
	return &CreateS3LiaraStorageBadRequest{}
}

/*
CreateS3LiaraStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateS3LiaraStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create s3 liara storage bad request response has a 2xx status code
func (o *CreateS3LiaraStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create s3 liara storage bad request response has a 3xx status code
func (o *CreateS3LiaraStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 liara storage bad request response has a 4xx status code
func (o *CreateS3LiaraStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create s3 liara storage bad request response has a 5xx status code
func (o *CreateS3LiaraStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create s3 liara storage bad request response a status code equal to that given
func (o *CreateS3LiaraStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create s3 liara storage bad request response
func (o *CreateS3LiaraStorageBadRequest) Code() int {
	return 400
}

func (o *CreateS3LiaraStorageBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/s3/liara][%d] createS3LiaraStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateS3LiaraStorageBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/s3/liara][%d] createS3LiaraStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateS3LiaraStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateS3LiaraStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateS3LiaraStorageInternalServerError creates a CreateS3LiaraStorageInternalServerError with default headers values
func NewCreateS3LiaraStorageInternalServerError() *CreateS3LiaraStorageInternalServerError {
	return &CreateS3LiaraStorageInternalServerError{}
}

/*
CreateS3LiaraStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateS3LiaraStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create s3 liara storage internal server error response has a 2xx status code
func (o *CreateS3LiaraStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create s3 liara storage internal server error response has a 3xx status code
func (o *CreateS3LiaraStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 liara storage internal server error response has a 4xx status code
func (o *CreateS3LiaraStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create s3 liara storage internal server error response has a 5xx status code
func (o *CreateS3LiaraStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create s3 liara storage internal server error response a status code equal to that given
func (o *CreateS3LiaraStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create s3 liara storage internal server error response
func (o *CreateS3LiaraStorageInternalServerError) Code() int {
	return 500
}

func (o *CreateS3LiaraStorageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/s3/liara][%d] createS3LiaraStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateS3LiaraStorageInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/s3/liara][%d] createS3LiaraStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateS3LiaraStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateS3LiaraStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}