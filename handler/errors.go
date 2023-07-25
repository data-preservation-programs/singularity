package handler

import (
	"fmt"

	"github.com/pkg/errors"
)

type ErrInvalidParameter struct {
	Err error
}

func (e ErrInvalidParameter) Unwrap() error {
	return e.Err
}
func (e ErrInvalidParameter) Error() string {
	return fmt.Sprintf("invalid parameter: %s", e.Err.Error())
}

func NewInvalidParameterErr(err string) ErrInvalidParameter {
	return ErrInvalidParameter{
		Err: errors.New(err),
	}
}

type ErrNotFound struct {
	Err error
}

func (e ErrNotFound) Unwrap() error {
	return e.Err
}
func (e ErrNotFound) Error() string {
	return fmt.Sprintf("not found: %s", e.Err.Error())
}

type ErrDuplicateRecord struct {
	Err error
}

func (e ErrDuplicateRecord) Unwrap() error {
	return e.Err
}

func (e ErrDuplicateRecord) Error() string {
	return fmt.Sprintf("duplicate record: %s", e.Err.Error())
}

func NewDuplicateRecordError(err string) ErrDuplicateRecord {
	return ErrDuplicateRecord{
		Err: errors.New(err),
	}
}
