package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		cfg := cors.Config{
			AllowOrigins:     "*",
			AllowMethods:     "GET,HEAD,POST,PUT,PATCH,DELETE,OPTIONS",
			AllowHeaders:     fmt.Sprintf("Content-Type,Accept,Authorization,Origin,%s", c.Get("Access-Control-Request-Headers")),
			ExposeHeaders:    "Content-Length,Content-Type",
			AllowCredentials: true,
			MaxAge:           int((12 * time.Hour).Seconds()),
		}

		return cors.New(cfg)(c)
	}
}
