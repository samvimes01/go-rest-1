package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samvimes01/go-rest1/handlers"
	"github.com/samvimes01/go-rest1/middlewares"
)

func SetupRoutes(app *fiber.App) {
	// public routes
	var publicRoutes fiber.Router = app.Group("/api/v1")

	publicRoutes.Post("/signup", handlers.Signup)
	publicRoutes.Post("/login", handlers.Login)
	publicRoutes.Get("/items", handlers.GetAllItems)
	publicRoutes.Get("/items/:id", handlers.GetItemByID)

	// private routes, authentication is required
	var privateRoutes fiber.Router = app.Group("/api/v1", middlewares.CreateJwtMiddleware())

	privateRoutes.Post("/items", handlers.CreateItem)
	privateRoutes.Put("/items/:id", handlers.UpdateItem)
	privateRoutes.Delete("/items/:id", handlers.DeleteItem)
}
