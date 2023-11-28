package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/samvimes01/go-rest1/models"
	"github.com/samvimes01/go-rest1/services"
)

func Signup(c *fiber.Ctx) error {
	var userInput *models.UserRequest = new(models.UserRequest)

	if err := c.BodyParser(userInput); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	errors := userInput.ValidateStruct()

	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "validation failed",
			Data:    errors,
		})
	}

	token, err := services.Signup(*userInput)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response[string]{
		Success: true,
		Message: "token data",
		Data:    token,
	})
}

func Login(c *fiber.Ctx) error {
	var userInput *models.UserRequest = new(models.UserRequest)

	if err := c.BodyParser(userInput); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	errors := userInput.ValidateStruct()

	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "validation failed",
			Data:    errors,
		})
	}

	token, err := services.Login(*userInput)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response[string]{
		Success: true,
		Message: "token data",
		Data:    token,
	})
}
