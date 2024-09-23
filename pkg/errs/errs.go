package errs

import (
	"fmt"

	"github.com/danielgtaylor/huma/v2"
)

var (
	ErrUnauthorized = fmt.Errorf("unauthorized access")
	ErrBadRequest   = fmt.Errorf("bad request")
	ErrNotFound     = fmt.Errorf("resource not found")
	ErrInternal     = fmt.Errorf("internal server error")
)

type Errors struct {
	Body []string `json:"body"`
}

type ErrorResponse struct {
	status  int
	message string
	Errors  Errors `json:"errors"`
}

func (e *ErrorResponse) Error() string {
	return e.message
}

func (e *ErrorResponse) GetStatus() int {
	return e.status
}

func NewError(status int, message string, errs ...error) huma.StatusError {
	errors := make([]string, len(errs))
	for i, err := range errs {
		errors[i] = err.Error()
	}
	return &ErrorResponse{
		status:  status,
		message: message,
		Errors: Errors{
			Body: errors,
		},
	}
}

// NewAppError converts an application error to a StatusError for Huma.
func NewAppError(err error) huma.StatusError {
	switch err {
	case ErrUnauthorized:
		return huma.Error401Unauthorized(err.Error())
	case ErrBadRequest:
		return huma.Error400BadRequest(err.Error())
	case ErrNotFound:
		return huma.Error404NotFound(err.Error())
	case ErrInternal:
		return huma.Error500InternalServerError(err.Error())
	default:
		return huma.Error500InternalServerError(ErrInternal.Error())
	}
}
