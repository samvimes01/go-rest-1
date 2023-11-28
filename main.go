package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samvimes01/go-rest1/routes"
)

func main() {
	var app *fiber.App = fiber.New()

	// dummy request handler 
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world!")
	})

  routes.SetupRoutes(app)

	app.Listen(":8080")
}