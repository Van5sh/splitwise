package helpers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ErrorsResponse struct {
	Code      string      `json:"code"`
	Message   string      `json:"message"`
	Details   interface{} `json:"details,omitempty"`
	Status    int         `json:"status"`
	Timestamp string      `json:"timestamp"`
	Err       error       `json:"-"`
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

func newAppError(code, message string, details interface{}, status int, err error) *ErrorsResponse {
	if status <= 0 {
		status = http.StatusInternalServerError
	}
	code = strings.TrimSpace(code)
	if code == "" {
		code = "internal_server_error"
	}
	message = strings.TrimSpace(message)
	if message == "" {
		message = http.StatusText(status)
	}

	return &ErrorsResponse{
		Code:      code,
		Message:   message,
		Details:   details,
		Status:    status,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Err:       err,
	}
}

func BadRequest(message string, details interface{}, err error) *ErrorsResponse {
	return newAppError("bad_request", message, details, 400, err)
}

func ValidationError(message string, details interface{}, err error) *ErrorsResponse {
	return newAppError("validation_error", message, details, 422, err)
}

func NotFound(message string, details interface{}, err error) *ErrorsResponse {
	return newAppError("not_found", message, details, 404, err)
}

func InternalServerError(message string, details interface{}, err error) *ErrorsResponse {
	return newAppError("internal_server_error", message, details, 500, err)
}

func Unauthorized(message string, details interface{}, err error) *ErrorsResponse {
	return newAppError("unauthorized", message, details, 401, err)
}

func Forbidden(message string, details interface{}, err error) *ErrorsResponse {
	return newAppError("forbidden", message, details, 403, err)
}

func NormalizeError(err error) *ErrorsResponse {
	if err == nil {
		return nil
	}

	var appErr *ErrorsResponse
	if errors.As(err, &appErr) {
		return appErr
	}

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		return newAppError(
			"http_error",
			fiberErr.Message,
			nil,
			fiberErr.Code,
			err,
		)
	}

	if errors.Is(err, sql.ErrNoRows) {
		return NotFound("Resource not found", nil, err)
	}

	// Some wrapped DB drivers lose sql.ErrNoRows identity but keep message text.
	if strings.Contains(strings.ToLower(err.Error()), "no rows in result set") {
		return NotFound("Resource not found", nil, err)
	}

	return InternalServerError("Internal Server Error", nil, err)
}
