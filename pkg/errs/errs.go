package errs

import (
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

var (
	ErrEntityExists = fmt.Errorf("entity_exists")

	ErrBadRequest           = fmt.Errorf("bad_request")
	ErrUnauthorized         = fmt.Errorf("unauthorized_access")
	ErrForbidden            = fmt.Errorf("forbidden")
	ErrNotFound             = fmt.Errorf("resource_not_found")
	ErrMethodNotAllowed     = fmt.Errorf("method_not_allowed")
	ErrNotAcceptable        = fmt.Errorf("not_acceptable")
	ErrConflict             = fmt.Errorf("conflict")
	ErrGone                 = fmt.Errorf("gone")
	ErrPreconditionFailed   = fmt.Errorf("precondition_failed")
	ErrUnsupportedMediaType = fmt.Errorf("unsupported_media_type")
	ErrUnprocessableEntity  = fmt.Errorf("unprocessable_entity")
	ErrTooManyRequests      = fmt.Errorf("too_many_requests")
	ErrInternal             = fmt.Errorf("internal_server_error")
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

func NewAppError(err error, detail string, errs ...error) huma.StatusError {
	switch err {
	case ErrEntityExists:
		return &huma.ErrorModel{
			Status: http.StatusConflict,
			Title:  err.Error(),
			Detail: detail,
		}
	case ErrInternal:
		return &huma.ErrorModel{
			Status: http.StatusInternalServerError,
			Title:  err.Error(),
			Detail: detail,
		}
	default:
		return &huma.ErrorModel{
			Status: http.StatusInternalServerError,
			Title:  ErrInternal.Error(),
			Detail: ErrInternal.Error(),
		}
	}
}
