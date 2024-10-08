package errs

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type AppError int

const (
	EntityExists AppError = iota + 4000
)

func (ae AppError) String() string {
	return appCodeNames[ae]
}

func (ae AppError) HTTPStatus() int {
	return appCodeStatus[ae]
}

var appCodeNames = [...]string{
	EntityExists: "entity_exists",
}

var appCodeStatus = [...]int{
	EntityExists: http.StatusConflict,
}

// NewAppError returns an error with the specified error code and message.
func NewAppError(appErr AppError, message string) huma.StatusError {
	return &ErrorResponse{
		Status:  appErr.HTTPStatus(),
		Code:    appErr.String(),
		Message: message,
	}
}
