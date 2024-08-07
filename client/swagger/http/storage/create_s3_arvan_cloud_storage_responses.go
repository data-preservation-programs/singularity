// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// CreateS3ArvanCloudStorageReader is a Reader for the CreateS3ArvanCloudStorage structure.
type CreateS3ArvanCloudStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateS3ArvanCloudStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateS3ArvanCloudStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateS3ArvanCloudStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateS3ArvanCloudStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /storage/s3/arvancloud] CreateS3ArvanCloudStorage", response, response.Code())
	}
}

// NewCreateS3ArvanCloudStorageOK creates a CreateS3ArvanCloudStorageOK with default headers values
func NewCreateS3ArvanCloudStorageOK() *CreateS3ArvanCloudStorageOK {
	return &CreateS3ArvanCloudStorageOK{}
}

/*
CreateS3ArvanCloudStorageOK describes a response with status code 200, with default header values.

OK
*/
type CreateS3ArvanCloudStorageOK struct {
	Payload *models.ModelStorage
}

// IsSuccess returns true when this create s3 arvan cloud storage o k response has a 2xx status code
func (o *CreateS3ArvanCloudStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create s3 arvan cloud storage o k response has a 3xx status code
func (o *CreateS3ArvanCloudStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 arvan cloud storage o k response has a 4xx status code
func (o *CreateS3ArvanCloudStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create s3 arvan cloud storage o k response has a 5xx status code
func (o *CreateS3ArvanCloudStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create s3 arvan cloud storage o k response a status code equal to that given
func (o *CreateS3ArvanCloudStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create s3 arvan cloud storage o k response
func (o *CreateS3ArvanCloudStorageOK) Code() int {
	return 200
}

func (o *CreateS3ArvanCloudStorageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/s3/arvancloud][%d] createS3ArvanCloudStorageOK %s", 200, payload)
}

func (o *CreateS3ArvanCloudStorageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/s3/arvancloud][%d] createS3ArvanCloudStorageOK %s", 200, payload)
}

func (o *CreateS3ArvanCloudStorageOK) GetPayload() *models.ModelStorage {
	return o.Payload
}

func (o *CreateS3ArvanCloudStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelStorage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateS3ArvanCloudStorageBadRequest creates a CreateS3ArvanCloudStorageBadRequest with default headers values
func NewCreateS3ArvanCloudStorageBadRequest() *CreateS3ArvanCloudStorageBadRequest {
	return &CreateS3ArvanCloudStorageBadRequest{}
}

/*
CreateS3ArvanCloudStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateS3ArvanCloudStorageBadRequest struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create s3 arvan cloud storage bad request response has a 2xx status code
func (o *CreateS3ArvanCloudStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create s3 arvan cloud storage bad request response has a 3xx status code
func (o *CreateS3ArvanCloudStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 arvan cloud storage bad request response has a 4xx status code
func (o *CreateS3ArvanCloudStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create s3 arvan cloud storage bad request response has a 5xx status code
func (o *CreateS3ArvanCloudStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create s3 arvan cloud storage bad request response a status code equal to that given
func (o *CreateS3ArvanCloudStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create s3 arvan cloud storage bad request response
func (o *CreateS3ArvanCloudStorageBadRequest) Code() int {
	return 400
}

func (o *CreateS3ArvanCloudStorageBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/s3/arvancloud][%d] createS3ArvanCloudStorageBadRequest %s", 400, payload)
}

func (o *CreateS3ArvanCloudStorageBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/s3/arvancloud][%d] createS3ArvanCloudStorageBadRequest %s", 400, payload)
}

func (o *CreateS3ArvanCloudStorageBadRequest) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateS3ArvanCloudStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateS3ArvanCloudStorageInternalServerError creates a CreateS3ArvanCloudStorageInternalServerError with default headers values
func NewCreateS3ArvanCloudStorageInternalServerError() *CreateS3ArvanCloudStorageInternalServerError {
	return &CreateS3ArvanCloudStorageInternalServerError{}
}

/*
CreateS3ArvanCloudStorageInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateS3ArvanCloudStorageInternalServerError struct {
	Payload *models.APIHTTPError
}

// IsSuccess returns true when this create s3 arvan cloud storage internal server error response has a 2xx status code
func (o *CreateS3ArvanCloudStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create s3 arvan cloud storage internal server error response has a 3xx status code
func (o *CreateS3ArvanCloudStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create s3 arvan cloud storage internal server error response has a 4xx status code
func (o *CreateS3ArvanCloudStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create s3 arvan cloud storage internal server error response has a 5xx status code
func (o *CreateS3ArvanCloudStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create s3 arvan cloud storage internal server error response a status code equal to that given
func (o *CreateS3ArvanCloudStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create s3 arvan cloud storage internal server error response
func (o *CreateS3ArvanCloudStorageInternalServerError) Code() int {
	return 500
}

func (o *CreateS3ArvanCloudStorageInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/s3/arvancloud][%d] createS3ArvanCloudStorageInternalServerError %s", 500, payload)
}

func (o *CreateS3ArvanCloudStorageInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/s3/arvancloud][%d] createS3ArvanCloudStorageInternalServerError %s", 500, payload)
}

func (o *CreateS3ArvanCloudStorageInternalServerError) GetPayload() *models.APIHTTPError {
	return o.Payload
}

func (o *CreateS3ArvanCloudStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
