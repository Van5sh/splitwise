package middleware

import (
	"errors"
	"fmt"
	"time"

	"github.com/Van5sh/new-splitwise/internal/helpers"
	"github.com/gofiber/fiber/v2"
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
		defer func() {
			if r := recover(); r != nil {
				logger.Error("panic recovered", map[string]interface{}{
					"error": r,
					// "requestId": requestID,
					"path":   c.Path(),
					"method": c.Method(),
				})

				_ = c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"code":      "internal_server_error",
					"message":   "Something went wrong",
					"timestamp": time.Now().UTC().Format(time.RFC3339),
					// "requestId": requestID,
				})
			}
		}()

		err = c.Next()

		if err != nil {
			var appErr *helpers.ErrorsResponse
			if errors.As(err, &appErr) {

				logger.Error("application error", map[string]interface{}{
					"code":    appErr.Code,
					"message": appErr.Message,
					"status":  appErr.Status,
					// "requestId": requestID,
					"path":   c.Path(),
					"method": c.Method(),
				})

				return c.Status(appErr.Status).JSON(fiber.Map{
					"code":      appErr.Code,
					"message":   appErr.Message,
					"details":   appErr.Details,
					"timestamp": appErr.Timestamp,
					// "requestId": requestID,
				})
			}

			// HANDLE FIBER HTTP ERROR
			if fErr, ok := err.(*fiber.Error); ok {
				logger.Error("fiber error", map[string]interface{}{
					"code":   fErr.Code,
					"error":  fErr.Message,
					"path":   c.Path(),
					"method": c.Method(),
				})

				return c.Status(fErr.Code).JSON(fiber.Map{
					"code":      "http_error",
					"message":   fErr.Message,
					"timestamp": time.Now().UTC().Format(time.RFC3339),
				})
			}

			// HANDLE UNKNOWN ERROR
			logger.Error("unhandled error", map[string]interface{}{
				"error": err.Error(),
				// "requestId": requestID,
				"path":   c.Path(),
				"method": c.Method(),
			})

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code":      "internal_server_error",
				"message":   "Internal Server Error",
				"timestamp": time.Now().UTC().Format(time.RFC3339),
				// "requestId": requestID,
			})
		}

		return nil
	}
}
