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
)

// PrepareToPackFileReader is a Reader for the PrepareToPackFile structure.
type PrepareToPackFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PrepareToPackFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPrepareToPackFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPrepareToPackFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPrepareToPackFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /file/{id}/prepare_to_pack] PrepareToPackFile", response, response.Code())
	}
}

// NewPrepareToPackFileOK creates a PrepareToPackFileOK with default headers values
func NewPrepareToPackFileOK() *PrepareToPackFileOK {
	return &PrepareToPackFileOK{}
}

/*
PrepareToPackFileOK describes a response with status code 200, with default header values.

OK
*/
type PrepareToPackFileOK struct {
	Payload int64
}

// IsSuccess returns true when this prepare to pack file o k response has a 2xx status code
func (o *PrepareToPackFileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this prepare to pack file o k response has a 3xx status code
func (o *PrepareToPackFileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prepare to pack file o k response has a 4xx status code
func (o *PrepareToPackFileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this prepare to pack file o k response has a 5xx status code
func (o *PrepareToPackFileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this prepare to pack file o k response a status code equal to that given
func (o *PrepareToPackFileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the prepare to pack file o k response
func (o *PrepareToPackFileOK) Code() int {
	return 200
}

func (o *PrepareToPackFileOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /file/{id}/prepare_to_pack][%d] prepareToPackFileOK %s", 200, payload)
}

func (o *PrepareToPackFileOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /file/{id}/prepare_to_pack][%d] prepareToPackFileOK %s", 200, payload)
}

func (o *PrepareToPackFileOK) GetPayload() int64 {
	return o.Payload
}

func (o *PrepareToPackFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrepareToPackFileBadRequest creates a PrepareToPackFileBadRequest with default headers values
func NewPrepareToPackFileBadRequest() *PrepareToPackFileBadRequest {
	return &PrepareToPackFileBadRequest{}
}

/*
PrepareToPackFileBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PrepareToPackFileBadRequest struct {
	Payload string
}

// IsSuccess returns true when this prepare to pack file bad request response has a 2xx status code
func (o *PrepareToPackFileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prepare to pack file bad request response has a 3xx status code
func (o *PrepareToPackFileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prepare to pack file bad request response has a 4xx status code
func (o *PrepareToPackFileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this prepare to pack file bad request response has a 5xx status code
func (o *PrepareToPackFileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this prepare to pack file bad request response a status code equal to that given
func (o *PrepareToPackFileBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the prepare to pack file bad request response
func (o *PrepareToPackFileBadRequest) Code() int {
	return 400
}

func (o *PrepareToPackFileBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /file/{id}/prepare_to_pack][%d] prepareToPackFileBadRequest %s", 400, payload)
}

func (o *PrepareToPackFileBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /file/{id}/prepare_to_pack][%d] prepareToPackFileBadRequest %s", 400, payload)
}

func (o *PrepareToPackFileBadRequest) GetPayload() string {
	return o.Payload
}

func (o *PrepareToPackFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrepareToPackFileInternalServerError creates a PrepareToPackFileInternalServerError with default headers values
func NewPrepareToPackFileInternalServerError() *PrepareToPackFileInternalServerError {
	return &PrepareToPackFileInternalServerError{}
}

/*
PrepareToPackFileInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PrepareToPackFileInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this prepare to pack file internal server error response has a 2xx status code
func (o *PrepareToPackFileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prepare to pack file internal server error response has a 3xx status code
func (o *PrepareToPackFileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prepare to pack file internal server error response has a 4xx status code
func (o *PrepareToPackFileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this prepare to pack file internal server error response has a 5xx status code
func (o *PrepareToPackFileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this prepare to pack file internal server error response a status code equal to that given
func (o *PrepareToPackFileInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the prepare to pack file internal server error response
func (o *PrepareToPackFileInternalServerError) Code() int {
	return 500
}

func (o *PrepareToPackFileInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /file/{id}/prepare_to_pack][%d] prepareToPackFileInternalServerError %s", 500, payload)
}

func (o *PrepareToPackFileInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /file/{id}/prepare_to_pack][%d] prepareToPackFileInternalServerError %s", 500, payload)
}

func (o *PrepareToPackFileInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *PrepareToPackFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
