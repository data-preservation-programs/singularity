// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// PostStorageLocalReader is a Reader for the PostStorageLocal structure.
type PostStorageLocalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageLocalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStorageLocalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStorageLocalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStorageLocalInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/local] PostStorageLocal", response, response.Code())
	}
}

// NewPostStorageLocalOK creates a PostStorageLocalOK with default headers values
func NewPostStorageLocalOK() *PostStorageLocalOK {
	return &PostStorageLocalOK{}
}

/*
PostStorageLocalOK describes a response with status code 200, with default header values.

OK
*/
type PostStorageLocalOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this post storage local o k response has a 2xx status code
func (o *PostStorageLocalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post storage local o k response has a 3xx status code
func (o *PostStorageLocalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage local o k response has a 4xx status code
func (o *PostStorageLocalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage local o k response has a 5xx status code
func (o *PostStorageLocalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage local o k response a status code equal to that given
func (o *PostStorageLocalOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post storage local o k response
func (o *PostStorageLocalOK) Code() int {
	return 200
}

func (o *PostStorageLocalOK) Error() string {
	return fmt.Sprintf("[POST /storage/local][%d] postStorageLocalOK  %+v", 200, o.Payload)
}

func (o *PostStorageLocalOK) String() string {
	return fmt.Sprintf("[POST /storage/local][%d] postStorageLocalOK  %+v", 200, o.Payload)
}

func (o *PostStorageLocalOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *PostStorageLocalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageLocalBadRequest creates a PostStorageLocalBadRequest with default headers values
func NewPostStorageLocalBadRequest() *PostStorageLocalBadRequest {
	return &PostStorageLocalBadRequest{}
}

/*
PostStorageLocalBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostStorageLocalBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage local bad request response has a 2xx status code
func (o *PostStorageLocalBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage local bad request response has a 3xx status code
func (o *PostStorageLocalBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage local bad request response has a 4xx status code
func (o *PostStorageLocalBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post storage local bad request response has a 5xx status code
func (o *PostStorageLocalBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage local bad request response a status code equal to that given
func (o *PostStorageLocalBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post storage local bad request response
func (o *PostStorageLocalBadRequest) Code() int {
	return 400
}

func (o *PostStorageLocalBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/local][%d] postStorageLocalBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageLocalBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/local][%d] postStorageLocalBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageLocalBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageLocalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageLocalInternalServerError creates a PostStorageLocalInternalServerError with default headers values
func NewPostStorageLocalInternalServerError() *PostStorageLocalInternalServerError {
	return &PostStorageLocalInternalServerError{}
}

/*
PostStorageLocalInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostStorageLocalInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage local internal server error response has a 2xx status code
func (o *PostStorageLocalInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage local internal server error response has a 3xx status code
func (o *PostStorageLocalInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage local internal server error response has a 4xx status code
func (o *PostStorageLocalInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage local internal server error response has a 5xx status code
func (o *PostStorageLocalInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post storage local internal server error response a status code equal to that given
func (o *PostStorageLocalInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post storage local internal server error response
func (o *PostStorageLocalInternalServerError) Code() int {
	return 500
}

func (o *PostStorageLocalInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/local][%d] postStorageLocalInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageLocalInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/local][%d] postStorageLocalInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageLocalInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageLocalInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
