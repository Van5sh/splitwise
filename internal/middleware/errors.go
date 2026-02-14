package middleware

import (
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
			appErr := helpers.NormalizeError(err)

			logger.Error("request error", map[string]interface{}{
				"code":    appErr.Code,
				"message": appErr.Message,
				"status":  appErr.Status,
				"path":    c.Path(),
				"method":  c.Method(),
			})

			return c.Status(appErr.Status).JSON(fiber.Map{
				"code":      appErr.Code,
				"message":   appErr.Message,
				"details":   appErr.Details,
				"timestamp": appErr.Timestamp,
			})
		}

		return nil
	}
}
