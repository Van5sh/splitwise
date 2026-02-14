package middleware

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

const defaultSlowRequestMS = 500

func RequestTimingMiddleware() fiber.Handler {
	threshold := time.Duration(getSlowRequestMS()) * time.Millisecond

	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		elapsed := time.Since(start)

		if elapsed >= threshold {
			fmt.Printf("[SLOW] %s %s | status=%d | duration=%s\n",
				c.Method(),
				c.Path(),
				c.Response().StatusCode(),
				elapsed.String(),
			)
		}

		return err
	}
}

func getSlowRequestMS() int {
	raw := os.Getenv("SLOW_REQUEST_MS")
	if raw == "" {
		return defaultSlowRequestMS
	}

	value, err := strconv.Atoi(raw)
	if err != nil || value <= 0 {
		return defaultSlowRequestMS
	}

	return value
}
