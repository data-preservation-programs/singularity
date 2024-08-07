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

// CreatePutioStorageReader is a Reader for the CreatePutioStorage structure.
type CreatePutioStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePutioStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePutioStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreatePutioStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreatePutioStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/putio] CreatePutioStorage", response, response.Code())
	}
}

// NewCreatePutioStorageOK creates a CreatePutioStorageOK with default headers values
func NewCreatePutioStorageOK() *CreatePutioStorageOK {
	return &CreatePutioStorageOK{}
}

/*
CreatePutioStorageOK describes a response with status code 200, with default header values.

OK
*/
type CreatePutioStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create putio storage o k response has a 2xx status code
func (o *CreatePutioStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create putio storage o k response has a 3xx status code
func (o *CreatePutioStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create putio storage o k response has a 4xx status code
func (o *CreatePutioStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create putio storage o k response has a 5xx status code
func (o *CreatePutioStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create putio storage o k response a status code equal to that given
func (o *CreatePutioStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create putio storage o k response
func (o *CreatePutioStorageOK) Code() int {
	return 200
}

func (o *CreatePutioStorageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/putio][%d] createPutioStorageOK %s", 200, payload)
}

func (o *CreatePutioStorageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/putio][%d] createPutioStorageOK %s", 200, payload)
}

func (o *CreatePutioStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreatePutioStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePutioStorageBadRequest creates a CreatePutioStorageBadRequest with default headers values
func NewCreatePutioStorageBadRequest() *CreatePutioStorageBadRequest {
	return &CreatePutioStorageBadRequest{}
}

/*
CreatePutioStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreatePutioStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create putio storage bad request response has a 2xx status code
func (o *CreatePutioStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create putio storage bad request response has a 3xx status code
func (o *CreatePutioStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create putio storage bad request response has a 4xx status code
func (o *CreatePutioStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create putio storage bad request response has a 5xx status code
func (o *CreatePutioStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create putio storage bad request response a status code equal to that given
func (o *CreatePutioStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create putio storage bad request response
func (o *CreatePutioStorageBadRequest) Code() int {
	return 400
}

func (o *CreatePutioStorageBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/putio][%d] createPutioStorageBadRequest %s", 400, payload)
}

func (o *CreatePutioStorageBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/putio][%d] createPutioStorageBadRequest %s", 400, payload)
}

func (o *CreatePutioStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreatePutioStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePutioStorageInternalServerError creates a CreatePutioStorageInternalServerError with default headers values
func NewCreatePutioStorageInternalServerError() *CreatePutioStorageInternalServerError {
	return &CreatePutioStorageInternalServerError{}
}

/*
CreatePutioStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreatePutioStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create putio storage internal server error response has a 2xx status code
func (o *CreatePutioStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create putio storage internal server error response has a 3xx status code
func (o *CreatePutioStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create putio storage internal server error response has a 4xx status code
func (o *CreatePutioStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create putio storage internal server error response has a 5xx status code
func (o *CreatePutioStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create putio storage internal server error response a status code equal to that given
func (o *CreatePutioStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create putio storage internal server error response
func (o *CreatePutioStorageInternalServerError) Code() int {
	return 500
}

func (o *CreatePutioStorageInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/putio][%d] createPutioStorageInternalServerError %s", 500, payload)
}

func (o *CreatePutioStorageInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/putio][%d] createPutioStorageInternalServerError %s", 500, payload)
}

func (o *CreatePutioStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreatePutioStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
