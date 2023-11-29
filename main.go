package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/samvimes01/go-rest1/db"
	"github.com/samvimes01/go-rest1/routes"
	"github.com/samvimes01/go-rest1/utils"
)

const DEFAULT_PORT = "8080"

func NewFiberApp() *fiber.App {
	var app *fiber.App = fiber.New()
	// dummy request handler
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world!")
	})
	routes.SetupRoutes(app)
	return app
}

func main() {
	var app *fiber.App = NewFiberApp()

	db.InitDatabase(utils.GetValue("DB_HOST"), utils.GetValue("DB_NAME"))

	var PORT string = os.Getenv("PORT")
	if PORT == "" {
		PORT = DEFAULT_PORT
	}

	app.Listen(fmt.Sprintf(":%s", PORT))
}
