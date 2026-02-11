package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

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

	return func(c *fiber.Ctx) error {
		requestID := c.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
			c.Set("X-Request-ID", requestID)
		}

		defer func() {
			if r := recover(); r != nil {
				if errorHandler != nil && errorHandler.Logger != nil {
					errorHandler.Logger.Error("Panic recovered in ErrorResponseMiddleware",
						map[string]interface{}{
							"error":     r,
							"requestId": requestID,
						})
				}

				_ = c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
					Code:      "internal_server_error",
					Message:   "An unexpected error occurred.",
					TimeStamp: time.Now().Unix(),
					RequestID: requestID,
				})
				return
			}

			if err := c.Context().Err(); err != nil {
				if errorHandler != nil && errorHandler.Logger != nil {
					errorHandler.Logger.Error("Error occurred during request processing",
						map[string]interface{}{
							"error":     err.Error(),
							"requestId": requestID,
						})
				}

				_ = c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
					Code:      "internal_server_error",
					Message:   "An unexpected error occurred.",
					TimeStamp: time.Now().Unix(),
					RequestID: requestID,
				})
				return
			}

			status := c.Response().StatusCode()

			if status >= 500 {
				if errorHandler != nil && errorHandler.Logger != nil {
					errorHandler.Logger.Error("Request resulted in server error",
						map[string]interface{}{
							"statusCode": status,
							"requestId":  requestID,
						})
				}

				_ = c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
					Code:      "internal_server_error",
					Message:   "An unexpected error occurred.",
					TimeStamp: time.Now().Unix(),
					RequestID: requestID,
				})
				return
			}

			if status >= 400 {
				if errorHandler != nil && errorHandler.Logger != nil {
					errorHandler.Logger.Error("Request resulted in client error",
						map[string]interface{}{
							"statusCode": status,
							"requestId":  requestID,
						})
				}
			}
		}()

		return c.Next()
	}
}

func statusCodeToErrorCode(status int) string {
	switch status {
	case fiber.StatusBadRequest:
		return "BAD_REQUEST"
	case fiber.StatusUnauthorized:
		return "UNAUTHORIZED"
	case fiber.StatusForbidden:
		return "FORBIDDEN"
	case fiber.StatusNotFound:
		return "NOT_FOUND"
	case fiber.StatusMethodNotAllowed:
		return "METHOD_NOT_ALLOWED"
	case fiber.StatusConflict:
		return "CONFLICT"
	case fiber.StatusTooManyRequests:
		return "RATE_LIMITED"
	default:
		if status >= 500 {
			return "SERVER_ERROR"
		}
		return "CLIENT_ERROR"
	}
}
