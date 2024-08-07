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

// CreateS3DreamhostStorageReader is a Reader for the CreateS3DreamhostStorage structure.
type CreateS3DreamhostStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateS3DreamhostStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateS3DreamhostStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateS3DreamhostStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateS3DreamhostStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/s3/dreamhost] CreateS3DreamhostStorage", response, response.Code())
	}
}

// NewCreateS3DreamhostStorageOK creates a CreateS3DreamhostStorageOK with default headers values
func NewCreateS3DreamhostStorageOK() *CreateS3DreamhostStorageOK {
	return &CreateS3DreamhostStorageOK{}
}

/*
CreateS3DreamhostStorageOK describes a response with status code 200, with default header values.

OK
*/
type CreateS3DreamhostStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create s3 dreamhost storage o k response has a 2xx status code
func (o *CreateS3DreamhostStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create s3 dreamhost storage o k response has a 3xx status code
func (o *CreateS3DreamhostStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 dreamhost storage o k response has a 4xx status code
func (o *CreateS3DreamhostStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create s3 dreamhost storage o k response has a 5xx status code
func (o *CreateS3DreamhostStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create s3 dreamhost storage o k response a status code equal to that given
func (o *CreateS3DreamhostStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create s3 dreamhost storage o k response
func (o *CreateS3DreamhostStorageOK) Code() int {
	return 200
}

func (o *CreateS3DreamhostStorageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/s3/dreamhost][%d] createS3DreamhostStorageOK %s", 200, payload)
}

func (o *CreateS3DreamhostStorageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/s3/dreamhost][%d] createS3DreamhostStorageOK %s", 200, payload)
}

func (o *CreateS3DreamhostStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreateS3DreamhostStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateS3DreamhostStorageBadRequest creates a CreateS3DreamhostStorageBadRequest with default headers values
func NewCreateS3DreamhostStorageBadRequest() *CreateS3DreamhostStorageBadRequest {
	return &CreateS3DreamhostStorageBadRequest{}
}

/*
CreateS3DreamhostStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateS3DreamhostStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create s3 dreamhost storage bad request response has a 2xx status code
func (o *CreateS3DreamhostStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create s3 dreamhost storage bad request response has a 3xx status code
func (o *CreateS3DreamhostStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 dreamhost storage bad request response has a 4xx status code
func (o *CreateS3DreamhostStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create s3 dreamhost storage bad request response has a 5xx status code
func (o *CreateS3DreamhostStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create s3 dreamhost storage bad request response a status code equal to that given
func (o *CreateS3DreamhostStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create s3 dreamhost storage bad request response
func (o *CreateS3DreamhostStorageBadRequest) Code() int {
	return 400
}

func (o *CreateS3DreamhostStorageBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/s3/dreamhost][%d] createS3DreamhostStorageBadRequest %s", 400, payload)
}

func (o *CreateS3DreamhostStorageBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/s3/dreamhost][%d] createS3DreamhostStorageBadRequest %s", 400, payload)
}

func (o *CreateS3DreamhostStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateS3DreamhostStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateS3DreamhostStorageInternalServerError creates a CreateS3DreamhostStorageInternalServerError with default headers values
func NewCreateS3DreamhostStorageInternalServerError() *CreateS3DreamhostStorageInternalServerError {
	return &CreateS3DreamhostStorageInternalServerError{}
}

/*
CreateS3DreamhostStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateS3DreamhostStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create s3 dreamhost storage internal server error response has a 2xx status code
func (o *CreateS3DreamhostStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create s3 dreamhost storage internal server error response has a 3xx status code
func (o *CreateS3DreamhostStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 dreamhost storage internal server error response has a 4xx status code
func (o *CreateS3DreamhostStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create s3 dreamhost storage internal server error response has a 5xx status code
func (o *CreateS3DreamhostStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create s3 dreamhost storage internal server error response a status code equal to that given
func (o *CreateS3DreamhostStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create s3 dreamhost storage internal server error response
func (o *CreateS3DreamhostStorageInternalServerError) Code() int {
	return 500
}

func (o *CreateS3DreamhostStorageInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/s3/dreamhost][%d] createS3DreamhostStorageInternalServerError %s", 500, payload)
}

func (o *CreateS3DreamhostStorageInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/s3/dreamhost][%d] createS3DreamhostStorageInternalServerError %s", 500, payload)
}

func (o *CreateS3DreamhostStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateS3DreamhostStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
