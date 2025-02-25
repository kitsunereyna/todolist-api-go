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

func GetAllTowatch(c *fiber.Ctx) error {
	towatches := []Towatch{}
	database.DB.Model(&models.Towatch{}).Find(&towatches)

	return helpers.SendSuccessJSON(c, towatches)
}

func GetTowatchById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID Format",
		})
	}

	towatch := Towatch{}
	query := Towatch{ID: id}
	err = database.DB.First(&towatch, &query).Error

	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    404,
			"message": "Todo not found",
		})
	}

	return helpers.SendSuccessJSON(c, towatch)
}

func CreateTowatch(c *fiber.Ctx) error {
	json := new(dto.CreateTowatchDto)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}

	newTowatch := Towatch{
		Title:      json.Title,
		StartDate:  json.StartDate,
		FinishDate: json.FinishDate,
		Episodes:   json.Episodes,
		Rating:     json.Rating,
		Studio:     json.Studio,
		Image:      json.Image,
	}

	err := database.DB.Create(&newTowatch).Error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return helpers.SendSuccessJSON(c, newTowatch)
}

func DeleteTowatchById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid ID format",
		})
	}

	foundTowatch := Towatch{}
	query := Towatch{
		ID: id,
	}

	err = database.DB.First(&foundTowatch, &query).Error
	if err == gorm.ErrRecordNotFound {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Towatch not found",
		})
	}

	database.DB.Delete(&foundTowatch)
	return helpers.SendSuccessJSON(c, nil)
}
