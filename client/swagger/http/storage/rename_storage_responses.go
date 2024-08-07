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

// RenameStorageReader is a Reader for the RenameStorage structure.
type RenameStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RenameStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRenameStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRenameStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRenameStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /storage/{name}/rename] RenameStorage", response, response.Code())
	}
}

// NewRenameStorageOK creates a RenameStorageOK with default headers values
func NewRenameStorageOK() *RenameStorageOK {
	return &RenameStorageOK{}
}

/*
RenameStorageOK describes a response with status code 200, with default header values.

OK
*/
type RenameStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this rename storage o k response has a 2xx status code
func (o *RenameStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rename storage o k response has a 3xx status code
func (o *RenameStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rename storage o k response has a 4xx status code
func (o *RenameStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this rename storage o k response has a 5xx status code
func (o *RenameStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this rename storage o k response a status code equal to that given
func (o *RenameStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the rename storage o k response
func (o *RenameStorageOK) Code() int {
	return 200
}

func (o *RenameStorageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/{name}/rename][%d] renameStorageOK %s", 200, payload)
}

func (o *RenameStorageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/{name}/rename][%d] renameStorageOK %s", 200, payload)
}

func (o *RenameStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *RenameStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenameStorageBadRequest creates a RenameStorageBadRequest with default headers values
func NewRenameStorageBadRequest() *RenameStorageBadRequest {
	return &RenameStorageBadRequest{}
}

/*
RenameStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RenameStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this rename storage bad request response has a 2xx status code
func (o *RenameStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rename storage bad request response has a 3xx status code
func (o *RenameStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rename storage bad request response has a 4xx status code
func (o *RenameStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this rename storage bad request response has a 5xx status code
func (o *RenameStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this rename storage bad request response a status code equal to that given
func (o *RenameStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the rename storage bad request response
func (o *RenameStorageBadRequest) Code() int {
	return 400
}

func (o *RenameStorageBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/{name}/rename][%d] renameStorageBadRequest %s", 400, payload)
}

func (o *RenameStorageBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/{name}/rename][%d] renameStorageBadRequest %s", 400, payload)
}

func (o *RenameStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *RenameStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenameStorageInternalServerError creates a RenameStorageInternalServerError with default headers values
func NewRenameStorageInternalServerError() *RenameStorageInternalServerError {
	return &RenameStorageInternalServerError{}
}

/*
RenameStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RenameStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this rename storage internal server error response has a 2xx status code
func (o *RenameStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rename storage internal server error response has a 3xx status code
func (o *RenameStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rename storage internal server error response has a 4xx status code
func (o *RenameStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this rename storage internal server error response has a 5xx status code
func (o *RenameStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this rename storage internal server error response a status code equal to that given
func (o *RenameStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the rename storage internal server error response
func (o *RenameStorageInternalServerError) Code() int {
	return 500
}

func (o *RenameStorageInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/{name}/rename][%d] renameStorageInternalServerError %s", 500, payload)
}

func (o *RenameStorageInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/{name}/rename][%d] renameStorageInternalServerError %s", 500, payload)
}

func (o *RenameStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *RenameStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
