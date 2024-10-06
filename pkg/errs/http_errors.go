package errs

import "net/http"

type HTTPError int

const (
	ErrBadRequest           HTTPError = http.StatusBadRequest
	ErrUnauthorized         HTTPError = http.StatusUnauthorized
	ErrForbidden            HTTPError = http.StatusForbidden
	ErrNotFound             HTTPError = http.StatusNotFound
	ErrMethodNotAllowed     HTTPError = http.StatusMethodNotAllowed
	ErrNotAcceptable        HTTPError = http.StatusNotAcceptable
	ErrConflict             HTTPError = http.StatusConflict
	ErrGone                 HTTPError = http.StatusGone
	ErrPreconditionFailed   HTTPError = http.StatusPreconditionFailed
	ErrUnsupportedMediaType HTTPError = http.StatusUnsupportedMediaType
	ErrUnprocessableEntity  HTTPError = http.StatusUnprocessableEntity
	ErrTooManyRequests      HTTPError = http.StatusTooManyRequests
	ErrInternal             HTTPError = http.StatusInternalServerError
	ErrNotImplemented       HTTPError = http.StatusNotImplemented
	ErrBadGateway           HTTPError = http.StatusBadGateway
	ErrServiceUnavailable   HTTPError = http.StatusServiceUnavailable
	ErrGatewayTimeout       HTTPError = http.StatusGatewayTimeout
)

func (he HTTPError) String() string {
	return httpCodeNames[he]
}

func (he HTTPError) HTTPStatus() int {
	return int(he)
}

var httpCodeNames = [...]string{
	ErrBadRequest:           "bad_request",
	ErrUnauthorized:         "unauthorized_access",
	ErrForbidden:            "forbidden",
	ErrNotFound:             "resource_not_found",
	ErrMethodNotAllowed:     "method_not_allowed",
	ErrNotAcceptable:        "not_acceptable",
	ErrConflict:             "conflict",
	ErrGone:                 "gone",
	ErrPreconditionFailed:   "precondition_failed",
	ErrUnsupportedMediaType: "unsupported_media_type",
	ErrUnprocessableEntity:  "unprocessable_entity",
	ErrTooManyRequests:      "too_many_requests",
	ErrInternal:             "internal_server_error",
	ErrNotImplemented:       "not_implemented",
	ErrBadGateway:           "bad_gateway",
	ErrServiceUnavailable:   "service_unavailable",
	ErrGatewayTimeout:       "gateway_timeout",
}
