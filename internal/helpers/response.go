package helpers

import "github.com/gofiber/fiber/v2"

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *fiber.Ctx, status string, message string, data interface{}) error {
	return c.JSON(APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func Error(c *fiber.Ctx, statusCode int, message string) error {
	c.Status(statusCode)
	return c.JSON(APIResponse{
		Status:  "error",
		Message: message,
	})
}
