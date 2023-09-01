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

// PostStorageSeafileReader is a Reader for the PostStorageSeafile structure.
type PostStorageSeafileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageSeafileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStorageSeafileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStorageSeafileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStorageSeafileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/seafile] PostStorageSeafile", response, response.Code())
	}
}

// NewPostStorageSeafileOK creates a PostStorageSeafileOK with default headers values
func NewPostStorageSeafileOK() *PostStorageSeafileOK {
	return &PostStorageSeafileOK{}
}

/*
PostStorageSeafileOK describes a response with status code 200, with default header values.

OK
*/
type PostStorageSeafileOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this post storage seafile o k response has a 2xx status code
func (o *PostStorageSeafileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post storage seafile o k response has a 3xx status code
func (o *PostStorageSeafileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage seafile o k response has a 4xx status code
func (o *PostStorageSeafileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage seafile o k response has a 5xx status code
func (o *PostStorageSeafileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage seafile o k response a status code equal to that given
func (o *PostStorageSeafileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post storage seafile o k response
func (o *PostStorageSeafileOK) Code() int {
	return 200
}

func (o *PostStorageSeafileOK) Error() string {
	return fmt.Sprintf("[POST /storage/seafile][%d] postStorageSeafileOK  %+v", 200, o.Payload)
}

func (o *PostStorageSeafileOK) String() string {
	return fmt.Sprintf("[POST /storage/seafile][%d] postStorageSeafileOK  %+v", 200, o.Payload)
}

func (o *PostStorageSeafileOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *PostStorageSeafileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageSeafileBadRequest creates a PostStorageSeafileBadRequest with default headers values
func NewPostStorageSeafileBadRequest() *PostStorageSeafileBadRequest {
	return &PostStorageSeafileBadRequest{}
}

/*
PostStorageSeafileBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostStorageSeafileBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage seafile bad request response has a 2xx status code
func (o *PostStorageSeafileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage seafile bad request response has a 3xx status code
func (o *PostStorageSeafileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage seafile bad request response has a 4xx status code
func (o *PostStorageSeafileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post storage seafile bad request response has a 5xx status code
func (o *PostStorageSeafileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage seafile bad request response a status code equal to that given
func (o *PostStorageSeafileBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post storage seafile bad request response
func (o *PostStorageSeafileBadRequest) Code() int {
	return 400
}

func (o *PostStorageSeafileBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/seafile][%d] postStorageSeafileBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageSeafileBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/seafile][%d] postStorageSeafileBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageSeafileBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageSeafileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageSeafileInternalServerError creates a PostStorageSeafileInternalServerError with default headers values
func NewPostStorageSeafileInternalServerError() *PostStorageSeafileInternalServerError {
	return &PostStorageSeafileInternalServerError{}
}

/*
PostStorageSeafileInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostStorageSeafileInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage seafile internal server error response has a 2xx status code
func (o *PostStorageSeafileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage seafile internal server error response has a 3xx status code
func (o *PostStorageSeafileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage seafile internal server error response has a 4xx status code
func (o *PostStorageSeafileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage seafile internal server error response has a 5xx status code
func (o *PostStorageSeafileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post storage seafile internal server error response a status code equal to that given
func (o *PostStorageSeafileInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post storage seafile internal server error response
func (o *PostStorageSeafileInternalServerError) Code() int {
	return 500
}

func (o *PostStorageSeafileInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/seafile][%d] postStorageSeafileInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageSeafileInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/seafile][%d] postStorageSeafileInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageSeafileInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageSeafileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}