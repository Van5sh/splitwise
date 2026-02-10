package main

import (
	database "github.com/Van5sh/new-splitwise/internal/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := database.ConnectToDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})
	go startServer()
	app.Listen(":8080")
}
