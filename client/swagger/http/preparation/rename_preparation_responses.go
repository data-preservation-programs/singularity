// Code generated by go-swagger; DO NOT EDIT.

package preparation

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

// RenamePreparationReader is a Reader for the RenamePreparation structure.
type RenamePreparationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RenamePreparationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRenamePreparationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRenamePreparationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRenamePreparationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /preparation/{name}/rename] RenamePreparation", response, response.Code())
	}
}

// NewRenamePreparationOK creates a RenamePreparationOK with default headers values
func NewRenamePreparationOK() *RenamePreparationOK {
	return &RenamePreparationOK{}
}

/*
RenamePreparationOK describes a response with status code 200, with default header values.

OK
*/
type RenamePreparationOK struct {
	Payload *models.ModelPreparation
}

// IsSuccess returns true when this rename preparation o k response has a 2xx status code
func (o *RenamePreparationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rename preparation o k response has a 3xx status code
func (o *RenamePreparationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rename preparation o k response has a 4xx status code
func (o *RenamePreparationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this rename preparation o k response has a 5xx status code
func (o *RenamePreparationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this rename preparation o k response a status code equal to that given
func (o *RenamePreparationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the rename preparation o k response
func (o *RenamePreparationOK) Code() int {
	return 200
}

func (o *RenamePreparationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /preparation/{name}/rename][%d] renamePreparationOK %s", 200, payload)
}

func (o *RenamePreparationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /preparation/{name}/rename][%d] renamePreparationOK %s", 200, payload)
}

func (o *RenamePreparationOK) GetPayload() *models.ModelPreparation {
	return o.Payload
}

func (o *RenamePreparationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelPreparation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenamePreparationBadRequest creates a RenamePreparationBadRequest with default headers values
func NewRenamePreparationBadRequest() *RenamePreparationBadRequest {
	return &RenamePreparationBadRequest{}
}

/*
RenamePreparationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RenamePreparationBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this rename preparation bad request response has a 2xx status code
func (o *RenamePreparationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rename preparation bad request response has a 3xx status code
func (o *RenamePreparationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rename preparation bad request response has a 4xx status code
func (o *RenamePreparationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this rename preparation bad request response has a 5xx status code
func (o *RenamePreparationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this rename preparation bad request response a status code equal to that given
func (o *RenamePreparationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the rename preparation bad request response
func (o *RenamePreparationBadRequest) Code() int {
	return 400
}

func (o *RenamePreparationBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /preparation/{name}/rename][%d] renamePreparationBadRequest %s", 400, payload)
}

func (o *RenamePreparationBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /preparation/{name}/rename][%d] renamePreparationBadRequest %s", 400, payload)
}

func (o *RenamePreparationBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *RenamePreparationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenamePreparationInternalServerError creates a RenamePreparationInternalServerError with default headers values
func NewRenamePreparationInternalServerError() *RenamePreparationInternalServerError {
	return &RenamePreparationInternalServerError{}
}

/*
RenamePreparationInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RenamePreparationInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this rename preparation internal server error response has a 2xx status code
func (o *RenamePreparationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rename preparation internal server error response has a 3xx status code
func (o *RenamePreparationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rename preparation internal server error response has a 4xx status code
func (o *RenamePreparationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this rename preparation internal server error response has a 5xx status code
func (o *RenamePreparationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this rename preparation internal server error response a status code equal to that given
func (o *RenamePreparationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the rename preparation internal server error response
func (o *RenamePreparationInternalServerError) Code() int {
	return 500
}

func (o *RenamePreparationInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /preparation/{name}/rename][%d] renamePreparationInternalServerError %s", 500, payload)
}

func (o *RenamePreparationInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /preparation/{name}/rename][%d] renamePreparationInternalServerError %s", 500, payload)
}

func (o *RenamePreparationInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *RenamePreparationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
