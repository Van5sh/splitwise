package main

import (
	"github.com/Van5sh/new-splitwise/database"
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
	app.Listen(":8080")
}
