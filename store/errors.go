package store

import (
	"strings"
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
