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

// PostStorageKoofrOtherReader is a Reader for the PostStorageKoofrOther structure.
type PostStorageKoofrOtherReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageKoofrOtherReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStorageKoofrOtherOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStorageKoofrOtherBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStorageKoofrOtherInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/koofr/other] PostStorageKoofrOther", response, response.Code())
	}
}

// NewPostStorageKoofrOtherOK creates a PostStorageKoofrOtherOK with default headers values
func NewPostStorageKoofrOtherOK() *PostStorageKoofrOtherOK {
	return &PostStorageKoofrOtherOK{}
}

/*
PostStorageKoofrOtherOK describes a response with status code 200, with default header values.

OK
*/
type PostStorageKoofrOtherOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this post storage koofr other o k response has a 2xx status code
func (o *PostStorageKoofrOtherOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post storage koofr other o k response has a 3xx status code
func (o *PostStorageKoofrOtherOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage koofr other o k response has a 4xx status code
func (o *PostStorageKoofrOtherOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage koofr other o k response has a 5xx status code
func (o *PostStorageKoofrOtherOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage koofr other o k response a status code equal to that given
func (o *PostStorageKoofrOtherOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post storage koofr other o k response
func (o *PostStorageKoofrOtherOK) Code() int {
	return 200
}

func (o *PostStorageKoofrOtherOK) Error() string {
	return fmt.Sprintf("[POST /storage/koofr/other][%d] postStorageKoofrOtherOK  %+v", 200, o.Payload)
}

func (o *PostStorageKoofrOtherOK) String() string {
	return fmt.Sprintf("[POST /storage/koofr/other][%d] postStorageKoofrOtherOK  %+v", 200, o.Payload)
}

func (o *PostStorageKoofrOtherOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *PostStorageKoofrOtherOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageKoofrOtherBadRequest creates a PostStorageKoofrOtherBadRequest with default headers values
func NewPostStorageKoofrOtherBadRequest() *PostStorageKoofrOtherBadRequest {
	return &PostStorageKoofrOtherBadRequest{}
}

/*
PostStorageKoofrOtherBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostStorageKoofrOtherBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage koofr other bad request response has a 2xx status code
func (o *PostStorageKoofrOtherBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage koofr other bad request response has a 3xx status code
func (o *PostStorageKoofrOtherBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage koofr other bad request response has a 4xx status code
func (o *PostStorageKoofrOtherBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post storage koofr other bad request response has a 5xx status code
func (o *PostStorageKoofrOtherBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage koofr other bad request response a status code equal to that given
func (o *PostStorageKoofrOtherBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post storage koofr other bad request response
func (o *PostStorageKoofrOtherBadRequest) Code() int {
	return 400
}

func (o *PostStorageKoofrOtherBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/koofr/other][%d] postStorageKoofrOtherBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageKoofrOtherBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/koofr/other][%d] postStorageKoofrOtherBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageKoofrOtherBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageKoofrOtherBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageKoofrOtherInternalServerError creates a PostStorageKoofrOtherInternalServerError with default headers values
func NewPostStorageKoofrOtherInternalServerError() *PostStorageKoofrOtherInternalServerError {
	return &PostStorageKoofrOtherInternalServerError{}
}

/*
PostStorageKoofrOtherInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostStorageKoofrOtherInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage koofr other internal server error response has a 2xx status code
func (o *PostStorageKoofrOtherInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage koofr other internal server error response has a 3xx status code
func (o *PostStorageKoofrOtherInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage koofr other internal server error response has a 4xx status code
func (o *PostStorageKoofrOtherInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage koofr other internal server error response has a 5xx status code
func (o *PostStorageKoofrOtherInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post storage koofr other internal server error response a status code equal to that given
func (o *PostStorageKoofrOtherInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post storage koofr other internal server error response
func (o *PostStorageKoofrOtherInternalServerError) Code() int {
	return 500
}

func (o *PostStorageKoofrOtherInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/koofr/other][%d] postStorageKoofrOtherInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageKoofrOtherInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/koofr/other][%d] postStorageKoofrOtherInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageKoofrOtherInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageKoofrOtherInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
