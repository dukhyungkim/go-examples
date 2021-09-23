package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := setup()

	if err := app.Listen(":3000"); err != nil {
		log.Fatalln(err)
	}
}

func setup() *fiber.App {
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	return app
}
