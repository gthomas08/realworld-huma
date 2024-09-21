package errs

import (
	"fmt"

	"github.com/danielgtaylor/huma/v2"
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

func Error401Unauthorized() huma.StatusError {
	return huma.Error401Unauthorized("unauthorized", fmt.Errorf("unauthorized"))
}

func Error403Forbidden() huma.StatusError {
	return huma.Error403Forbidden("forbidden", fmt.Errorf("forbidden"))
}

func Error404NotFound() huma.StatusError {
	return huma.Error404NotFound("not found", fmt.Errorf("not found"))
}

func Error422UnprocessableEntity() huma.StatusError {
	return huma.Error422UnprocessableEntity("unprocessable entity", fmt.Errorf("unprocessable entity"))
}
