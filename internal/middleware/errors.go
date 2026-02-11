package middleware

import (
	"fmt"
	"time"

	"github.com/Van5sh/new-splitwise/internal/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Logger interface {
	Error(msg string, fields map[string]interface{})
}

type defaultLogger struct{}

func (l defaultLogger) Error(msg string, fields map[string]interface{}) {
	fmt.Printf("[ERROR] %s | %v\n", msg, fields)
}

func ErrorResponseMiddleware() fiber.Handler {
	logger := defaultLogger{}

	return func(c *fiber.Ctx) (err error) {
		requestID := c.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
			c.Set("X-Request-ID", requestID)
		}

		defer func() {
			if r := recover(); r != nil {

				logger.Error("panic recovered", map[string]interface{}{
					"error":     r,
					"requestId": requestID,
					"path":      c.Path(),
					"method":    c.Method(),
				})

				_ = c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"code":      "internal_server_error",
					"message":   "Something went wrong",
					"timestamp": time.Now().Unix(),
					"requestId": requestID,
				})
			}
		}()

		err = c.Next()

		if err != nil {

			// HANDLE CUSTOM APP ERROR
			if appErr, ok := err.(*helpers.ErrorsResponse); ok {

				logger.Error("application error", map[string]interface{}{
					"code":      appErr.Code,
					"message":   appErr.Message,
					"status":    appErr.Status,
					"requestId": requestID,
					"path":      c.Path(),
					"method":    c.Method(),
				})

				return c.Status(appErr.Status).JSON(fiber.Map{
					"code":      appErr.Code,
					"message":   appErr.Message,
					"details":   appErr.Details,
					"timestamp": time.Now().Unix(),
					"requestId": requestID,
				})
			}

			// HANDLE UNKNOWN ERROR
			logger.Error("unhandled error", map[string]interface{}{
				"error":     err.Error(),
				"requestId": requestID,
				"path":      c.Path(),
				"method":    c.Method(),
			})

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code":      "internal_server_error",
				"message":   "Internal Server Error",
				"timestamp": time.Now().Unix(),
				"requestId": requestID,
			})
		}

		return nil
	}
}
