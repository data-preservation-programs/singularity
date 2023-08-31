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

// PostStorageOnedriveReader is a Reader for the PostStorageOnedrive structure.
type PostStorageOnedriveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageOnedriveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStorageOnedriveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStorageOnedriveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStorageOnedriveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/onedrive] PostStorageOnedrive", response, response.Code())
	}
}

// NewPostStorageOnedriveOK creates a PostStorageOnedriveOK with default headers values
func NewPostStorageOnedriveOK() *PostStorageOnedriveOK {
	return &PostStorageOnedriveOK{}
}

/*
PostStorageOnedriveOK describes a response with status code 200, with default header values.

OK
*/
type PostStorageOnedriveOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this post storage onedrive o k response has a 2xx status code
func (o *PostStorageOnedriveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post storage onedrive o k response has a 3xx status code
func (o *PostStorageOnedriveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage onedrive o k response has a 4xx status code
func (o *PostStorageOnedriveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage onedrive o k response has a 5xx status code
func (o *PostStorageOnedriveOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage onedrive o k response a status code equal to that given
func (o *PostStorageOnedriveOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post storage onedrive o k response
func (o *PostStorageOnedriveOK) Code() int {
	return 200
}

func (o *PostStorageOnedriveOK) Error() string {
	return fmt.Sprintf("[POST /storage/onedrive][%d] postStorageOnedriveOK  %+v", 200, o.Payload)
}

func (o *PostStorageOnedriveOK) String() string {
	return fmt.Sprintf("[POST /storage/onedrive][%d] postStorageOnedriveOK  %+v", 200, o.Payload)
}

func (o *PostStorageOnedriveOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *PostStorageOnedriveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageOnedriveBadRequest creates a PostStorageOnedriveBadRequest with default headers values
func NewPostStorageOnedriveBadRequest() *PostStorageOnedriveBadRequest {
	return &PostStorageOnedriveBadRequest{}
}

/*
PostStorageOnedriveBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostStorageOnedriveBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage onedrive bad request response has a 2xx status code
func (o *PostStorageOnedriveBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage onedrive bad request response has a 3xx status code
func (o *PostStorageOnedriveBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage onedrive bad request response has a 4xx status code
func (o *PostStorageOnedriveBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post storage onedrive bad request response has a 5xx status code
func (o *PostStorageOnedriveBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage onedrive bad request response a status code equal to that given
func (o *PostStorageOnedriveBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post storage onedrive bad request response
func (o *PostStorageOnedriveBadRequest) Code() int {
	return 400
}

func (o *PostStorageOnedriveBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/onedrive][%d] postStorageOnedriveBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageOnedriveBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/onedrive][%d] postStorageOnedriveBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageOnedriveBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageOnedriveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageOnedriveInternalServerError creates a PostStorageOnedriveInternalServerError with default headers values
func NewPostStorageOnedriveInternalServerError() *PostStorageOnedriveInternalServerError {
	return &PostStorageOnedriveInternalServerError{}
}

/*
PostStorageOnedriveInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostStorageOnedriveInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage onedrive internal server error response has a 2xx status code
func (o *PostStorageOnedriveInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage onedrive internal server error response has a 3xx status code
func (o *PostStorageOnedriveInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage onedrive internal server error response has a 4xx status code
func (o *PostStorageOnedriveInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage onedrive internal server error response has a 5xx status code
func (o *PostStorageOnedriveInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post storage onedrive internal server error response a status code equal to that given
func (o *PostStorageOnedriveInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post storage onedrive internal server error response
func (o *PostStorageOnedriveInternalServerError) Code() int {
	return 500
}

func (o *PostStorageOnedriveInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/onedrive][%d] postStorageOnedriveInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageOnedriveInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/onedrive][%d] postStorageOnedriveInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageOnedriveInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageOnedriveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
