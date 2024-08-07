// Code generated by go-swagger; DO NOT EDIT.

package wallet_association

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

// AttachWalletReader is a Reader for the AttachWallet structure.
type AttachWalletReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttachWalletReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAttachWalletOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAttachWalletBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAttachWalletInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /preparation/{id}/wallet/{wallet}] AttachWallet", response, response.Code())
	}
}

// NewAttachWalletOK creates a AttachWalletOK with default headers values
func NewAttachWalletOK() *AttachWalletOK {
	return &AttachWalletOK{}
}

/*
AttachWalletOK describes a response with status code 200, with default header values.

OK
*/
type AttachWalletOK struct {
	Payload *models.ModelPreparation
}

// IsSuccess returns true when this attach wallet o k response has a 2xx status code
func (o *AttachWalletOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this attach wallet o k response has a 3xx status code
func (o *AttachWalletOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach wallet o k response has a 4xx status code
func (o *AttachWalletOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this attach wallet o k response has a 5xx status code
func (o *AttachWalletOK) IsServerError() bool {
	return false
}

// IsCode returns true when this attach wallet o k response a status code equal to that given
func (o *AttachWalletOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the attach wallet o k response
func (o *AttachWalletOK) Code() int {
	return 200
}

func (o *AttachWalletOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /preparation/{id}/wallet/{wallet}][%d] attachWalletOK %s", 200, payload)
}

func (o *AttachWalletOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /preparation/{id}/wallet/{wallet}][%d] attachWalletOK %s", 200, payload)
}

func (o *AttachWalletOK) GetPayload() *models.ModelPreparation {
	return o.Payload
}

func (o *AttachWalletOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelPreparation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachWalletBadRequest creates a AttachWalletBadRequest with default headers values
func NewAttachWalletBadRequest() *AttachWalletBadRequest {
	return &AttachWalletBadRequest{}
}

/*
AttachWalletBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AttachWalletBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this attach wallet bad request response has a 2xx status code
func (o *AttachWalletBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach wallet bad request response has a 3xx status code
func (o *AttachWalletBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach wallet bad request response has a 4xx status code
func (o *AttachWalletBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach wallet bad request response has a 5xx status code
func (o *AttachWalletBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this attach wallet bad request response a status code equal to that given
func (o *AttachWalletBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the attach wallet bad request response
func (o *AttachWalletBadRequest) Code() int {
	return 400
}

func (o *AttachWalletBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /preparation/{id}/wallet/{wallet}][%d] attachWalletBadRequest %s", 400, payload)
}

func (o *AttachWalletBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /preparation/{id}/wallet/{wallet}][%d] attachWalletBadRequest %s", 400, payload)
}

func (o *AttachWalletBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *AttachWalletBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachWalletInternalServerError creates a AttachWalletInternalServerError with default headers values
func NewAttachWalletInternalServerError() *AttachWalletInternalServerError {
	return &AttachWalletInternalServerError{}
}

/*
AttachWalletInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type AttachWalletInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this attach wallet internal server error response has a 2xx status code
func (o *AttachWalletInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach wallet internal server error response has a 3xx status code
func (o *AttachWalletInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach wallet internal server error response has a 4xx status code
func (o *AttachWalletInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this attach wallet internal server error response has a 5xx status code
func (o *AttachWalletInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this attach wallet internal server error response a status code equal to that given
func (o *AttachWalletInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the attach wallet internal server error response
func (o *AttachWalletInternalServerError) Code() int {
	return 500
}

func (o *AttachWalletInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /preparation/{id}/wallet/{wallet}][%d] attachWalletInternalServerError %s", 500, payload)
}

func (o *AttachWalletInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /preparation/{id}/wallet/{wallet}][%d] attachWalletInternalServerError %s", 500, payload)
}

func (o *AttachWalletInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *AttachWalletInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
