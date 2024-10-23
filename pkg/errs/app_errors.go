package errs

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type AppError int

const (
	InvalidCredentials AppError = iota + 4000
	EntityExists
	EntityNotFound
	InvalidOperation
)

func (ae AppError) String() string {
	return appCodeNames[ae]
}

func (ae AppError) HTTPStatus() int {
	return appCodeStatus[ae]
}

var appCodeNames = [...]string{
	InvalidCredentials: "invalid_credentials",
	EntityExists:       "entity_exists",
	EntityNotFound:     "entity_not_found",
	InvalidOperation:   "invalid_operation",
}

var appCodeStatus = [...]int{
	InvalidCredentials: http.StatusUnauthorized,
	EntityExists:       http.StatusConflict,
	EntityNotFound:     http.StatusNotFound,
	InvalidOperation:   http.StatusConflict,
}

// NewAppError returns an error with the specified error code and message.
func NewAppError(appErr AppError, message string) huma.StatusError {
	return &ErrorResponse{
		Status:  appErr.HTTPStatus(),
		Code:    appErr.String(),
		Message: message,
	}
}
