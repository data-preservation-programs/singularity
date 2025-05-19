package handlererror

import (
	"errors"
	stderrors "errors"
	"strings"
)

var ErrInvalidParameter = stderrors.New("invalid parameter")
var ErrNotFound = stderrors.New("not found")
var ErrDuplicateRecord = stderrors.New("duplicate record")
var ErrDuplicateKey = stderrors.New("duplicated key not allowed")

// IsDuplicateKeyError matches even if the error is deeply wrapped
func IsDuplicateKeyError(err error) bool {
	for err != nil {
		if errors.Is(err, ErrDuplicateKey) ||
			strings.Contains(err.Error(), "duplicate key value") ||
			strings.Contains(err.Error(), "duplicated key not allowed") {
			return true
		}
		err = errors.Unwrap(err)
	}
	return false
}
