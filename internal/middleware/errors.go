package middleware

import "github.com/gofiber/fiber/v2"

type ErrorResponse struct {
	Code      string      `json:"code"`
	Message   string      `json:"message"`
	Details   interface{} `json:"details,omitempty"`
	TimeStamp int64       `json:"timestamp"`
	RequestID string      `json:"requestId,omitempty"`
}

type Logger interface {
	Error(msg string, fields ...map[string]interface{})
}

type ErrorHandler struct {
	ExposeErrorDetails bool
	Logger             Logger
}

func ErrorResponseMiddleware() fiber.Handler {
	var errorHandler *ErrorHandler
}
