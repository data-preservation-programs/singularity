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

// PostSourceGcsDatasetDatasetNameReader is a Reader for the PostSourceGcsDatasetDatasetName structure.
type PostSourceGcsDatasetDatasetNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSourceGcsDatasetDatasetNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSourceGcsDatasetDatasetNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSourceGcsDatasetDatasetNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSourceGcsDatasetDatasetNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /source/gcs/dataset/{datasetName}] PostSourceGcsDatasetDatasetName", response, response.Code())
	}
}

// NewPostSourceGcsDatasetDatasetNameOK creates a PostSourceGcsDatasetDatasetNameOK with default headers values
func NewPostSourceGcsDatasetDatasetNameOK() *PostSourceGcsDatasetDatasetNameOK {
	return &PostSourceGcsDatasetDatasetNameOK{}
}

/*
PostSourceGcsDatasetDatasetNameOK describes a response with status code 200, with default header values.

OK
*/
type PostSourceGcsDatasetDatasetNameOK struct {
	Payload *models.ModelSource
}

// IsSuccess returns true when this post source gcs dataset dataset name o k response has a 2xx status code
func (o *PostSourceGcsDatasetDatasetNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post source gcs dataset dataset name o k response has a 3xx status code
func (o *PostSourceGcsDatasetDatasetNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source gcs dataset dataset name o k response has a 4xx status code
func (o *PostSourceGcsDatasetDatasetNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source gcs dataset dataset name o k response has a 5xx status code
func (o *PostSourceGcsDatasetDatasetNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post source gcs dataset dataset name o k response a status code equal to that given
func (o *PostSourceGcsDatasetDatasetNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post source gcs dataset dataset name o k response
func (o *PostSourceGcsDatasetDatasetNameOK) Code() int {
	return 200
}

func (o *PostSourceGcsDatasetDatasetNameOK) Error() string {
	return fmt.Sprintf("[POST /source/gcs/dataset/{datasetName}][%d] postSourceGcsDatasetDatasetNameOK  %+v", 200, o.Payload)
}

func (o *PostSourceGcsDatasetDatasetNameOK) String() string {
	return fmt.Sprintf("[POST /source/gcs/dataset/{datasetName}][%d] postSourceGcsDatasetDatasetNameOK  %+v", 200, o.Payload)
}

func (o *PostSourceGcsDatasetDatasetNameOK) GetPayload() *models.ModelSource {
	return o.Payload
}

func (o *PostSourceGcsDatasetDatasetNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceGcsDatasetDatasetNameBadRequest creates a PostSourceGcsDatasetDatasetNameBadRequest with default headers values
func NewPostSourceGcsDatasetDatasetNameBadRequest() *PostSourceGcsDatasetDatasetNameBadRequest {
	return &PostSourceGcsDatasetDatasetNameBadRequest{}
}

/*
PostSourceGcsDatasetDatasetNameBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostSourceGcsDatasetDatasetNameBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source gcs dataset dataset name bad request response has a 2xx status code
func (o *PostSourceGcsDatasetDatasetNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source gcs dataset dataset name bad request response has a 3xx status code
func (o *PostSourceGcsDatasetDatasetNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source gcs dataset dataset name bad request response has a 4xx status code
func (o *PostSourceGcsDatasetDatasetNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post source gcs dataset dataset name bad request response has a 5xx status code
func (o *PostSourceGcsDatasetDatasetNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post source gcs dataset dataset name bad request response a status code equal to that given
func (o *PostSourceGcsDatasetDatasetNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post source gcs dataset dataset name bad request response
func (o *PostSourceGcsDatasetDatasetNameBadRequest) Code() int {
	return 400
}

func (o *PostSourceGcsDatasetDatasetNameBadRequest) Error() string {
	return fmt.Sprintf("[POST /source/gcs/dataset/{datasetName}][%d] postSourceGcsDatasetDatasetNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceGcsDatasetDatasetNameBadRequest) String() string {
	return fmt.Sprintf("[POST /source/gcs/dataset/{datasetName}][%d] postSourceGcsDatasetDatasetNameBadRequest  %+v", 400, o.Payload)
}

func (o *PostSourceGcsDatasetDatasetNameBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourceGcsDatasetDatasetNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSourceGcsDatasetDatasetNameInternalServerError creates a PostSourceGcsDatasetDatasetNameInternalServerError with default headers values
func NewPostSourceGcsDatasetDatasetNameInternalServerError() *PostSourceGcsDatasetDatasetNameInternalServerError {
	return &PostSourceGcsDatasetDatasetNameInternalServerError{}
}

/*
PostSourceGcsDatasetDatasetNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostSourceGcsDatasetDatasetNameInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this post source gcs dataset dataset name internal server error response has a 2xx status code
func (o *PostSourceGcsDatasetDatasetNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post source gcs dataset dataset name internal server error response has a 3xx status code
func (o *PostSourceGcsDatasetDatasetNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post source gcs dataset dataset name internal server error response has a 4xx status code
func (o *PostSourceGcsDatasetDatasetNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post source gcs dataset dataset name internal server error response has a 5xx status code
func (o *PostSourceGcsDatasetDatasetNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post source gcs dataset dataset name internal server error response a status code equal to that given
func (o *PostSourceGcsDatasetDatasetNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post source gcs dataset dataset name internal server error response
func (o *PostSourceGcsDatasetDatasetNameInternalServerError) Code() int {
	return 500
}

func (o *PostSourceGcsDatasetDatasetNameInternalServerError) Error() string {
	return fmt.Sprintf("[POST /source/gcs/dataset/{datasetName}][%d] postSourceGcsDatasetDatasetNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceGcsDatasetDatasetNameInternalServerError) String() string {
	return fmt.Sprintf("[POST /source/gcs/dataset/{datasetName}][%d] postSourceGcsDatasetDatasetNameInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSourceGcsDatasetDatasetNameInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *PostSourceGcsDatasetDatasetNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
