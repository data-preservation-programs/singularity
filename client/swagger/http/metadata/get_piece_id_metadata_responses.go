// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// GetPieceIDMetadataReader is a Reader for the GetPieceIDMetadata structure.
type GetPieceIDMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPieceIDMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPieceIDMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPieceIDMetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPieceIDMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPieceIDMetadataInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /piece/{id}/metadata] GetPieceIDMetadata", response, response.Code())
	}
}

// NewGetPieceIDMetadataOK creates a GetPieceIDMetadataOK with default headers values
func NewGetPieceIDMetadataOK() *GetPieceIDMetadataOK {
	return &GetPieceIDMetadataOK{}
}

/*
GetPieceIDMetadataOK describes a response with status code 200, with default header values.

OK
*/
type GetPieceIDMetadataOK struct {
	Payload models.StorePieceReader
}

// IsSuccess returns true when this get piece Id metadata o k response has a 2xx status code
func (o *GetPieceIDMetadataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get piece Id metadata o k response has a 3xx status code
func (o *GetPieceIDMetadataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get piece Id metadata o k response has a 4xx status code
func (o *GetPieceIDMetadataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get piece Id metadata o k response has a 5xx status code
func (o *GetPieceIDMetadataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get piece Id metadata o k response a status code equal to that given
func (o *GetPieceIDMetadataOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get piece Id metadata o k response
func (o *GetPieceIDMetadataOK) Code() int {
	return 200
}

func (o *GetPieceIDMetadataOK) Error() string {
	return fmt.Sprintf("[GET /piece/{id}/metadata][%d] getPieceIdMetadataOK  %+v", 200, o.Payload)
}

func (o *GetPieceIDMetadataOK) String() string {
	return fmt.Sprintf("[GET /piece/{id}/metadata][%d] getPieceIdMetadataOK  %+v", 200, o.Payload)
}

func (o *GetPieceIDMetadataOK) GetPayload() models.StorePieceReader {
	return o.Payload
}

func (o *GetPieceIDMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPieceIDMetadataBadRequest creates a GetPieceIDMetadataBadRequest with default headers values
func NewGetPieceIDMetadataBadRequest() *GetPieceIDMetadataBadRequest {
	return &GetPieceIDMetadataBadRequest{}
}

/*
GetPieceIDMetadataBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetPieceIDMetadataBadRequest struct {
	Payload string
}

// IsSuccess returns true when this get piece Id metadata bad request response has a 2xx status code
func (o *GetPieceIDMetadataBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get piece Id metadata bad request response has a 3xx status code
func (o *GetPieceIDMetadataBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get piece Id metadata bad request response has a 4xx status code
func (o *GetPieceIDMetadataBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get piece Id metadata bad request response has a 5xx status code
func (o *GetPieceIDMetadataBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get piece Id metadata bad request response a status code equal to that given
func (o *GetPieceIDMetadataBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get piece Id metadata bad request response
func (o *GetPieceIDMetadataBadRequest) Code() int {
	return 400
}

func (o *GetPieceIDMetadataBadRequest) Error() string {
	return fmt.Sprintf("[GET /piece/{id}/metadata][%d] getPieceIdMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *GetPieceIDMetadataBadRequest) String() string {
	return fmt.Sprintf("[GET /piece/{id}/metadata][%d] getPieceIdMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *GetPieceIDMetadataBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetPieceIDMetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPieceIDMetadataNotFound creates a GetPieceIDMetadataNotFound with default headers values
func NewGetPieceIDMetadataNotFound() *GetPieceIDMetadataNotFound {
	return &GetPieceIDMetadataNotFound{}
}

/*
GetPieceIDMetadataNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetPieceIDMetadataNotFound struct {
	Payload string
}

// IsSuccess returns true when this get piece Id metadata not found response has a 2xx status code
func (o *GetPieceIDMetadataNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get piece Id metadata not found response has a 3xx status code
func (o *GetPieceIDMetadataNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get piece Id metadata not found response has a 4xx status code
func (o *GetPieceIDMetadataNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get piece Id metadata not found response has a 5xx status code
func (o *GetPieceIDMetadataNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get piece Id metadata not found response a status code equal to that given
func (o *GetPieceIDMetadataNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get piece Id metadata not found response
func (o *GetPieceIDMetadataNotFound) Code() int {
	return 404
}

func (o *GetPieceIDMetadataNotFound) Error() string {
	return fmt.Sprintf("[GET /piece/{id}/metadata][%d] getPieceIdMetadataNotFound  %+v", 404, o.Payload)
}

func (o *GetPieceIDMetadataNotFound) String() string {
	return fmt.Sprintf("[GET /piece/{id}/metadata][%d] getPieceIdMetadataNotFound  %+v", 404, o.Payload)
}

func (o *GetPieceIDMetadataNotFound) GetPayload() string {
	return o.Payload
}

func (o *GetPieceIDMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPieceIDMetadataInternalServerError creates a GetPieceIDMetadataInternalServerError with default headers values
func NewGetPieceIDMetadataInternalServerError() *GetPieceIDMetadataInternalServerError {
	return &GetPieceIDMetadataInternalServerError{}
}

/*
GetPieceIDMetadataInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetPieceIDMetadataInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this get piece Id metadata internal server error response has a 2xx status code
func (o *GetPieceIDMetadataInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get piece Id metadata internal server error response has a 3xx status code
func (o *GetPieceIDMetadataInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get piece Id metadata internal server error response has a 4xx status code
func (o *GetPieceIDMetadataInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get piece Id metadata internal server error response has a 5xx status code
func (o *GetPieceIDMetadataInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get piece Id metadata internal server error response a status code equal to that given
func (o *GetPieceIDMetadataInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get piece Id metadata internal server error response
func (o *GetPieceIDMetadataInternalServerError) Code() int {
	return 500
}

func (o *GetPieceIDMetadataInternalServerError) Error() string {
	return fmt.Sprintf("[GET /piece/{id}/metadata][%d] getPieceIdMetadataInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPieceIDMetadataInternalServerError) String() string {
	return fmt.Sprintf("[GET /piece/{id}/metadata][%d] getPieceIdMetadataInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPieceIDMetadataInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *GetPieceIDMetadataInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}