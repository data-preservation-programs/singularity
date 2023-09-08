package handlererror

import (
	"github.com/cockroachdb/errors"
)

var ErrInvalidParameter = errors.New("invalid parameter")

var ErrNotFound = errors.New("not found")

var ErrDuplicateRecord = errors.New("duplicate record")
