// Code generated by go-swagger; DO NOT EDIT.

package data_source

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// GetSourceIDFilesReader is a Reader for the GetSourceIDFiles structure.
type GetSourceIDFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSourceIDFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSourceIDFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSourceIDFilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSourceIDFilesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /source/{id}/files] GetSourceIDFiles", response, response.Code())
	}
}

// NewGetSourceIDFilesOK creates a GetSourceIDFilesOK with default headers values
func NewGetSourceIDFilesOK() *GetSourceIDFilesOK {
	return &GetSourceIDFilesOK{}
}

/*
GetSourceIDFilesOK describes a response with status code 200, with default header values.

OK
*/
type GetSourceIDFilesOK struct {
	Payload []*models.ModelFile
}

// IsSuccess returns true when this get source Id files o k response has a 2xx status code
func (o *GetSourceIDFilesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get source Id files o k response has a 3xx status code
func (o *GetSourceIDFilesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get source Id files o k response has a 4xx status code
func (o *GetSourceIDFilesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get source Id files o k response has a 5xx status code
func (o *GetSourceIDFilesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get source Id files o k response a status code equal to that given
func (o *GetSourceIDFilesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get source Id files o k response
func (o *GetSourceIDFilesOK) Code() int {
	return 200
}

func (o *GetSourceIDFilesOK) Error() string {
	return fmt.Sprintf("[GET /source/{id}/files][%d] getSourceIdFilesOK  %+v", 200, o.Payload)
}

func (o *GetSourceIDFilesOK) String() string {
	return fmt.Sprintf("[GET /source/{id}/files][%d] getSourceIdFilesOK  %+v", 200, o.Payload)
}

func (o *GetSourceIDFilesOK) GetPayload() []*models.ModelFile {
	return o.Payload
}

func (o *GetSourceIDFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSourceIDFilesBadRequest creates a GetSourceIDFilesBadRequest with default headers values
func NewGetSourceIDFilesBadRequest() *GetSourceIDFilesBadRequest {
	return &GetSourceIDFilesBadRequest{}
}

/*
GetSourceIDFilesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSourceIDFilesBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this get source Id files bad request response has a 2xx status code
func (o *GetSourceIDFilesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get source Id files bad request response has a 3xx status code
func (o *GetSourceIDFilesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get source Id files bad request response has a 4xx status code
func (o *GetSourceIDFilesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get source Id files bad request response has a 5xx status code
func (o *GetSourceIDFilesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get source Id files bad request response a status code equal to that given
func (o *GetSourceIDFilesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get source Id files bad request response
func (o *GetSourceIDFilesBadRequest) Code() int {
	return 400
}

func (o *GetSourceIDFilesBadRequest) Error() string {
	return fmt.Sprintf("[GET /source/{id}/files][%d] getSourceIdFilesBadRequest  %+v", 400, o.Payload)
}

func (o *GetSourceIDFilesBadRequest) String() string {
	return fmt.Sprintf("[GET /source/{id}/files][%d] getSourceIdFilesBadRequest  %+v", 400, o.Payload)
}

func (o *GetSourceIDFilesBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *GetSourceIDFilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSourceIDFilesInternalServerError creates a GetSourceIDFilesInternalServerError with default headers values
func NewGetSourceIDFilesInternalServerError() *GetSourceIDFilesInternalServerError {
	return &GetSourceIDFilesInternalServerError{}
}

/*
GetSourceIDFilesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetSourceIDFilesInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this get source Id files internal server error response has a 2xx status code
func (o *GetSourceIDFilesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get source Id files internal server error response has a 3xx status code
func (o *GetSourceIDFilesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get source Id files internal server error response has a 4xx status code
func (o *GetSourceIDFilesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get source Id files internal server error response has a 5xx status code
func (o *GetSourceIDFilesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get source Id files internal server error response a status code equal to that given
func (o *GetSourceIDFilesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get source Id files internal server error response
func (o *GetSourceIDFilesInternalServerError) Code() int {
	return 500
}

func (o *GetSourceIDFilesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /source/{id}/files][%d] getSourceIdFilesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSourceIDFilesInternalServerError) String() string {
	return fmt.Sprintf("[GET /source/{id}/files][%d] getSourceIdFilesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSourceIDFilesInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *GetSourceIDFilesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
