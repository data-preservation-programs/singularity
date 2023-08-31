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

// PostStorageSftpReader is a Reader for the PostStorageSftp structure.
type PostStorageSftpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageSftpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStorageSftpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStorageSftpBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStorageSftpInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/sftp] PostStorageSftp", response, response.Code())
	}
}

// NewPostStorageSftpOK creates a PostStorageSftpOK with default headers values
func NewPostStorageSftpOK() *PostStorageSftpOK {
	return &PostStorageSftpOK{}
}

/*
PostStorageSftpOK describes a response with status code 200, with default header values.

OK
*/
type PostStorageSftpOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this post storage sftp o k response has a 2xx status code
func (o *PostStorageSftpOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post storage sftp o k response has a 3xx status code
func (o *PostStorageSftpOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage sftp o k response has a 4xx status code
func (o *PostStorageSftpOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage sftp o k response has a 5xx status code
func (o *PostStorageSftpOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage sftp o k response a status code equal to that given
func (o *PostStorageSftpOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post storage sftp o k response
func (o *PostStorageSftpOK) Code() int {
	return 200
}

func (o *PostStorageSftpOK) Error() string {
	return fmt.Sprintf("[POST /storage/sftp][%d] postStorageSftpOK  %+v", 200, o.Payload)
}

func (o *PostStorageSftpOK) String() string {
	return fmt.Sprintf("[POST /storage/sftp][%d] postStorageSftpOK  %+v", 200, o.Payload)
}

func (o *PostStorageSftpOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *PostStorageSftpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageSftpBadRequest creates a PostStorageSftpBadRequest with default headers values
func NewPostStorageSftpBadRequest() *PostStorageSftpBadRequest {
	return &PostStorageSftpBadRequest{}
}

/*
PostStorageSftpBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostStorageSftpBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage sftp bad request response has a 2xx status code
func (o *PostStorageSftpBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage sftp bad request response has a 3xx status code
func (o *PostStorageSftpBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage sftp bad request response has a 4xx status code
func (o *PostStorageSftpBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post storage sftp bad request response has a 5xx status code
func (o *PostStorageSftpBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage sftp bad request response a status code equal to that given
func (o *PostStorageSftpBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post storage sftp bad request response
func (o *PostStorageSftpBadRequest) Code() int {
	return 400
}

func (o *PostStorageSftpBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/sftp][%d] postStorageSftpBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageSftpBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/sftp][%d] postStorageSftpBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageSftpBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageSftpBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageSftpInternalServerError creates a PostStorageSftpInternalServerError with default headers values
func NewPostStorageSftpInternalServerError() *PostStorageSftpInternalServerError {
	return &PostStorageSftpInternalServerError{}
}

/*
PostStorageSftpInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostStorageSftpInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage sftp internal server error response has a 2xx status code
func (o *PostStorageSftpInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage sftp internal server error response has a 3xx status code
func (o *PostStorageSftpInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage sftp internal server error response has a 4xx status code
func (o *PostStorageSftpInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage sftp internal server error response has a 5xx status code
func (o *PostStorageSftpInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post storage sftp internal server error response a status code equal to that given
func (o *PostStorageSftpInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post storage sftp internal server error response
func (o *PostStorageSftpInternalServerError) Code() int {
	return 500
}

func (o *PostStorageSftpInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/sftp][%d] postStorageSftpInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageSftpInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/sftp][%d] postStorageSftpInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageSftpInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageSftpInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
