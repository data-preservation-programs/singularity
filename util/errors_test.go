package util

import (
	"fmt"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/stretchr/testify/require"
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

type TestError struct {
	msg string
}

func (e TestError) Error() string {
	return e.msg
}

func TestAggregateErrorOthers(t *testing.T) {
	err1 := errors.New("error1")
	err2 := errors.New("error2")
	err3 := errors.New("error3")
	testError := TestError{msg: "test error"}
	agg := AggregateError{Errors: []error{err1, err2, err3}}
	err := agg.Unwrap()
	require.ErrorContains(t, err, "error3")
	require.True(t, errors.Is(err, err3))
	require.True(t, errors.As(err, &err3))
	require.Equal(t, "error1, error2, error3", fmt.Sprintf("%v", agg))
	require.Contains(t, fmt.Sprintf("%+v", agg), "error1, error2, error3\n")
	require.Equal(t, "error1, error2, error3", fmt.Sprintf("%s", agg))
	require.Equal(t, "\"error1, error2, error3\"", fmt.Sprintf("%q", agg))

	require.True(t, agg.Is(err1))
	require.True(t, agg.As(&err1))
	require.False(t, agg.Is(testError))
	require.False(t, agg.As(&testError))
}
