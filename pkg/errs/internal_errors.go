package errs

import "errors"

var (
	ErrNotFound = errors.New("entity not found")
)
