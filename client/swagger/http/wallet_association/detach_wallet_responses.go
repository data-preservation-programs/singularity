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

// DetachWalletReader is a Reader for the DetachWallet structure.
type DetachWalletReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetachWalletReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDetachWalletOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDetachWalletBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDetachWalletInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /preparation/{id}/wallet/{wallet}] DetachWallet", response, response.Code())
	}
}

// NewDetachWalletOK creates a DetachWalletOK with default headers values
func NewDetachWalletOK() *DetachWalletOK {
	return &DetachWalletOK{}
}

/*
DetachWalletOK describes a response with status code 200, with default header values.

OK
*/
type DetachWalletOK struct {
	Payload *models.ModelPreparation
}

// IsSuccess returns true when this detach wallet o k response has a 2xx status code
func (o *DetachWalletOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this detach wallet o k response has a 3xx status code
func (o *DetachWalletOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this detach wallet o k response has a 4xx status code
func (o *DetachWalletOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this detach wallet o k response has a 5xx status code
func (o *DetachWalletOK) IsServerError() bool {
	return false
}

// IsCode returns true when this detach wallet o k response a status code equal to that given
func (o *DetachWalletOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the detach wallet o k response
func (o *DetachWalletOK) Code() int {
	return 200
}

func (o *DetachWalletOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /preparation/{id}/wallet/{wallet}][%d] detachWalletOK %s", 200, payload)
}

func (o *DetachWalletOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /preparation/{id}/wallet/{wallet}][%d] detachWalletOK %s", 200, payload)
}

func (o *DetachWalletOK) GetPayload() *models.ModelPreparation {
	return o.Payload
}

func (o *DetachWalletOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelPreparation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetachWalletBadRequest creates a DetachWalletBadRequest with default headers values
func NewDetachWalletBadRequest() *DetachWalletBadRequest {
	return &DetachWalletBadRequest{}
}

/*
DetachWalletBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DetachWalletBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this detach wallet bad request response has a 2xx status code
func (o *DetachWalletBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this detach wallet bad request response has a 3xx status code
func (o *DetachWalletBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this detach wallet bad request response has a 4xx status code
func (o *DetachWalletBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this detach wallet bad request response has a 5xx status code
func (o *DetachWalletBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this detach wallet bad request response a status code equal to that given
func (o *DetachWalletBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the detach wallet bad request response
func (o *DetachWalletBadRequest) Code() int {
	return 400
}

func (o *DetachWalletBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /preparation/{id}/wallet/{wallet}][%d] detachWalletBadRequest %s", 400, payload)
}

func (o *DetachWalletBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /preparation/{id}/wallet/{wallet}][%d] detachWalletBadRequest %s", 400, payload)
}

func (o *DetachWalletBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *DetachWalletBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetachWalletInternalServerError creates a DetachWalletInternalServerError with default headers values
func NewDetachWalletInternalServerError() *DetachWalletInternalServerError {
	return &DetachWalletInternalServerError{}
}

/*
DetachWalletInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DetachWalletInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this detach wallet internal server error response has a 2xx status code
func (o *DetachWalletInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this detach wallet internal server error response has a 3xx status code
func (o *DetachWalletInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this detach wallet internal server error response has a 4xx status code
func (o *DetachWalletInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this detach wallet internal server error response has a 5xx status code
func (o *DetachWalletInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this detach wallet internal server error response a status code equal to that given
func (o *DetachWalletInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the detach wallet internal server error response
func (o *DetachWalletInternalServerError) Code() int {
	return 500
}

func (o *DetachWalletInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /preparation/{id}/wallet/{wallet}][%d] detachWalletInternalServerError %s", 500, payload)
}

func (o *DetachWalletInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /preparation/{id}/wallet/{wallet}][%d] detachWalletInternalServerError %s", 500, payload)
}

func (o *DetachWalletInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *DetachWalletInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
