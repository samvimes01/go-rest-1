package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/samvimes01/go-rest1/models"
	"github.com/samvimes01/go-rest1/services"
)

func GetAllItems(c *fiber.Ctx) error {
	var items []models.Item = services.GetAllItems()

	return c.JSON(models.Response[[]models.Item]{
		Success: true,
		Message: "All items data",
		Data:    items,
	})
}

func GetItemByID(c *fiber.Ctx) error {
	var itemID string = c.Params("id")

	item, err := services.GetItemByID(itemID)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response[models.Item]{
		Success: true,
		Message: "item found",
		Data:    item,
	})
}

func CreateItem(c *fiber.Ctx) error {
	var itemInput *models.ItemRequest = new(models.ItemRequest)

	if err := c.BodyParser(itemInput); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	errors := itemInput.ValidateStruct()

	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "validation failed",
			Data:    errors,
		})
	}

	var createdItem models.Item = services.CreateItem(itemInput)

	return c.Status(http.StatusCreated).JSON(models.Response[models.Item]{
		Success: true,
		Message: "item created",
		Data:    createdItem,
	})
}

func UpdateItem(c *fiber.Ctx) error {
  itemInput := &models.ItemRequest{}

	if err := c.BodyParser(itemInput); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	errors := itemInput.ValidateStruct()

	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "validation failed",
			Data:    errors,
		})
	}

	var itemID string = c.Params("id")

	updatedItem, err := services.UpdateItem(itemInput, itemID)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response[models.Item]{
		Success: true,
		Message: "item updated",
		Data:    updatedItem,
	})
}

func DeleteItem(c *fiber.Ctx) error {
	var itemID string = c.Params("id")

	result := services.DeleteItem(itemID)

  if result {
    return c.JSON(models.Response[any]{
      Success: true,
      Message: "item deleted",
    })
  }

  return c.Status(http.StatusNotFound).JSON(models.Response[any]{
		Success: false,
		Message: "item failed to delete",
	})
}
