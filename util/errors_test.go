package util

import (
	"testing"

	"github.com/pkg/errors"
)

func TestAggregateError(t *testing.T) {
	tests := []struct {
		name     string
		errors   []error
		expected string
	}{
		{
			name:     "multiple errors",
			errors:   []error{errors.New("error1"), errors.New("error2"), errors.New("error3")},
			expected: "error1, error2, error3",
		},
		{
			name:     "single error",
			errors:   []error{errors.New("only error")},
			expected: "only error",
		},
		{
			name:     "no errors",
			errors:   []error{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aggErr := AggregateError{Errors: tt.errors}
			if aggErr.Error() != tt.expected {
				t.Errorf("got %q, want %q", aggErr.Error(), tt.expected)
			}
		})
	}
}
