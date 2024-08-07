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

// CreatePremiumizemeStorageReader is a Reader for the CreatePremiumizemeStorage structure.
type CreatePremiumizemeStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePremiumizemeStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePremiumizemeStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreatePremiumizemeStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreatePremiumizemeStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/premiumizeme] CreatePremiumizemeStorage", response, response.Code())
	}
}

// NewCreatePremiumizemeStorageOK creates a CreatePremiumizemeStorageOK with default headers values
func NewCreatePremiumizemeStorageOK() *CreatePremiumizemeStorageOK {
	return &CreatePremiumizemeStorageOK{}
}

/*
CreatePremiumizemeStorageOK describes a response with status code 200, with default header values.

OK
*/
type CreatePremiumizemeStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create premiumizeme storage o k response has a 2xx status code
func (o *CreatePremiumizemeStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create premiumizeme storage o k response has a 3xx status code
func (o *CreatePremiumizemeStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create premiumizeme storage o k response has a 4xx status code
func (o *CreatePremiumizemeStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create premiumizeme storage o k response has a 5xx status code
func (o *CreatePremiumizemeStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create premiumizeme storage o k response a status code equal to that given
func (o *CreatePremiumizemeStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create premiumizeme storage o k response
func (o *CreatePremiumizemeStorageOK) Code() int {
	return 200
}

func (o *CreatePremiumizemeStorageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/premiumizeme][%d] createPremiumizemeStorageOK %s", 200, payload)
}

func (o *CreatePremiumizemeStorageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/premiumizeme][%d] createPremiumizemeStorageOK %s", 200, payload)
}

func (o *CreatePremiumizemeStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreatePremiumizemeStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePremiumizemeStorageBadRequest creates a CreatePremiumizemeStorageBadRequest with default headers values
func NewCreatePremiumizemeStorageBadRequest() *CreatePremiumizemeStorageBadRequest {
	return &CreatePremiumizemeStorageBadRequest{}
}

/*
CreatePremiumizemeStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreatePremiumizemeStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create premiumizeme storage bad request response has a 2xx status code
func (o *CreatePremiumizemeStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create premiumizeme storage bad request response has a 3xx status code
func (o *CreatePremiumizemeStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create premiumizeme storage bad request response has a 4xx status code
func (o *CreatePremiumizemeStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create premiumizeme storage bad request response has a 5xx status code
func (o *CreatePremiumizemeStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create premiumizeme storage bad request response a status code equal to that given
func (o *CreatePremiumizemeStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create premiumizeme storage bad request response
func (o *CreatePremiumizemeStorageBadRequest) Code() int {
	return 400
}

func (o *CreatePremiumizemeStorageBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/premiumizeme][%d] createPremiumizemeStorageBadRequest %s", 400, payload)
}

func (o *CreatePremiumizemeStorageBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/premiumizeme][%d] createPremiumizemeStorageBadRequest %s", 400, payload)
}

func (o *CreatePremiumizemeStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreatePremiumizemeStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePremiumizemeStorageInternalServerError creates a CreatePremiumizemeStorageInternalServerError with default headers values
func NewCreatePremiumizemeStorageInternalServerError() *CreatePremiumizemeStorageInternalServerError {
	return &CreatePremiumizemeStorageInternalServerError{}
}

/*
CreatePremiumizemeStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreatePremiumizemeStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create premiumizeme storage internal server error response has a 2xx status code
func (o *CreatePremiumizemeStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create premiumizeme storage internal server error response has a 3xx status code
func (o *CreatePremiumizemeStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create premiumizeme storage internal server error response has a 4xx status code
func (o *CreatePremiumizemeStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create premiumizeme storage internal server error response has a 5xx status code
func (o *CreatePremiumizemeStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create premiumizeme storage internal server error response a status code equal to that given
func (o *CreatePremiumizemeStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create premiumizeme storage internal server error response
func (o *CreatePremiumizemeStorageInternalServerError) Code() int {
	return 500
}

func (o *CreatePremiumizemeStorageInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/premiumizeme][%d] createPremiumizemeStorageInternalServerError %s", 500, payload)
}

func (o *CreatePremiumizemeStorageInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/premiumizeme][%d] createPremiumizemeStorageInternalServerError %s", 500, payload)
}

func (o *CreatePremiumizemeStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreatePremiumizemeStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
