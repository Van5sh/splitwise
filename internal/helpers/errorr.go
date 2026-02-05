package helpers

import (
	"fmt"
	"time"
)

type ErrorsResponse struct {
	Code      string      `json:"code"`
	Message   string      `json:"message"`
	Details   interface{} `json:"details,omitempty"`
	Status    int         `json:"status"`
	Timestamp string      `json:"timestamp"`
	Err       error       `json:"error"`
}

func (e *ErrorsResponse) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *ErrorsResponse) Unwrap() error {
	return e.Err
}

func NewAppError(code string, message string, details interface{}, status int, err error) *ErrorsResponse {
	return &ErrorsResponse{
		Code:      code,
		Message:   message,
		Details:   details,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Status:    status,
		Err:       err,
	}
}

func (e *ErrorsResponse) BadRequest(message string, details interface{}, err error) *ErrorsResponse {
	return NewAppError("bad_request", message, details, 400, err)
}

func (e *ErrorsResponse) ValidationError(message string, details interface{}, err error) *ErrorsResponse {
	return NewAppError("validation_error", message, details, 422, err)
}

func (e *ErrorsResponse) NotFound(message string, details interface{}, err error) *ErrorsResponse {
	return NewAppError("not_found", message, details, 404, err)
}

func (e *ErrorsResponse) InternalServerError(message string, details interface{}, err error) *ErrorsResponse {
	return NewAppError("internal_server_error", message, details, 500, err)
}

func (e *ErrorsResponse) Unauthorized(message string, details interface{}, err error) *ErrorsResponse {
	return NewAppError("unauthorized", message, details, 401, err)
}
