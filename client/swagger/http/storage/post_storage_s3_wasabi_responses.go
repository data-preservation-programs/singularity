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

// PostStorageS3WasabiReader is a Reader for the PostStorageS3Wasabi structure.
type PostStorageS3WasabiReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageS3WasabiReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStorageS3WasabiOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStorageS3WasabiBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStorageS3WasabiInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/s3/wasabi] PostStorageS3Wasabi", response, response.Code())
	}
}

// NewPostStorageS3WasabiOK creates a PostStorageS3WasabiOK with default headers values
func NewPostStorageS3WasabiOK() *PostStorageS3WasabiOK {
	return &PostStorageS3WasabiOK{}
}

/*
PostStorageS3WasabiOK describes a response with status code 200, with default header values.

OK
*/
type PostStorageS3WasabiOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this post storage s3 wasabi o k response has a 2xx status code
func (o *PostStorageS3WasabiOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post storage s3 wasabi o k response has a 3xx status code
func (o *PostStorageS3WasabiOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage s3 wasabi o k response has a 4xx status code
func (o *PostStorageS3WasabiOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage s3 wasabi o k response has a 5xx status code
func (o *PostStorageS3WasabiOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage s3 wasabi o k response a status code equal to that given
func (o *PostStorageS3WasabiOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post storage s3 wasabi o k response
func (o *PostStorageS3WasabiOK) Code() int {
	return 200
}

func (o *PostStorageS3WasabiOK) Error() string {
	return fmt.Sprintf("[POST /storage/s3/wasabi][%d] postStorageS3WasabiOK  %+v", 200, o.Payload)
}

func (o *PostStorageS3WasabiOK) String() string {
	return fmt.Sprintf("[POST /storage/s3/wasabi][%d] postStorageS3WasabiOK  %+v", 200, o.Payload)
}

func (o *PostStorageS3WasabiOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *PostStorageS3WasabiOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageS3WasabiBadRequest creates a PostStorageS3WasabiBadRequest with default headers values
func NewPostStorageS3WasabiBadRequest() *PostStorageS3WasabiBadRequest {
	return &PostStorageS3WasabiBadRequest{}
}

/*
PostStorageS3WasabiBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostStorageS3WasabiBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage s3 wasabi bad request response has a 2xx status code
func (o *PostStorageS3WasabiBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage s3 wasabi bad request response has a 3xx status code
func (o *PostStorageS3WasabiBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage s3 wasabi bad request response has a 4xx status code
func (o *PostStorageS3WasabiBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post storage s3 wasabi bad request response has a 5xx status code
func (o *PostStorageS3WasabiBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage s3 wasabi bad request response a status code equal to that given
func (o *PostStorageS3WasabiBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post storage s3 wasabi bad request response
func (o *PostStorageS3WasabiBadRequest) Code() int {
	return 400
}

func (o *PostStorageS3WasabiBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/s3/wasabi][%d] postStorageS3WasabiBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageS3WasabiBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/s3/wasabi][%d] postStorageS3WasabiBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageS3WasabiBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageS3WasabiBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageS3WasabiInternalServerError creates a PostStorageS3WasabiInternalServerError with default headers values
func NewPostStorageS3WasabiInternalServerError() *PostStorageS3WasabiInternalServerError {
	return &PostStorageS3WasabiInternalServerError{}
}

/*
PostStorageS3WasabiInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostStorageS3WasabiInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage s3 wasabi internal server error response has a 2xx status code
func (o *PostStorageS3WasabiInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage s3 wasabi internal server error response has a 3xx status code
func (o *PostStorageS3WasabiInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage s3 wasabi internal server error response has a 4xx status code
func (o *PostStorageS3WasabiInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage s3 wasabi internal server error response has a 5xx status code
func (o *PostStorageS3WasabiInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post storage s3 wasabi internal server error response a status code equal to that given
func (o *PostStorageS3WasabiInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post storage s3 wasabi internal server error response
func (o *PostStorageS3WasabiInternalServerError) Code() int {
	return 500
}

func (o *PostStorageS3WasabiInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/s3/wasabi][%d] postStorageS3WasabiInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageS3WasabiInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/s3/wasabi][%d] postStorageS3WasabiInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageS3WasabiInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageS3WasabiInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
