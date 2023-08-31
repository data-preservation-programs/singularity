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

// GetStorageNameExplorePathReader is a Reader for the GetStorageNameExplorePath structure.
type GetStorageNameExplorePathReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStorageNameExplorePathReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStorageNameExplorePathOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetStorageNameExplorePathBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetStorageNameExplorePathInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /storage/{name}/explore/{path}] GetStorageNameExplorePath", response, response.Code())
	}
}

// NewGetStorageNameExplorePathOK creates a GetStorageNameExplorePathOK with default headers values
func NewGetStorageNameExplorePathOK() *GetStorageNameExplorePathOK {
	return &GetStorageNameExplorePathOK{}
}

/*
GetStorageNameExplorePathOK describes a response with status code 200, with default header values.

OK
*/
type GetStorageNameExplorePathOK struct {
	Payload []*models.StorageDirEntry
}

// IsSuccess returns true when this get storage name explore path o k response has a 2xx status code
func (o *GetStorageNameExplorePathOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get storage name explore path o k response has a 3xx status code
func (o *GetStorageNameExplorePathOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get storage name explore path o k response has a 4xx status code
func (o *GetStorageNameExplorePathOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get storage name explore path o k response has a 5xx status code
func (o *GetStorageNameExplorePathOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get storage name explore path o k response a status code equal to that given
func (o *GetStorageNameExplorePathOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get storage name explore path o k response
func (o *GetStorageNameExplorePathOK) Code() int {
	return 200
}

func (o *GetStorageNameExplorePathOK) Error() string {
	return fmt.Sprintf("[GET /storage/{name}/explore/{path}][%d] getStorageNameExplorePathOK  %+v", 200, o.Payload)
}

func (o *GetStorageNameExplorePathOK) String() string {
	return fmt.Sprintf("[GET /storage/{name}/explore/{path}][%d] getStorageNameExplorePathOK  %+v", 200, o.Payload)
}

func (o *GetStorageNameExplorePathOK) GetPayload() []*models.StorageDirEntry {
	return o.Payload
}

func (o *GetStorageNameExplorePathOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStorageNameExplorePathBadRequest creates a GetStorageNameExplorePathBadRequest with default headers values
func NewGetStorageNameExplorePathBadRequest() *GetStorageNameExplorePathBadRequest {
	return &GetStorageNameExplorePathBadRequest{}
}

/*
GetStorageNameExplorePathBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetStorageNameExplorePathBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this get storage name explore path bad request response has a 2xx status code
func (o *GetStorageNameExplorePathBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get storage name explore path bad request response has a 3xx status code
func (o *GetStorageNameExplorePathBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get storage name explore path bad request response has a 4xx status code
func (o *GetStorageNameExplorePathBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get storage name explore path bad request response has a 5xx status code
func (o *GetStorageNameExplorePathBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get storage name explore path bad request response a status code equal to that given
func (o *GetStorageNameExplorePathBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get storage name explore path bad request response
func (o *GetStorageNameExplorePathBadRequest) Code() int {
	return 400
}

func (o *GetStorageNameExplorePathBadRequest) Error() string {
	return fmt.Sprintf("[GET /storage/{name}/explore/{path}][%d] getStorageNameExplorePathBadRequest  %+v", 400, o.Payload)
}

func (o *GetStorageNameExplorePathBadRequest) String() string {
	return fmt.Sprintf("[GET /storage/{name}/explore/{path}][%d] getStorageNameExplorePathBadRequest  %+v", 400, o.Payload)
}

func (o *GetStorageNameExplorePathBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *GetStorageNameExplorePathBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStorageNameExplorePathInternalServerError creates a GetStorageNameExplorePathInternalServerError with default headers values
func NewGetStorageNameExplorePathInternalServerError() *GetStorageNameExplorePathInternalServerError {
	return &GetStorageNameExplorePathInternalServerError{}
}

/*
GetStorageNameExplorePathInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetStorageNameExplorePathInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this get storage name explore path internal server error response has a 2xx status code
func (o *GetStorageNameExplorePathInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get storage name explore path internal server error response has a 3xx status code
func (o *GetStorageNameExplorePathInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get storage name explore path internal server error response has a 4xx status code
func (o *GetStorageNameExplorePathInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get storage name explore path internal server error response has a 5xx status code
func (o *GetStorageNameExplorePathInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get storage name explore path internal server error response a status code equal to that given
func (o *GetStorageNameExplorePathInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get storage name explore path internal server error response
func (o *GetStorageNameExplorePathInternalServerError) Code() int {
	return 500
}

func (o *GetStorageNameExplorePathInternalServerError) Error() string {
	return fmt.Sprintf("[GET /storage/{name}/explore/{path}][%d] getStorageNameExplorePathInternalServerError  %+v", 500, o.Payload)
}

func (o *GetStorageNameExplorePathInternalServerError) String() string {
	return fmt.Sprintf("[GET /storage/{name}/explore/{path}][%d] getStorageNameExplorePathInternalServerError  %+v", 500, o.Payload)
}

func (o *GetStorageNameExplorePathInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *GetStorageNameExplorePathInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
