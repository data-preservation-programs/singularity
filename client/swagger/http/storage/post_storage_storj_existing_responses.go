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

// PostStorageStorjExistingReader is a Reader for the PostStorageStorjExisting structure.
type PostStorageStorjExistingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageStorjExistingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStorageStorjExistingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStorageStorjExistingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStorageStorjExistingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/storj/existing] PostStorageStorjExisting", response, response.Code())
	}
}

// NewPostStorageStorjExistingOK creates a PostStorageStorjExistingOK with default headers values
func NewPostStorageStorjExistingOK() *PostStorageStorjExistingOK {
	return &PostStorageStorjExistingOK{}
}

/*
PostStorageStorjExistingOK describes a response with status code 200, with default header values.

OK
*/
type PostStorageStorjExistingOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this post storage storj existing o k response has a 2xx status code
func (o *PostStorageStorjExistingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post storage storj existing o k response has a 3xx status code
func (o *PostStorageStorjExistingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage storj existing o k response has a 4xx status code
func (o *PostStorageStorjExistingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage storj existing o k response has a 5xx status code
func (o *PostStorageStorjExistingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage storj existing o k response a status code equal to that given
func (o *PostStorageStorjExistingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post storage storj existing o k response
func (o *PostStorageStorjExistingOK) Code() int {
	return 200
}

func (o *PostStorageStorjExistingOK) Error() string {
	return fmt.Sprintf("[POST /storage/storj/existing][%d] postStorageStorjExistingOK  %+v", 200, o.Payload)
}

func (o *PostStorageStorjExistingOK) String() string {
	return fmt.Sprintf("[POST /storage/storj/existing][%d] postStorageStorjExistingOK  %+v", 200, o.Payload)
}

func (o *PostStorageStorjExistingOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *PostStorageStorjExistingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageStorjExistingBadRequest creates a PostStorageStorjExistingBadRequest with default headers values
func NewPostStorageStorjExistingBadRequest() *PostStorageStorjExistingBadRequest {
	return &PostStorageStorjExistingBadRequest{}
}

/*
PostStorageStorjExistingBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostStorageStorjExistingBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage storj existing bad request response has a 2xx status code
func (o *PostStorageStorjExistingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage storj existing bad request response has a 3xx status code
func (o *PostStorageStorjExistingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage storj existing bad request response has a 4xx status code
func (o *PostStorageStorjExistingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post storage storj existing bad request response has a 5xx status code
func (o *PostStorageStorjExistingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage storj existing bad request response a status code equal to that given
func (o *PostStorageStorjExistingBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post storage storj existing bad request response
func (o *PostStorageStorjExistingBadRequest) Code() int {
	return 400
}

func (o *PostStorageStorjExistingBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/storj/existing][%d] postStorageStorjExistingBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageStorjExistingBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/storj/existing][%d] postStorageStorjExistingBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageStorjExistingBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageStorjExistingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageStorjExistingInternalServerError creates a PostStorageStorjExistingInternalServerError with default headers values
func NewPostStorageStorjExistingInternalServerError() *PostStorageStorjExistingInternalServerError {
	return &PostStorageStorjExistingInternalServerError{}
}

/*
PostStorageStorjExistingInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostStorageStorjExistingInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage storj existing internal server error response has a 2xx status code
func (o *PostStorageStorjExistingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage storj existing internal server error response has a 3xx status code
func (o *PostStorageStorjExistingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage storj existing internal server error response has a 4xx status code
func (o *PostStorageStorjExistingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage storj existing internal server error response has a 5xx status code
func (o *PostStorageStorjExistingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post storage storj existing internal server error response a status code equal to that given
func (o *PostStorageStorjExistingInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post storage storj existing internal server error response
func (o *PostStorageStorjExistingInternalServerError) Code() int {
	return 500
}

func (o *PostStorageStorjExistingInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/storj/existing][%d] postStorageStorjExistingInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageStorjExistingInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/storj/existing][%d] postStorageStorjExistingInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageStorjExistingInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageStorjExistingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}