package controllers

import (
	"strconv"

	"github.com/DiarCode/todo-go-api/src/config/database"
	"github.com/DiarCode/todo-go-api/src/dto"
	"github.com/DiarCode/todo-go-api/src/helpers"
	"github.com/DiarCode/todo-go-api/src/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllTowatchCategories(c *fiber.Ctx) error {
	categories := []TowatchCategory{}
	database.DB.Model(&models.TowatchCategory{}).Find(&categories)

	return helpers.SendSuccessJSON(c, categories)
}

func GetTowatchCategoryById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID Format",
		})
	}

	category := TowatchCategory{}
	query := TowatchCategory{ID: id}
	err = database.DB.First(&category, &query).Error

	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Towatch category not found",
		})
	}

	return helpers.SendSuccessJSON(c, category)
}

func CreateTowatchCategory(c *fiber.Ctx) error {
	json := new(dto.CreateTowatchCategoryDto)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	newCategory := TowatchCategory{
		Value: json.Value,
		Color: json.Color,
	}

	err := database.DB.Create(&newCategory).Error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return helpers.SendSuccessJSON(c, newCategory)
}

func DeleteTowatchCategoryById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID format",
		})
	}

	foundCategory := TowatchCategory{}
	query := TowatchCategory{
		ID: id,
	}

	err = database.DB.First(&foundCategory, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Towatch category not found",
		})
	}

	database.DB.Delete(&foundCategory)
	return helpers.SendSuccessJSON(c, nil)
}
