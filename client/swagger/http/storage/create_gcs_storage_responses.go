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

// CreateGcsStorageReader is a Reader for the CreateGcsStorage structure.
type CreateGcsStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateGcsStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateGcsStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateGcsStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateGcsStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/gcs] CreateGcsStorage", response, response.Code())
	}
}

// NewCreateGcsStorageOK creates a CreateGcsStorageOK with default headers values
func NewCreateGcsStorageOK() *CreateGcsStorageOK {
	return &CreateGcsStorageOK{}
}

/*
CreateGcsStorageOK describes a response with status code 200, with default header values.

OK
*/
type CreateGcsStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create gcs storage o k response has a 2xx status code
func (o *CreateGcsStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create gcs storage o k response has a 3xx status code
func (o *CreateGcsStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create gcs storage o k response has a 4xx status code
func (o *CreateGcsStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create gcs storage o k response has a 5xx status code
func (o *CreateGcsStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create gcs storage o k response a status code equal to that given
func (o *CreateGcsStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create gcs storage o k response
func (o *CreateGcsStorageOK) Code() int {
	return 200
}

func (o *CreateGcsStorageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/gcs][%d] createGcsStorageOK %s", 200, payload)
}

func (o *CreateGcsStorageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/gcs][%d] createGcsStorageOK %s", 200, payload)
}

func (o *CreateGcsStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreateGcsStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGcsStorageBadRequest creates a CreateGcsStorageBadRequest with default headers values
func NewCreateGcsStorageBadRequest() *CreateGcsStorageBadRequest {
	return &CreateGcsStorageBadRequest{}
}

/*
CreateGcsStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateGcsStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create gcs storage bad request response has a 2xx status code
func (o *CreateGcsStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create gcs storage bad request response has a 3xx status code
func (o *CreateGcsStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create gcs storage bad request response has a 4xx status code
func (o *CreateGcsStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create gcs storage bad request response has a 5xx status code
func (o *CreateGcsStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create gcs storage bad request response a status code equal to that given
func (o *CreateGcsStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create gcs storage bad request response
func (o *CreateGcsStorageBadRequest) Code() int {
	return 400
}

func (o *CreateGcsStorageBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/gcs][%d] createGcsStorageBadRequest %s", 400, payload)
}

func (o *CreateGcsStorageBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/gcs][%d] createGcsStorageBadRequest %s", 400, payload)
}

func (o *CreateGcsStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateGcsStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGcsStorageInternalServerError creates a CreateGcsStorageInternalServerError with default headers values
func NewCreateGcsStorageInternalServerError() *CreateGcsStorageInternalServerError {
	return &CreateGcsStorageInternalServerError{}
}

/*
CreateGcsStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateGcsStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create gcs storage internal server error response has a 2xx status code
func (o *CreateGcsStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create gcs storage internal server error response has a 3xx status code
func (o *CreateGcsStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create gcs storage internal server error response has a 4xx status code
func (o *CreateGcsStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create gcs storage internal server error response has a 5xx status code
func (o *CreateGcsStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create gcs storage internal server error response a status code equal to that given
func (o *CreateGcsStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create gcs storage internal server error response
func (o *CreateGcsStorageInternalServerError) Code() int {
	return 500
}

func (o *CreateGcsStorageInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/gcs][%d] createGcsStorageInternalServerError %s", 500, payload)
}

func (o *CreateGcsStorageInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/gcs][%d] createGcsStorageInternalServerError %s", 500, payload)
}

func (o *CreateGcsStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateGcsStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
