package errs

import (
	"fmt"

	"github.com/danielgtaylor/huma/v2"
)

var (
	ErrEntityExists = fmt.Errorf("entity exists")

	ErrBadRequest           = fmt.Errorf("bad request")
	ErrUnauthorized         = fmt.Errorf("unauthorized access")
	ErrForbidden            = fmt.Errorf("forbidden")
	ErrNotFound             = fmt.Errorf("resource not found")
	ErrMethodNotAllowed     = fmt.Errorf("method not allowed")
	ErrNotAcceptable        = fmt.Errorf("not acceptable")
	ErrConflict             = fmt.Errorf("conflict")
	ErrGone                 = fmt.Errorf("gone")
	ErrPreconditionFailed   = fmt.Errorf("precondition failed")
	ErrUnsupportedMediaType = fmt.Errorf("unsupported media type")
	ErrUnprocessableEntity  = fmt.Errorf("unprocessable entity")
	ErrTooManyRequests      = fmt.Errorf("too many requests")
	ErrInternal             = fmt.Errorf("internal server error")
)

type Errors struct {
	Body []string `json:"body"`
}

type ErrorResponse struct {
	status  int
	Message string `json:"message"`
	Errors  Errors `json:"errors"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
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
		Message: message,
		Errors: Errors{
			Body: errors,
		},
	}
}

func NewAppError(err error) huma.StatusError {
	switch err {
	case ErrEntityExists:
		return huma.Error409Conflict(err.Error())
	default:
		return huma.Error500InternalServerError(ErrInternal.Error())
	}
}

// NewHTTPError converts an application error to a StatusError for Huma.
func NewHTTPError(err error) huma.StatusError {
	switch err {
	case ErrBadRequest:
		return huma.Error400BadRequest(err.Error())
	case ErrUnauthorized:
		return huma.Error401Unauthorized(err.Error())
	case ErrForbidden:
		return huma.Error403Forbidden(err.Error())
	case ErrNotFound:
		return huma.Error404NotFound(err.Error())
	case ErrMethodNotAllowed:
		return huma.Error405MethodNotAllowed(err.Error())
	case ErrNotAcceptable:
		return huma.Error406NotAcceptable(err.Error())
	case ErrConflict:
		return huma.Error409Conflict(err.Error())
	case ErrGone:
		return huma.Error410Gone(err.Error())
	case ErrPreconditionFailed:
		return huma.Error412PreconditionFailed(err.Error())
	case ErrUnsupportedMediaType:
		return huma.Error415UnsupportedMediaType(err.Error())
	case ErrUnprocessableEntity:
		return huma.Error422UnprocessableEntity(err.Error())
	case ErrTooManyRequests:
		return huma.Error429TooManyRequests(err.Error())
	case ErrInternal:
		return huma.Error500InternalServerError(err.Error())
	default:
		return huma.Error500InternalServerError(ErrInternal.Error())
	}
}
