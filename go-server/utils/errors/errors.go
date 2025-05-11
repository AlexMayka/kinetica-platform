package errors

import (
	"fmt"
)

type ErrorResponse struct {
	Code    int
	Message string
	Tag     string
}

var registry = make(map[error]ErrorResponse)

func Registration(err error, resp ErrorResponse) {
	registry[err] = resp
}

func MapError(err error) ErrorResponse {
	if resp, ok := registry[err]; ok {
		return resp
	}
	return ErrorResponse{Code: 9999, Message: "Unknown error occurred", Tag: "unknown"}
}

func PrintError(err error) {
	e := MapError(err)
	fmt.Printf("❌ [%s] %s (code: %d)\n", e.Tag, e.Message, e.Code)
}
