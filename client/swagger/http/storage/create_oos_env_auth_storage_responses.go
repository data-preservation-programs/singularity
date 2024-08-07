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

// CreateOosEnvAuthStorageReader is a Reader for the CreateOosEnvAuthStorage structure.
type CreateOosEnvAuthStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOosEnvAuthStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOosEnvAuthStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateOosEnvAuthStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateOosEnvAuthStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/oos/env_auth] CreateOosEnv_authStorage", response, response.Code())
	}
}

// NewCreateOosEnvAuthStorageOK creates a CreateOosEnvAuthStorageOK with default headers values
func NewCreateOosEnvAuthStorageOK() *CreateOosEnvAuthStorageOK {
	return &CreateOosEnvAuthStorageOK{}
}

/*
CreateOosEnvAuthStorageOK describes a response with status code 200, with default header values.

OK
*/
type CreateOosEnvAuthStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create oos env auth storage o k response has a 2xx status code
func (o *CreateOosEnvAuthStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create oos env auth storage o k response has a 3xx status code
func (o *CreateOosEnvAuthStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create oos env auth storage o k response has a 4xx status code
func (o *CreateOosEnvAuthStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create oos env auth storage o k response has a 5xx status code
func (o *CreateOosEnvAuthStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create oos env auth storage o k response a status code equal to that given
func (o *CreateOosEnvAuthStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create oos env auth storage o k response
func (o *CreateOosEnvAuthStorageOK) Code() int {
	return 200
}

func (o *CreateOosEnvAuthStorageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/oos/env_auth][%d] createOosEnvAuthStorageOK %s", 200, payload)
}

func (o *CreateOosEnvAuthStorageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/oos/env_auth][%d] createOosEnvAuthStorageOK %s", 200, payload)
}

func (o *CreateOosEnvAuthStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreateOosEnvAuthStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOosEnvAuthStorageBadRequest creates a CreateOosEnvAuthStorageBadRequest with default headers values
func NewCreateOosEnvAuthStorageBadRequest() *CreateOosEnvAuthStorageBadRequest {
	return &CreateOosEnvAuthStorageBadRequest{}
}

/*
CreateOosEnvAuthStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateOosEnvAuthStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create oos env auth storage bad request response has a 2xx status code
func (o *CreateOosEnvAuthStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create oos env auth storage bad request response has a 3xx status code
func (o *CreateOosEnvAuthStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create oos env auth storage bad request response has a 4xx status code
func (o *CreateOosEnvAuthStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create oos env auth storage bad request response has a 5xx status code
func (o *CreateOosEnvAuthStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create oos env auth storage bad request response a status code equal to that given
func (o *CreateOosEnvAuthStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create oos env auth storage bad request response
func (o *CreateOosEnvAuthStorageBadRequest) Code() int {
	return 400
}

func (o *CreateOosEnvAuthStorageBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/oos/env_auth][%d] createOosEnvAuthStorageBadRequest %s", 400, payload)
}

func (o *CreateOosEnvAuthStorageBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/oos/env_auth][%d] createOosEnvAuthStorageBadRequest %s", 400, payload)
}

func (o *CreateOosEnvAuthStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateOosEnvAuthStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOosEnvAuthStorageInternalServerError creates a CreateOosEnvAuthStorageInternalServerError with default headers values
func NewCreateOosEnvAuthStorageInternalServerError() *CreateOosEnvAuthStorageInternalServerError {
	return &CreateOosEnvAuthStorageInternalServerError{}
}

/*
CreateOosEnvAuthStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateOosEnvAuthStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create oos env auth storage internal server error response has a 2xx status code
func (o *CreateOosEnvAuthStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create oos env auth storage internal server error response has a 3xx status code
func (o *CreateOosEnvAuthStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create oos env auth storage internal server error response has a 4xx status code
func (o *CreateOosEnvAuthStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create oos env auth storage internal server error response has a 5xx status code
func (o *CreateOosEnvAuthStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create oos env auth storage internal server error response a status code equal to that given
func (o *CreateOosEnvAuthStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create oos env auth storage internal server error response
func (o *CreateOosEnvAuthStorageInternalServerError) Code() int {
	return 500
}

func (o *CreateOosEnvAuthStorageInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/oos/env_auth][%d] createOosEnvAuthStorageInternalServerError %s", 500, payload)
}

func (o *CreateOosEnvAuthStorageInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/oos/env_auth][%d] createOosEnvAuthStorageInternalServerError %s", 500, payload)
}

func (o *CreateOosEnvAuthStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateOosEnvAuthStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
