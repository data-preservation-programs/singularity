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

// PostStorageMegaReader is a Reader for the PostStorageMega structure.
type PostStorageMegaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStorageMegaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStorageMegaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStorageMegaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStorageMegaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/mega] PostStorageMega", response, response.Code())
	}
}

// NewPostStorageMegaOK creates a PostStorageMegaOK with default headers values
func NewPostStorageMegaOK() *PostStorageMegaOK {
	return &PostStorageMegaOK{}
}

/*
PostStorageMegaOK describes a response with status code 200, with default header values.

OK
*/
type PostStorageMegaOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this post storage mega o k response has a 2xx status code
func (o *PostStorageMegaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post storage mega o k response has a 3xx status code
func (o *PostStorageMegaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage mega o k response has a 4xx status code
func (o *PostStorageMegaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage mega o k response has a 5xx status code
func (o *PostStorageMegaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage mega o k response a status code equal to that given
func (o *PostStorageMegaOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post storage mega o k response
func (o *PostStorageMegaOK) Code() int {
	return 200
}

func (o *PostStorageMegaOK) Error() string {
	return fmt.Sprintf("[POST /storage/mega][%d] postStorageMegaOK  %+v", 200, o.Payload)
}

func (o *PostStorageMegaOK) String() string {
	return fmt.Sprintf("[POST /storage/mega][%d] postStorageMegaOK  %+v", 200, o.Payload)
}

func (o *PostStorageMegaOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *PostStorageMegaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageMegaBadRequest creates a PostStorageMegaBadRequest with default headers values
func NewPostStorageMegaBadRequest() *PostStorageMegaBadRequest {
	return &PostStorageMegaBadRequest{}
}

/*
PostStorageMegaBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostStorageMegaBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage mega bad request response has a 2xx status code
func (o *PostStorageMegaBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage mega bad request response has a 3xx status code
func (o *PostStorageMegaBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage mega bad request response has a 4xx status code
func (o *PostStorageMegaBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post storage mega bad request response has a 5xx status code
func (o *PostStorageMegaBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post storage mega bad request response a status code equal to that given
func (o *PostStorageMegaBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post storage mega bad request response
func (o *PostStorageMegaBadRequest) Code() int {
	return 400
}

func (o *PostStorageMegaBadRequest) Error() string {
	return fmt.Sprintf("[POST /storage/mega][%d] postStorageMegaBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageMegaBadRequest) String() string {
	return fmt.Sprintf("[POST /storage/mega][%d] postStorageMegaBadRequest  %+v", 400, o.Payload)
}

func (o *PostStorageMegaBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageMegaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStorageMegaInternalServerError creates a PostStorageMegaInternalServerError with default headers values
func NewPostStorageMegaInternalServerError() *PostStorageMegaInternalServerError {
	return &PostStorageMegaInternalServerError{}
}

/*
PostStorageMegaInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostStorageMegaInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post storage mega internal server error response has a 2xx status code
func (o *PostStorageMegaInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post storage mega internal server error response has a 3xx status code
func (o *PostStorageMegaInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post storage mega internal server error response has a 4xx status code
func (o *PostStorageMegaInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post storage mega internal server error response has a 5xx status code
func (o *PostStorageMegaInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post storage mega internal server error response a status code equal to that given
func (o *PostStorageMegaInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post storage mega internal server error response
func (o *PostStorageMegaInternalServerError) Code() int {
	return 500
}

func (o *PostStorageMegaInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/mega][%d] postStorageMegaInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageMegaInternalServerError) String() string {
	return fmt.Sprintf("[POST /storage/mega][%d] postStorageMegaInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStorageMegaInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostStorageMegaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
