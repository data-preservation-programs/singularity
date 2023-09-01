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

// PostStorageYandexReader is a Reader for the PostStorageYandex structure.
type PostStorageYandexReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageYandexReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStorageYandexOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStorageYandexBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStorageYandexInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/yandex] PostStorageYandex", response, response.Code())
	}
}

// NewPostStorageYandexOK creates a PostStorageYandexOK with default headers values
func NewPostStorageYandexOK() *PostStorageYandexOK {
	return &PostStorageYandexOK{}
}

/*
PostStorageYandexOK describes a response with status code 200, with default header values.

OK
*/
type PostStorageYandexOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this post storage yandex o k response has a 2xx status code
func (o *PostStorageYandexOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post storage yandex o k response has a 3xx status code
func (o *PostStorageYandexOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage yandex o k response has a 4xx status code
func (o *PostStorageYandexOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage yandex o k response has a 5xx status code
func (o *PostStorageYandexOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage yandex o k response a status code equal to that given
func (o *PostStorageYandexOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post storage yandex o k response
func (o *PostStorageYandexOK) Code() int {
	return 200
}

func (o *PostStorageYandexOK) Error() string {
	return fmt.Sprintf("[POST /storage/yandex][%d] postStorageYandexOK  %+v", 200, o.Payload)
}

func (o *PostStorageYandexOK) String() string {
	return fmt.Sprintf("[POST /storage/yandex][%d] postStorageYandexOK  %+v", 200, o.Payload)
}

func (o *PostStorageYandexOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *PostStorageYandexOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageYandexBadRequest creates a PostStorageYandexBadRequest with default headers values
func NewPostStorageYandexBadRequest() *PostStorageYandexBadRequest {
	return &PostStorageYandexBadRequest{}
}

/*
PostStorageYandexBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostStorageYandexBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage yandex bad request response has a 2xx status code
func (o *PostStorageYandexBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage yandex bad request response has a 3xx status code
func (o *PostStorageYandexBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage yandex bad request response has a 4xx status code
func (o *PostStorageYandexBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post storage yandex bad request response has a 5xx status code
func (o *PostStorageYandexBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage yandex bad request response a status code equal to that given
func (o *PostStorageYandexBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post storage yandex bad request response
func (o *PostStorageYandexBadRequest) Code() int {
	return 400
}

func (o *PostStorageYandexBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/yandex][%d] postStorageYandexBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageYandexBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/yandex][%d] postStorageYandexBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageYandexBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageYandexBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageYandexInternalServerError creates a PostStorageYandexInternalServerError with default headers values
func NewPostStorageYandexInternalServerError() *PostStorageYandexInternalServerError {
	return &PostStorageYandexInternalServerError{}
}

/*
PostStorageYandexInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostStorageYandexInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage yandex internal server error response has a 2xx status code
func (o *PostStorageYandexInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage yandex internal server error response has a 3xx status code
func (o *PostStorageYandexInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage yandex internal server error response has a 4xx status code
func (o *PostStorageYandexInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage yandex internal server error response has a 5xx status code
func (o *PostStorageYandexInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post storage yandex internal server error response a status code equal to that given
func (o *PostStorageYandexInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post storage yandex internal server error response
func (o *PostStorageYandexInternalServerError) Code() int {
	return 500
}

func (o *PostStorageYandexInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/yandex][%d] postStorageYandexInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageYandexInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/yandex][%d] postStorageYandexInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageYandexInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageYandexInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}