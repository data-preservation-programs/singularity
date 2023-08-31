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

// PostStorageJottacloudReader is a Reader for the PostStorageJottacloud structure.
type PostStorageJottacloudReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageJottacloudReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStorageJottacloudOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStorageJottacloudBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStorageJottacloudInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/jottacloud] PostStorageJottacloud", response, response.Code())
	}
}

// NewPostStorageJottacloudOK creates a PostStorageJottacloudOK with default headers values
func NewPostStorageJottacloudOK() *PostStorageJottacloudOK {
	return &PostStorageJottacloudOK{}
}

/*
PostStorageJottacloudOK describes a response with status code 200, with default header values.

OK
*/
type PostStorageJottacloudOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this post storage jottacloud o k response has a 2xx status code
func (o *PostStorageJottacloudOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post storage jottacloud o k response has a 3xx status code
func (o *PostStorageJottacloudOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage jottacloud o k response has a 4xx status code
func (o *PostStorageJottacloudOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage jottacloud o k response has a 5xx status code
func (o *PostStorageJottacloudOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage jottacloud o k response a status code equal to that given
func (o *PostStorageJottacloudOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post storage jottacloud o k response
func (o *PostStorageJottacloudOK) Code() int {
	return 200
}

func (o *PostStorageJottacloudOK) Error() string {
	return fmt.Sprintf("[POST /storage/jottacloud][%d] postStorageJottacloudOK  %+v", 200, o.Payload)
}

func (o *PostStorageJottacloudOK) String() string {
	return fmt.Sprintf("[POST /storage/jottacloud][%d] postStorageJottacloudOK  %+v", 200, o.Payload)
}

func (o *PostStorageJottacloudOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *PostStorageJottacloudOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageJottacloudBadRequest creates a PostStorageJottacloudBadRequest with default headers values
func NewPostStorageJottacloudBadRequest() *PostStorageJottacloudBadRequest {
	return &PostStorageJottacloudBadRequest{}
}

/*
PostStorageJottacloudBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostStorageJottacloudBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage jottacloud bad request response has a 2xx status code
func (o *PostStorageJottacloudBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage jottacloud bad request response has a 3xx status code
func (o *PostStorageJottacloudBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage jottacloud bad request response has a 4xx status code
func (o *PostStorageJottacloudBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post storage jottacloud bad request response has a 5xx status code
func (o *PostStorageJottacloudBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage jottacloud bad request response a status code equal to that given
func (o *PostStorageJottacloudBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post storage jottacloud bad request response
func (o *PostStorageJottacloudBadRequest) Code() int {
	return 400
}

func (o *PostStorageJottacloudBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/jottacloud][%d] postStorageJottacloudBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageJottacloudBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/jottacloud][%d] postStorageJottacloudBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageJottacloudBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageJottacloudBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageJottacloudInternalServerError creates a PostStorageJottacloudInternalServerError with default headers values
func NewPostStorageJottacloudInternalServerError() *PostStorageJottacloudInternalServerError {
	return &PostStorageJottacloudInternalServerError{}
}

/*
PostStorageJottacloudInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostStorageJottacloudInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage jottacloud internal server error response has a 2xx status code
func (o *PostStorageJottacloudInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage jottacloud internal server error response has a 3xx status code
func (o *PostStorageJottacloudInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage jottacloud internal server error response has a 4xx status code
func (o *PostStorageJottacloudInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage jottacloud internal server error response has a 5xx status code
func (o *PostStorageJottacloudInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post storage jottacloud internal server error response a status code equal to that given
func (o *PostStorageJottacloudInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post storage jottacloud internal server error response
func (o *PostStorageJottacloudInternalServerError) Code() int {
	return 500
}

func (o *PostStorageJottacloudInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/jottacloud][%d] postStorageJottacloudInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageJottacloudInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/jottacloud][%d] postStorageJottacloudInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageJottacloudInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageJottacloudInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
