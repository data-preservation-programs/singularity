// Code generated by go-swagger; DO NOT EDIT.

package data_source

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// PostSourceIDCheckReader is a Reader for the PostSourceIDCheck structure.
type PostSourceIDCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSourceIDCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSourceIDCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSourceIDCheckBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSourceIDCheckInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /source/{id}/check] PostSourceIDCheck", response, response.Code())
	}
}

// NewPostSourceIDCheckOK creates a PostSourceIDCheckOK with default headers values
func NewPostSourceIDCheckOK() *PostSourceIDCheckOK {
	return &PostSourceIDCheckOK{}
}

/*
PostSourceIDCheckOK describes a response with status code 200, with default header values.

OK
*/
type PostSourceIDCheckOK struct {
	Payload []*models.GithubComDataPreservationProgramsSingularityHandlerDatasourceEntry
}

// IsSuccess returns true when this post source Id check o k response has a 2xx status code
func (o *PostSourceIDCheckOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post source Id check o k response has a 3xx status code
func (o *PostSourceIDCheckOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source Id check o k response has a 4xx status code
func (o *PostSourceIDCheckOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source Id check o k response has a 5xx status code
func (o *PostSourceIDCheckOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post source Id check o k response a status code equal to that given
func (o *PostSourceIDCheckOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post source Id check o k response
func (o *PostSourceIDCheckOK) Code() int {
	return 200
}

func (o *PostSourceIDCheckOK) Error() string {
	return fmt.Sprintf("[POST /source/{id}/check][%d] postSourceIdCheckOK  %+v", 200, o.Payload)
}

func (o *PostSourceIDCheckOK) String() string {
	return fmt.Sprintf("[POST /source/{id}/check][%d] postSourceIdCheckOK  %+v", 200, o.Payload)
}

func (o *PostSourceIDCheckOK) GetPayload() []*models.GithubComDataPreservationProgramsSingularityHandlerDatasourceEntry {
	return o.Payload
}

func (o *PostSourceIDCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceIDCheckBadRequest creates a PostSourceIDCheckBadRequest with default headers values
func NewPostSourceIDCheckBadRequest() *PostSourceIDCheckBadRequest {
	return &PostSourceIDCheckBadRequest{}
}

/*
PostSourceIDCheckBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostSourceIDCheckBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source Id check bad request response has a 2xx status code
func (o *PostSourceIDCheckBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source Id check bad request response has a 3xx status code
func (o *PostSourceIDCheckBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source Id check bad request response has a 4xx status code
func (o *PostSourceIDCheckBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post source Id check bad request response has a 5xx status code
func (o *PostSourceIDCheckBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post source Id check bad request response a status code equal to that given
func (o *PostSourceIDCheckBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post source Id check bad request response
func (o *PostSourceIDCheckBadRequest) Code() int {
	return 400
}

func (o *PostSourceIDCheckBadRequest) Error() string {
	return fmt.Sprintf("[POST /source/{id}/check][%d] postSourceIdCheckBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceIDCheckBadRequest) String() string {
	return fmt.Sprintf("[POST /source/{id}/check][%d] postSourceIdCheckBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceIDCheckBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourceIDCheckBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceIDCheckInternalServerError creates a PostSourceIDCheckInternalServerError with default headers values
func NewPostSourceIDCheckInternalServerError() *PostSourceIDCheckInternalServerError {
	return &PostSourceIDCheckInternalServerError{}
}

/*
PostSourceIDCheckInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostSourceIDCheckInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source Id check internal server error response has a 2xx status code
func (o *PostSourceIDCheckInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source Id check internal server error response has a 3xx status code
func (o *PostSourceIDCheckInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source Id check internal server error response has a 4xx status code
func (o *PostSourceIDCheckInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source Id check internal server error response has a 5xx status code
func (o *PostSourceIDCheckInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post source Id check internal server error response a status code equal to that given
func (o *PostSourceIDCheckInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post source Id check internal server error response
func (o *PostSourceIDCheckInternalServerError) Code() int {
	return 500
}

func (o *PostSourceIDCheckInternalServerError) Error() string {
	return fmt.Sprintf("[POST /source/{id}/check][%d] postSourceIdCheckInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceIDCheckInternalServerError) String() string {
	return fmt.Sprintf("[POST /source/{id}/check][%d] postSourceIdCheckInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceIDCheckInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourceIDCheckInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
