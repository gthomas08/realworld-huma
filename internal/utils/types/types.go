package types

// RequestBody is a generic struct for the request body.
type RequestBody[T any] struct {
	Body T
}

// ResponseBody is a generic struct for the response body.
type ResponseBody[T any] struct {
	Body T
}
