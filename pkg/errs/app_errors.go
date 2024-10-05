package errs

import "net/http"

type AppError int

const (
	ErrEntityExists AppError = iota + 4000
)

func (ae AppError) String() string {
	return appCodeNames[ae]
}

func (ae AppError) HTTPStatus() int {
	return appCodeStatus[ae]
}

var appCodeNames = [...]string{
	ErrEntityExists: "entity_exists",
}

var appCodeStatus = [...]int{
	ErrEntityExists: http.StatusConflict,
}
