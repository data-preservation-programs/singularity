// Code generated by go-swagger; DO NOT EDIT.

package wallet_association

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// PostPreparationIDWalletWalletReader is a Reader for the PostPreparationIDWalletWallet structure.
type PostPreparationIDWalletWalletReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPreparationIDWalletWalletReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostPreparationIDWalletWalletOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostPreparationIDWalletWalletBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostPreparationIDWalletWalletInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /preparation/{id}/wallet/{wallet}] PostPreparationIDWalletWallet", response, response.Code())
	}
}

// NewPostPreparationIDWalletWalletOK creates a PostPreparationIDWalletWalletOK with default headers values
func NewPostPreparationIDWalletWalletOK() *PostPreparationIDWalletWalletOK {
	return &PostPreparationIDWalletWalletOK{}
}

/*
PostPreparationIDWalletWalletOK describes a response with status code 200, with default header values.

OK
*/
type PostPreparationIDWalletWalletOK struct {
	Payload *models.ModelPreparation
}

// IsSuccess returns true when this post preparation Id wallet wallet o k response has a 2xx status code
func (o *PostPreparationIDWalletWalletOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post preparation Id wallet wallet o k response has a 3xx status code
func (o *PostPreparationIDWalletWalletOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post preparation Id wallet wallet o k response has a 4xx status code
func (o *PostPreparationIDWalletWalletOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post preparation Id wallet wallet o k response has a 5xx status code
func (o *PostPreparationIDWalletWalletOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post preparation Id wallet wallet o k response a status code equal to that given
func (o *PostPreparationIDWalletWalletOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post preparation Id wallet wallet o k response
func (o *PostPreparationIDWalletWalletOK) Code() int {
	return 200
}

func (o *PostPreparationIDWalletWalletOK) Error() string {
	return fmt.Sprintf("[POST /preparation/{id}/wallet/{wallet}][%d] postPreparationIdWalletWalletOK  %+v", 200, o.Payload)
}

func (o *PostPreparationIDWalletWalletOK) String() string {
	return fmt.Sprintf("[POST /preparation/{id}/wallet/{wallet}][%d] postPreparationIdWalletWalletOK  %+v", 200, o.Payload)
}

func (o *PostPreparationIDWalletWalletOK) GetPayload() *models.ModelPreparation {
	return o.Payload
}

func (o *PostPreparationIDWalletWalletOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelPreparation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPreparationIDWalletWalletBadRequest creates a PostPreparationIDWalletWalletBadRequest with default headers values
func NewPostPreparationIDWalletWalletBadRequest() *PostPreparationIDWalletWalletBadRequest {
	return &PostPreparationIDWalletWalletBadRequest{}
}

/*
PostPreparationIDWalletWalletBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostPreparationIDWalletWalletBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post preparation Id wallet wallet bad request response has a 2xx status code
func (o *PostPreparationIDWalletWalletBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post preparation Id wallet wallet bad request response has a 3xx status code
func (o *PostPreparationIDWalletWalletBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post preparation Id wallet wallet bad request response has a 4xx status code
func (o *PostPreparationIDWalletWalletBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post preparation Id wallet wallet bad request response has a 5xx status code
func (o *PostPreparationIDWalletWalletBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post preparation Id wallet wallet bad request response a status code equal to that given
func (o *PostPreparationIDWalletWalletBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post preparation Id wallet wallet bad request response
func (o *PostPreparationIDWalletWalletBadRequest) Code() int {
	return 400
}

func (o *PostPreparationIDWalletWalletBadRequest) Error() string {
	return fmt.Sprintf("[POST /preparation/{id}/wallet/{wallet}][%d] postPreparationIdWalletWalletBadRequest  %+v", 400, o.Payload)
}

func (o *PostPreparationIDWalletWalletBadRequest) String() string {
	return fmt.Sprintf("[POST /preparation/{id}/wallet/{wallet}][%d] postPreparationIdWalletWalletBadRequest  %+v", 400, o.Payload)
}

func (o *PostPreparationIDWalletWalletBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostPreparationIDWalletWalletBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPreparationIDWalletWalletInternalServerError creates a PostPreparationIDWalletWalletInternalServerError with default headers values
func NewPostPreparationIDWalletWalletInternalServerError() *PostPreparationIDWalletWalletInternalServerError {
	return &PostPreparationIDWalletWalletInternalServerError{}
}

/*
PostPreparationIDWalletWalletInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostPreparationIDWalletWalletInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post preparation Id wallet wallet internal server error response has a 2xx status code
func (o *PostPreparationIDWalletWalletInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post preparation Id wallet wallet internal server error response has a 3xx status code
func (o *PostPreparationIDWalletWalletInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post preparation Id wallet wallet internal server error response has a 4xx status code
func (o *PostPreparationIDWalletWalletInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post preparation Id wallet wallet internal server error response has a 5xx status code
func (o *PostPreparationIDWalletWalletInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post preparation Id wallet wallet internal server error response a status code equal to that given
func (o *PostPreparationIDWalletWalletInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post preparation Id wallet wallet internal server error response
func (o *PostPreparationIDWalletWalletInternalServerError) Code() int {
	return 500
}

func (o *PostPreparationIDWalletWalletInternalServerError) Error() string {
	return fmt.Sprintf("[POST /preparation/{id}/wallet/{wallet}][%d] postPreparationIdWalletWalletInternalServerError  %+v", 500, o.Payload)
}

func (o *PostPreparationIDWalletWalletInternalServerError) String() string {
	return fmt.Sprintf("[POST /preparation/{id}/wallet/{wallet}][%d] postPreparationIdWalletWalletInternalServerError  %+v", 500, o.Payload)
}

func (o *PostPreparationIDWalletWalletInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostPreparationIDWalletWalletInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
