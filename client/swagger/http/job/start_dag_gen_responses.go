// Code generated by go-swagger; DO NOT EDIT.

package job

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

// StartDagGenReader is a Reader for the StartDagGen structure.
type StartDagGenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartDagGenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartDagGenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStartDagGenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStartDagGenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /preparation/{id}/source/{name}/start-daggen] StartDagGen", response, response.Code())
	}
}

// NewStartDagGenOK creates a StartDagGenOK with default headers values
func NewStartDagGenOK() *StartDagGenOK {
	return &StartDagGenOK{}
}

/*
StartDagGenOK describes a response with status code 200, with default header values.

OK
*/
type StartDagGenOK struct {
	Payload *models.ModelJob
}

// IsSuccess returns true when this start dag gen o k response has a 2xx status code
func (o *StartDagGenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start dag gen o k response has a 3xx status code
func (o *StartDagGenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start dag gen o k response has a 4xx status code
func (o *StartDagGenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this start dag gen o k response has a 5xx status code
func (o *StartDagGenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this start dag gen o k response a status code equal to that given
func (o *StartDagGenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the start dag gen o k response
func (o *StartDagGenOK) Code() int {
	return 200
}

func (o *StartDagGenOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/start-daggen][%d] startDagGenOK %s", 200, payload)
}

func (o *StartDagGenOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/start-daggen][%d] startDagGenOK %s", 200, payload)
}

func (o *StartDagGenOK) GetPayload() *models.ModelJob {
	return o.Payload
}

func (o *StartDagGenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartDagGenBadRequest creates a StartDagGenBadRequest with default headers values
func NewStartDagGenBadRequest() *StartDagGenBadRequest {
	return &StartDagGenBadRequest{}
}

/*
StartDagGenBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StartDagGenBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this start dag gen bad request response has a 2xx status code
func (o *StartDagGenBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start dag gen bad request response has a 3xx status code
func (o *StartDagGenBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start dag gen bad request response has a 4xx status code
func (o *StartDagGenBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this start dag gen bad request response has a 5xx status code
func (o *StartDagGenBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this start dag gen bad request response a status code equal to that given
func (o *StartDagGenBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the start dag gen bad request response
func (o *StartDagGenBadRequest) Code() int {
	return 400
}

func (o *StartDagGenBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/start-daggen][%d] startDagGenBadRequest %s", 400, payload)
}

func (o *StartDagGenBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/start-daggen][%d] startDagGenBadRequest %s", 400, payload)
}

func (o *StartDagGenBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *StartDagGenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartDagGenInternalServerError creates a StartDagGenInternalServerError with default headers values
func NewStartDagGenInternalServerError() *StartDagGenInternalServerError {
	return &StartDagGenInternalServerError{}
}

/*
StartDagGenInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type StartDagGenInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this start dag gen internal server error response has a 2xx status code
func (o *StartDagGenInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start dag gen internal server error response has a 3xx status code
func (o *StartDagGenInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start dag gen internal server error response has a 4xx status code
func (o *StartDagGenInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this start dag gen internal server error response has a 5xx status code
func (o *StartDagGenInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this start dag gen internal server error response a status code equal to that given
func (o *StartDagGenInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the start dag gen internal server error response
func (o *StartDagGenInternalServerError) Code() int {
	return 500
}

func (o *StartDagGenInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/start-daggen][%d] startDagGenInternalServerError %s", 500, payload)
}

func (o *StartDagGenInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/start-daggen][%d] startDagGenInternalServerError %s", 500, payload)
}

func (o *StartDagGenInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *StartDagGenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
