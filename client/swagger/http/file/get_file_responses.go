// Code generated by go-swagger; DO NOT EDIT.

package file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// GetFileReader is a Reader for the GetFile structure.
type GetFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /file/{id}] GetFile", response, response.Code())
	}
}

// NewGetFileOK creates a GetFileOK with default headers values
func NewGetFileOK() *GetFileOK {
	return &GetFileOK{}
}

/*
GetFileOK describes a response with status code 200, with default header values.

OK
*/
type GetFileOK struct {
	Payload *models.ModelFile
}

// IsSuccess returns true when this get file o k response has a 2xx status code
func (o *GetFileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get file o k response has a 3xx status code
func (o *GetFileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file o k response has a 4xx status code
func (o *GetFileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get file o k response has a 5xx status code
func (o *GetFileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get file o k response a status code equal to that given
func (o *GetFileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get file o k response
func (o *GetFileOK) Code() int {
	return 200
}

func (o *GetFileOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file/{id}][%d] getFileOK %s", 200, payload)
}

func (o *GetFileOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file/{id}][%d] getFileOK %s", 200, payload)
}

func (o *GetFileOK) GetPayload() *models.ModelFile {
	return o.Payload
}

func (o *GetFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelFile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFileInternalServerError creates a GetFileInternalServerError with default headers values
func NewGetFileInternalServerError() *GetFileInternalServerError {
	return &GetFileInternalServerError{}
}

/*
GetFileInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetFileInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this get file internal server error response has a 2xx status code
func (o *GetFileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get file internal server error response has a 3xx status code
func (o *GetFileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file internal server error response has a 4xx status code
func (o *GetFileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get file internal server error response has a 5xx status code
func (o *GetFileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get file internal server error response a status code equal to that given
func (o *GetFileInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get file internal server error response
func (o *GetFileInternalServerError) Code() int {
	return 500
}

func (o *GetFileInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file/{id}][%d] getFileInternalServerError %s", 500, payload)
}

func (o *GetFileInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file/{id}][%d] getFileInternalServerError %s", 500, payload)
}

func (o *GetFileInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *GetFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
