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

// DeleteStorageNameReader is a Reader for the DeleteStorageName structure.
type DeleteStorageNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteStorageNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteStorageNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteStorageNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteStorageNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /storage/{name}] DeleteStorageName", response, response.Code())
	}
}

// NewDeleteStorageNameNoContent creates a DeleteStorageNameNoContent with default headers values
func NewDeleteStorageNameNoContent() *DeleteStorageNameNoContent {
	return &DeleteStorageNameNoContent{}
}

/*
DeleteStorageNameNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteStorageNameNoContent struct {
}

// IsSuccess returns true when this delete storage name no content response has a 2xx status code
func (o *DeleteStorageNameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete storage name no content response has a 3xx status code
func (o *DeleteStorageNameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete storage name no content response has a 4xx status code
func (o *DeleteStorageNameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete storage name no content response has a 5xx status code
func (o *DeleteStorageNameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete storage name no content response a status code equal to that given
func (o *DeleteStorageNameNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete storage name no content response
func (o *DeleteStorageNameNoContent) Code() int {
	return 204
}

func (o *DeleteStorageNameNoContent) Error() string {
	return fmt.Sprintf("[DELETE /storage/{name}][%d] deleteStorageNameNoContent ", 204)
}

func (o *DeleteStorageNameNoContent) String() string {
	return fmt.Sprintf("[DELETE /storage/{name}][%d] deleteStorageNameNoContent ", 204)
}

func (o *DeleteStorageNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteStorageNameBadRequest creates a DeleteStorageNameBadRequest with default headers values
func NewDeleteStorageNameBadRequest() *DeleteStorageNameBadRequest {
	return &DeleteStorageNameBadRequest{}
}

/*
DeleteStorageNameBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteStorageNameBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this delete storage name bad request response has a 2xx status code
func (o *DeleteStorageNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete storage name bad request response has a 3xx status code
func (o *DeleteStorageNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete storage name bad request response has a 4xx status code
func (o *DeleteStorageNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete storage name bad request response has a 5xx status code
func (o *DeleteStorageNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete storage name bad request response a status code equal to that given
func (o *DeleteStorageNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete storage name bad request response
func (o *DeleteStorageNameBadRequest) Code() int {
	return 400
}

func (o *DeleteStorageNameBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /storage/{name}][%d] deleteStorageNameBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteStorageNameBadRequest) String() string {
	return fmt.Sprintf("[DELETE /storage/{name}][%d] deleteStorageNameBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteStorageNameBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *DeleteStorageNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteStorageNameInternalServerError creates a DeleteStorageNameInternalServerError with default headers values
func NewDeleteStorageNameInternalServerError() *DeleteStorageNameInternalServerError {
	return &DeleteStorageNameInternalServerError{}
}

/*
DeleteStorageNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteStorageNameInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this delete storage name internal server error response has a 2xx status code
func (o *DeleteStorageNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete storage name internal server error response has a 3xx status code
func (o *DeleteStorageNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete storage name internal server error response has a 4xx status code
func (o *DeleteStorageNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete storage name internal server error response has a 5xx status code
func (o *DeleteStorageNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete storage name internal server error response a status code equal to that given
func (o *DeleteStorageNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete storage name internal server error response
func (o *DeleteStorageNameInternalServerError) Code() int {
	return 500
}

func (o *DeleteStorageNameInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /storage/{name}][%d] deleteStorageNameInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteStorageNameInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /storage/{name}][%d] deleteStorageNameInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteStorageNameInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *DeleteStorageNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}