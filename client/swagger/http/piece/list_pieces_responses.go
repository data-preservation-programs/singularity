// Code generated by go-swagger; DO NOT EDIT.

package piece

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// ListPiecesReader is a Reader for the ListPieces structure.
type ListPiecesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPiecesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPiecesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListPiecesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListPiecesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /preparation/{id}/piece] ListPieces", response, response.Code())
	}
}

// NewListPiecesOK creates a ListPiecesOK with default headers values
func NewListPiecesOK() *ListPiecesOK {
	return &ListPiecesOK{}
}

/*
ListPiecesOK describes a response with status code 200, with default header values.

OK
*/
type ListPiecesOK struct {
	Payload []*models.DataprepPieceList
}

// IsSuccess returns true when this list pieces o k response has a 2xx status code
func (o *ListPiecesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list pieces o k response has a 3xx status code
func (o *ListPiecesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list pieces o k response has a 4xx status code
func (o *ListPiecesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list pieces o k response has a 5xx status code
func (o *ListPiecesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list pieces o k response a status code equal to that given
func (o *ListPiecesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list pieces o k response
func (o *ListPiecesOK) Code() int {
	return 200
}

func (o *ListPiecesOK) Error() string {
	return fmt.Sprintf("[GET /preparation/{id}/piece][%d] listPiecesOK  %+v", 200, o.Payload)
}

func (o *ListPiecesOK) String() string {
	return fmt.Sprintf("[GET /preparation/{id}/piece][%d] listPiecesOK  %+v", 200, o.Payload)
}

func (o *ListPiecesOK) GetPayload() []*models.DataprepPieceList {
	return o.Payload
}

func (o *ListPiecesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPiecesBadRequest creates a ListPiecesBadRequest with default headers values
func NewListPiecesBadRequest() *ListPiecesBadRequest {
	return &ListPiecesBadRequest{}
}

/*
ListPiecesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ListPiecesBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this list pieces bad request response has a 2xx status code
func (o *ListPiecesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list pieces bad request response has a 3xx status code
func (o *ListPiecesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list pieces bad request response has a 4xx status code
func (o *ListPiecesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list pieces bad request response has a 5xx status code
func (o *ListPiecesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list pieces bad request response a status code equal to that given
func (o *ListPiecesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list pieces bad request response
func (o *ListPiecesBadRequest) Code() int {
	return 400
}

func (o *ListPiecesBadRequest) Error() string {
	return fmt.Sprintf("[GET /preparation/{id}/piece][%d] listPiecesBadRequest  %+v", 400, o.Payload)
}

func (o *ListPiecesBadRequest) String() string {
	return fmt.Sprintf("[GET /preparation/{id}/piece][%d] listPiecesBadRequest  %+v", 400, o.Payload)
}

func (o *ListPiecesBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *ListPiecesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPiecesInternalServerError creates a ListPiecesInternalServerError with default headers values
func NewListPiecesInternalServerError() *ListPiecesInternalServerError {
	return &ListPiecesInternalServerError{}
}

/*
ListPiecesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ListPiecesInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this list pieces internal server error response has a 2xx status code
func (o *ListPiecesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list pieces internal server error response has a 3xx status code
func (o *ListPiecesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list pieces internal server error response has a 4xx status code
func (o *ListPiecesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list pieces internal server error response has a 5xx status code
func (o *ListPiecesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list pieces internal server error response a status code equal to that given
func (o *ListPiecesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list pieces internal server error response
func (o *ListPiecesInternalServerError) Code() int {
	return 500
}

func (o *ListPiecesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /preparation/{id}/piece][%d] listPiecesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListPiecesInternalServerError) String() string {
	return fmt.Sprintf("[GET /preparation/{id}/piece][%d] listPiecesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListPiecesInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *ListPiecesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
