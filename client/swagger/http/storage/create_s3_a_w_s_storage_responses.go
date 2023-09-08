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

// CreateS3AWSStorageReader is a Reader for the CreateS3AWSStorage structure.
type CreateS3AWSStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateS3AWSStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateS3AWSStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateS3AWSStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateS3AWSStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/s3/aws] CreateS3AWSStorage", response, response.Code())
	}
}

// NewCreateS3AWSStorageOK creates a CreateS3AWSStorageOK with default headers values
func NewCreateS3AWSStorageOK() *CreateS3AWSStorageOK {
	return &CreateS3AWSStorageOK{}
}

/*
CreateS3AWSStorageOK describes a response with status code 200, with default header values.

OK
*/
type CreateS3AWSStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create s3 a w s storage o k response has a 2xx status code
func (o *CreateS3AWSStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create s3 a w s storage o k response has a 3xx status code
func (o *CreateS3AWSStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 a w s storage o k response has a 4xx status code
func (o *CreateS3AWSStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create s3 a w s storage o k response has a 5xx status code
func (o *CreateS3AWSStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create s3 a w s storage o k response a status code equal to that given
func (o *CreateS3AWSStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create s3 a w s storage o k response
func (o *CreateS3AWSStorageOK) Code() int {
	return 200
}

func (o *CreateS3AWSStorageOK) Error() string {
	return fmt.Sprintf("[POST /storage/s3/aws][%d] createS3AWSStorageOK  %+v", 200, o.Payload)
}

func (o *CreateS3AWSStorageOK) String() string {
	return fmt.Sprintf("[POST /storage/s3/aws][%d] createS3AWSStorageOK  %+v", 200, o.Payload)
}

func (o *CreateS3AWSStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreateS3AWSStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateS3AWSStorageBadRequest creates a CreateS3AWSStorageBadRequest with default headers values
func NewCreateS3AWSStorageBadRequest() *CreateS3AWSStorageBadRequest {
	return &CreateS3AWSStorageBadRequest{}
}

/*
CreateS3AWSStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateS3AWSStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create s3 a w s storage bad request response has a 2xx status code
func (o *CreateS3AWSStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create s3 a w s storage bad request response has a 3xx status code
func (o *CreateS3AWSStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 a w s storage bad request response has a 4xx status code
func (o *CreateS3AWSStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create s3 a w s storage bad request response has a 5xx status code
func (o *CreateS3AWSStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create s3 a w s storage bad request response a status code equal to that given
func (o *CreateS3AWSStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create s3 a w s storage bad request response
func (o *CreateS3AWSStorageBadRequest) Code() int {
	return 400
}

func (o *CreateS3AWSStorageBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/s3/aws][%d] createS3AWSStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateS3AWSStorageBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/s3/aws][%d] createS3AWSStorageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateS3AWSStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateS3AWSStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateS3AWSStorageInternalServerError creates a CreateS3AWSStorageInternalServerError with default headers values
func NewCreateS3AWSStorageInternalServerError() *CreateS3AWSStorageInternalServerError {
	return &CreateS3AWSStorageInternalServerError{}
}

/*
CreateS3AWSStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateS3AWSStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create s3 a w s storage internal server error response has a 2xx status code
func (o *CreateS3AWSStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create s3 a w s storage internal server error response has a 3xx status code
func (o *CreateS3AWSStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 a w s storage internal server error response has a 4xx status code
func (o *CreateS3AWSStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create s3 a w s storage internal server error response has a 5xx status code
func (o *CreateS3AWSStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create s3 a w s storage internal server error response a status code equal to that given
func (o *CreateS3AWSStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create s3 a w s storage internal server error response
func (o *CreateS3AWSStorageInternalServerError) Code() int {
	return 500
}

func (o *CreateS3AWSStorageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/s3/aws][%d] createS3AWSStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateS3AWSStorageInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/s3/aws][%d] createS3AWSStorageInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateS3AWSStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateS3AWSStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}