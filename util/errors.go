package util

import (
	"fmt"
	"strings"

	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
)

type AggregateError struct {
	Errors []error
}

func (a AggregateError) Error() string {
	errors := make([]string, len(a.Errors))
	for i, err := range a.Errors {
		errors[i] = err.Error()
	}

	return strings.Join(errors, ", ")
}

func (a AggregateError) Unwrap() error {
	if len(a.Errors) == 0 {
		return nil
	}

	return a.Errors[len(a.Errors)-1]
}

func (a AggregateError) Is(err error) bool {
	for _, e := range a.Errors {
		if errors.Is(e, err) {
			return true
		}
	}

	return false
}

func (a AggregateError) As(err interface{}) bool {
	for _, e := range a.Errors {
		if errors.As(e, err) {
			return true
		}
	}

	return false
}

func (a AggregateError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = s.Write([]byte(a.Error()))
			for _, e := range a.Errors {
				_, _ = s.Write([]byte("\n"))
				errors.FormatError(e, s, verb)
			}
			return
		}
		fallthrough
	case 's':
		_, _ = s.Write([]byte(a.Error()))
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", a.Error())
	}
}

func IsDuplicateKeyError(err error) bool {
	return errors.Is(err, gorm.ErrDuplicatedKey) || (err != nil && strings.Contains(err.Error(), "constraint failed"))
}

func IsForeignKeyConstraintError(err error) bool {
	return errors.Is(err, gorm.ErrForeignKeyViolated) || (err != nil && strings.Contains(err.Error(), "foreign key constraint"))
}
