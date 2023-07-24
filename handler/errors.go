package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type Error struct {
	Err            error
	HTTPStatusCode int
}

type HTTPError struct {
	Err string `json:"err"`
}

func (e *Error) HTTPResponse(c echo.Context) error {
	if e == nil {
		return c.String(http.StatusOK, "OK")
	}

	if e.HTTPStatusCode == 0 {
		e.HTTPStatusCode = http.StatusInternalServerError
	}

	return c.JSON(e.HTTPStatusCode, HTTPError{Err: e.Error()})
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func NewBadRequestError(err error) *Error {
	return &Error{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
	}
}

func NewBadRequestString(err string) *Error {
	return &Error{
		Err:            errors.New(err),
		HTTPStatusCode: http.StatusBadRequest,
	}
}

func NewNotFoundError(err error) *Error {
	return &Error{
		Err:            err,
		HTTPStatusCode: http.StatusNotFound,
	}
}

func NewNotFoundString(err string) *Error {
	return &Error{
		Err:            errors.New(err),
		HTTPStatusCode: http.StatusNotFound,
	}
}

func NewConflictError(err error) *Error {
	return &Error{
		Err:            err,
		HTTPStatusCode: http.StatusConflict,
	}
}

func NewConflictString(err string) *Error {
	return &Error{
		Err:            errors.New(err),
		HTTPStatusCode: http.StatusConflict,
	}
}

func NewHTTPError(code int, err string) *Error {
	return &Error{
		Err:            errors.New(err),
		HTTPStatusCode: code,
	}
}

func NewHandlerError(err error) *Error {
	return &Error{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
	}
}
