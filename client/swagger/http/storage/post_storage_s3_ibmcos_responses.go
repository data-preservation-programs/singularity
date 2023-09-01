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

// PostStorageS3IbmcosReader is a Reader for the PostStorageS3Ibmcos structure.
type PostStorageS3IbmcosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageS3IbmcosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStorageS3IbmcosOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStorageS3IbmcosBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStorageS3IbmcosInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/s3/ibmcos] PostStorageS3Ibmcos", response, response.Code())
	}
}

// NewPostStorageS3IbmcosOK creates a PostStorageS3IbmcosOK with default headers values
func NewPostStorageS3IbmcosOK() *PostStorageS3IbmcosOK {
	return &PostStorageS3IbmcosOK{}
}

/*
PostStorageS3IbmcosOK describes a response with status code 200, with default header values.

OK
*/
type PostStorageS3IbmcosOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this post storage s3 ibmcos o k response has a 2xx status code
func (o *PostStorageS3IbmcosOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post storage s3 ibmcos o k response has a 3xx status code
func (o *PostStorageS3IbmcosOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage s3 ibmcos o k response has a 4xx status code
func (o *PostStorageS3IbmcosOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage s3 ibmcos o k response has a 5xx status code
func (o *PostStorageS3IbmcosOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage s3 ibmcos o k response a status code equal to that given
func (o *PostStorageS3IbmcosOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post storage s3 ibmcos o k response
func (o *PostStorageS3IbmcosOK) Code() int {
	return 200
}

func (o *PostStorageS3IbmcosOK) Error() string {
	return fmt.Sprintf("[POST /storage/s3/ibmcos][%d] postStorageS3IbmcosOK  %+v", 200, o.Payload)
}

func (o *PostStorageS3IbmcosOK) String() string {
	return fmt.Sprintf("[POST /storage/s3/ibmcos][%d] postStorageS3IbmcosOK  %+v", 200, o.Payload)
}

func (o *PostStorageS3IbmcosOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *PostStorageS3IbmcosOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageS3IbmcosBadRequest creates a PostStorageS3IbmcosBadRequest with default headers values
func NewPostStorageS3IbmcosBadRequest() *PostStorageS3IbmcosBadRequest {
	return &PostStorageS3IbmcosBadRequest{}
}

/*
PostStorageS3IbmcosBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostStorageS3IbmcosBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage s3 ibmcos bad request response has a 2xx status code
func (o *PostStorageS3IbmcosBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage s3 ibmcos bad request response has a 3xx status code
func (o *PostStorageS3IbmcosBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage s3 ibmcos bad request response has a 4xx status code
func (o *PostStorageS3IbmcosBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post storage s3 ibmcos bad request response has a 5xx status code
func (o *PostStorageS3IbmcosBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage s3 ibmcos bad request response a status code equal to that given
func (o *PostStorageS3IbmcosBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post storage s3 ibmcos bad request response
func (o *PostStorageS3IbmcosBadRequest) Code() int {
	return 400
}

func (o *PostStorageS3IbmcosBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/s3/ibmcos][%d] postStorageS3IbmcosBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageS3IbmcosBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/s3/ibmcos][%d] postStorageS3IbmcosBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageS3IbmcosBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageS3IbmcosBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageS3IbmcosInternalServerError creates a PostStorageS3IbmcosInternalServerError with default headers values
func NewPostStorageS3IbmcosInternalServerError() *PostStorageS3IbmcosInternalServerError {
	return &PostStorageS3IbmcosInternalServerError{}
}

/*
PostStorageS3IbmcosInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostStorageS3IbmcosInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage s3 ibmcos internal server error response has a 2xx status code
func (o *PostStorageS3IbmcosInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage s3 ibmcos internal server error response has a 3xx status code
func (o *PostStorageS3IbmcosInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage s3 ibmcos internal server error response has a 4xx status code
func (o *PostStorageS3IbmcosInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage s3 ibmcos internal server error response has a 5xx status code
func (o *PostStorageS3IbmcosInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post storage s3 ibmcos internal server error response a status code equal to that given
func (o *PostStorageS3IbmcosInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post storage s3 ibmcos internal server error response
func (o *PostStorageS3IbmcosInternalServerError) Code() int {
	return 500
}

func (o *PostStorageS3IbmcosInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/s3/ibmcos][%d] postStorageS3IbmcosInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageS3IbmcosInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/s3/ibmcos][%d] postStorageS3IbmcosInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageS3IbmcosInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageS3IbmcosInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}