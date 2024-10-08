package errs

import "net/http"

type HTTPError int

const (
	BadRequest           HTTPError = http.StatusBadRequest
	Unauthorized         HTTPError = http.StatusUnauthorized
	Forbidden            HTTPError = http.StatusForbidden
	NotFound             HTTPError = http.StatusNotFound
	MethodNotAllowed     HTTPError = http.StatusMethodNotAllowed
	NotAcceptable        HTTPError = http.StatusNotAcceptable
	Conflict             HTTPError = http.StatusConflict
	Gone                 HTTPError = http.StatusGone
	PreconditionFailed   HTTPError = http.StatusPreconditionFailed
	UnsupportedMediaType HTTPError = http.StatusUnsupportedMediaType
	UnprocessableEntity  HTTPError = http.StatusUnprocessableEntity
	TooManyRequests      HTTPError = http.StatusTooManyRequests
	Internal             HTTPError = http.StatusInternalServerError
	NotImplemented       HTTPError = http.StatusNotImplemented
	BadGateway           HTTPError = http.StatusBadGateway
	ServiceUnavailable   HTTPError = http.StatusServiceUnavailable
	GatewayTimeout       HTTPError = http.StatusGatewayTimeout
)

func (he HTTPError) String() string {
	return httpCodeNames[he]
}

func (he HTTPError) HTTPStatus() int {
	return int(he)
}

var httpCodeNames = [...]string{
	BadRequest:           "bad_request",
	Unauthorized:         "unauthorized_access",
	Forbidden:            "forbidden",
	NotFound:             "resource_not_found",
	MethodNotAllowed:     "method_not_allowed",
	NotAcceptable:        "not_acceptable",
	Conflict:             "conflict",
	Gone:                 "gone",
	PreconditionFailed:   "precondition_failed",
	UnsupportedMediaType: "unsupported_media_type",
	UnprocessableEntity:  "unprocessable_entity",
	TooManyRequests:      "too_many_requests",
	Internal:             "internal_server_error",
	NotImplemented:       "not_implemented",
	BadGateway:           "bad_gateway",
	ServiceUnavailable:   "service_unavailable",
	GatewayTimeout:       "gateway_timeout",
}
