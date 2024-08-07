// Code generated by go-swagger; DO NOT EDIT.

package deal_schedule

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

// ResumeScheduleReader is a Reader for the ResumeSchedule structure.
type ResumeScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResumeScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResumeScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewResumeScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewResumeScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /schedule/{id}/resume] ResumeSchedule", response, response.Code())
	}
}

// NewResumeScheduleOK creates a ResumeScheduleOK with default headers values
func NewResumeScheduleOK() *ResumeScheduleOK {
	return &ResumeScheduleOK{}
}

/*
ResumeScheduleOK describes a response with status code 200, with default header values.

OK
*/
type ResumeScheduleOK struct {
	Payload *models.ModelSchedule
}

// IsSuccess returns true when this resume schedule o k response has a 2xx status code
func (o *ResumeScheduleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resume schedule o k response has a 3xx status code
func (o *ResumeScheduleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resume schedule o k response has a 4xx status code
func (o *ResumeScheduleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this resume schedule o k response has a 5xx status code
func (o *ResumeScheduleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this resume schedule o k response a status code equal to that given
func (o *ResumeScheduleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the resume schedule o k response
func (o *ResumeScheduleOK) Code() int {
	return 200
}

func (o *ResumeScheduleOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /schedule/{id}/resume][%d] resumeScheduleOK %s", 200, payload)
}

func (o *ResumeScheduleOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /schedule/{id}/resume][%d] resumeScheduleOK %s", 200, payload)
}

func (o *ResumeScheduleOK) GetPayload() *models.ModelSchedule {
	return o.Payload
}

func (o *ResumeScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelSchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResumeScheduleBadRequest creates a ResumeScheduleBadRequest with default headers values
func NewResumeScheduleBadRequest() *ResumeScheduleBadRequest {
	return &ResumeScheduleBadRequest{}
}

/*
ResumeScheduleBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ResumeScheduleBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this resume schedule bad request response has a 2xx status code
func (o *ResumeScheduleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resume schedule bad request response has a 3xx status code
func (o *ResumeScheduleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resume schedule bad request response has a 4xx status code
func (o *ResumeScheduleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this resume schedule bad request response has a 5xx status code
func (o *ResumeScheduleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this resume schedule bad request response a status code equal to that given
func (o *ResumeScheduleBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the resume schedule bad request response
func (o *ResumeScheduleBadRequest) Code() int {
	return 400
}

func (o *ResumeScheduleBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /schedule/{id}/resume][%d] resumeScheduleBadRequest %s", 400, payload)
}

func (o *ResumeScheduleBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /schedule/{id}/resume][%d] resumeScheduleBadRequest %s", 400, payload)
}

func (o *ResumeScheduleBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *ResumeScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResumeScheduleInternalServerError creates a ResumeScheduleInternalServerError with default headers values
func NewResumeScheduleInternalServerError() *ResumeScheduleInternalServerError {
	return &ResumeScheduleInternalServerError{}
}

/*
ResumeScheduleInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ResumeScheduleInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this resume schedule internal server error response has a 2xx status code
func (o *ResumeScheduleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this resume schedule internal server error response has a 3xx status code
func (o *ResumeScheduleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resume schedule internal server error response has a 4xx status code
func (o *ResumeScheduleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this resume schedule internal server error response has a 5xx status code
func (o *ResumeScheduleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this resume schedule internal server error response a status code equal to that given
func (o *ResumeScheduleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the resume schedule internal server error response
func (o *ResumeScheduleInternalServerError) Code() int {
	return 500
}

func (o *ResumeScheduleInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /schedule/{id}/resume][%d] resumeScheduleInternalServerError %s", 500, payload)
}

func (o *ResumeScheduleInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /schedule/{id}/resume][%d] resumeScheduleInternalServerError %s", 500, payload)
}

func (o *ResumeScheduleInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *ResumeScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
