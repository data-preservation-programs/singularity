// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// CreateS3CloudflareStorageReader is a Reader for the CreateS3CloudflareStorage structure.
type CreateS3CloudflareStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateS3CloudflareStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateS3CloudflareStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateS3CloudflareStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateS3CloudflareStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/s3/cloudflare] CreateS3CloudflareStorage", response, response.Code())
	}
}

// NewCreateS3CloudflareStorageOK creates a CreateS3CloudflareStorageOK with default headers values
func NewCreateS3CloudflareStorageOK() *CreateS3CloudflareStorageOK {
	return &CreateS3CloudflareStorageOK{}
}

/*
CreateS3CloudflareStorageOK describes a response with status code 200, with default header values.

OK
*/
type CreateS3CloudflareStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create s3 cloudflare storage o k response has a 2xx status code
func (o *CreateS3CloudflareStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create s3 cloudflare storage o k response has a 3xx status code
func (o *CreateS3CloudflareStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 cloudflare storage o k response has a 4xx status code
func (o *CreateS3CloudflareStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create s3 cloudflare storage o k response has a 5xx status code
func (o *CreateS3CloudflareStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create s3 cloudflare storage o k response a status code equal to that given
func (o *CreateS3CloudflareStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create s3 cloudflare storage o k response
func (o *CreateS3CloudflareStorageOK) Code() int {
	return 200
}

func (o *CreateS3CloudflareStorageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/s3/cloudflare][%d] createS3CloudflareStorageOK %s", 200, payload)
}

func (o *CreateS3CloudflareStorageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/s3/cloudflare][%d] createS3CloudflareStorageOK %s", 200, payload)
}

func (o *CreateS3CloudflareStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreateS3CloudflareStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateS3CloudflareStorageBadRequest creates a CreateS3CloudflareStorageBadRequest with default headers values
func NewCreateS3CloudflareStorageBadRequest() *CreateS3CloudflareStorageBadRequest {
	return &CreateS3CloudflareStorageBadRequest{}
}

/*
CreateS3CloudflareStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateS3CloudflareStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create s3 cloudflare storage bad request response has a 2xx status code
func (o *CreateS3CloudflareStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create s3 cloudflare storage bad request response has a 3xx status code
func (o *CreateS3CloudflareStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 cloudflare storage bad request response has a 4xx status code
func (o *CreateS3CloudflareStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create s3 cloudflare storage bad request response has a 5xx status code
func (o *CreateS3CloudflareStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create s3 cloudflare storage bad request response a status code equal to that given
func (o *CreateS3CloudflareStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create s3 cloudflare storage bad request response
func (o *CreateS3CloudflareStorageBadRequest) Code() int {
	return 400
}

func (o *CreateS3CloudflareStorageBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/s3/cloudflare][%d] createS3CloudflareStorageBadRequest %s", 400, payload)
}

func (o *CreateS3CloudflareStorageBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/s3/cloudflare][%d] createS3CloudflareStorageBadRequest %s", 400, payload)
}

func (o *CreateS3CloudflareStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateS3CloudflareStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateS3CloudflareStorageInternalServerError creates a CreateS3CloudflareStorageInternalServerError with default headers values
func NewCreateS3CloudflareStorageInternalServerError() *CreateS3CloudflareStorageInternalServerError {
	return &CreateS3CloudflareStorageInternalServerError{}
}

/*
CreateS3CloudflareStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateS3CloudflareStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create s3 cloudflare storage internal server error response has a 2xx status code
func (o *CreateS3CloudflareStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create s3 cloudflare storage internal server error response has a 3xx status code
func (o *CreateS3CloudflareStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 cloudflare storage internal server error response has a 4xx status code
func (o *CreateS3CloudflareStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create s3 cloudflare storage internal server error response has a 5xx status code
func (o *CreateS3CloudflareStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create s3 cloudflare storage internal server error response a status code equal to that given
func (o *CreateS3CloudflareStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create s3 cloudflare storage internal server error response
func (o *CreateS3CloudflareStorageInternalServerError) Code() int {
	return 500
}

func (o *CreateS3CloudflareStorageInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/s3/cloudflare][%d] createS3CloudflareStorageInternalServerError %s", 500, payload)
}

func (o *CreateS3CloudflareStorageInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/s3/cloudflare][%d] createS3CloudflareStorageInternalServerError %s", 500, payload)
}

func (o *CreateS3CloudflareStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateS3CloudflareStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
