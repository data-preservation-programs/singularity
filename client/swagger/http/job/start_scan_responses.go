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

// StartScanReader is a Reader for the StartScan structure.
type StartScanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartScanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartScanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStartScanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStartScanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /preparation/{id}/source/{name}/start-scan] StartScan", response, response.Code())
	}
}

// NewStartScanOK creates a StartScanOK with default headers values
func NewStartScanOK() *StartScanOK {
	return &StartScanOK{}
}

/*
StartScanOK describes a response with status code 200, with default header values.

OK
*/
type StartScanOK struct {
	Payload *models.ModelJob
}

// IsSuccess returns true when this start scan o k response has a 2xx status code
func (o *StartScanOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start scan o k response has a 3xx status code
func (o *StartScanOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start scan o k response has a 4xx status code
func (o *StartScanOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this start scan o k response has a 5xx status code
func (o *StartScanOK) IsServerError() bool {
	return false
}

// IsCode returns true when this start scan o k response a status code equal to that given
func (o *StartScanOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the start scan o k response
func (o *StartScanOK) Code() int {
	return 200
}

func (o *StartScanOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/start-scan][%d] startScanOK %s", 200, payload)
}

func (o *StartScanOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/start-scan][%d] startScanOK %s", 200, payload)
}

func (o *StartScanOK) GetPayload() *models.ModelJob {
	return o.Payload
}

func (o *StartScanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartScanBadRequest creates a StartScanBadRequest with default headers values
func NewStartScanBadRequest() *StartScanBadRequest {
	return &StartScanBadRequest{}
}

/*
StartScanBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StartScanBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this start scan bad request response has a 2xx status code
func (o *StartScanBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start scan bad request response has a 3xx status code
func (o *StartScanBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start scan bad request response has a 4xx status code
func (o *StartScanBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this start scan bad request response has a 5xx status code
func (o *StartScanBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this start scan bad request response a status code equal to that given
func (o *StartScanBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the start scan bad request response
func (o *StartScanBadRequest) Code() int {
	return 400
}

func (o *StartScanBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/start-scan][%d] startScanBadRequest %s", 400, payload)
}

func (o *StartScanBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/start-scan][%d] startScanBadRequest %s", 400, payload)
}

func (o *StartScanBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *StartScanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartScanInternalServerError creates a StartScanInternalServerError with default headers values
func NewStartScanInternalServerError() *StartScanInternalServerError {
	return &StartScanInternalServerError{}
}

/*
StartScanInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type StartScanInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this start scan internal server error response has a 2xx status code
func (o *StartScanInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start scan internal server error response has a 3xx status code
func (o *StartScanInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start scan internal server error response has a 4xx status code
func (o *StartScanInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this start scan internal server error response has a 5xx status code
func (o *StartScanInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this start scan internal server error response a status code equal to that given
func (o *StartScanInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the start scan internal server error response
func (o *StartScanInternalServerError) Code() int {
	return 500
}

func (o *StartScanInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/start-scan][%d] startScanInternalServerError %s", 500, payload)
}

func (o *StartScanInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /preparation/{id}/source/{name}/start-scan][%d] startScanInternalServerError %s", 500, payload)
}

func (o *StartScanInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *StartScanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
