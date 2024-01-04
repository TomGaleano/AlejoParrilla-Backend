package controller

import (
	"github.com/TomGaleano/Golang/database"
	"github.com/TomGaleano/Golang/models"
	"github.com/gofiber/fiber/v2"
)

func CreateMenu(c *fiber.Ctx) error {
	var menu models.Menu
	if err := c.BodyParser(&menu); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Unable to parse request body",
			"error":   err.Error(),
		})
	}
	if err := database.DB_SEC.Create(&menu).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Unable to create menu",
			"error":   err.Error(),
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "Menu created successfully",
		"menu":    menu,
	})
}

// Hacer peticiones de estilo /api/allmenu y obtiene todos los platos
func AllMenu(c *fiber.Ctx) error {
	var menus []models.Menu
	if err := database.DB_SEC.Find(&menus).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error retrieving menus",
		})
	}
	return c.JSON(fiber.Map{
		"data": menus,
	})
}

// Hacer peticiones de estilo /api/filtermenu?category=Carnes para obtener todos los platos de la categoria Carnes
func FilterMenu(c *fiber.Ctx) error {
	category := c.Query("category")
	var menus []models.Menu
	if err := database.DB_SEC.Where("category = ?", category).Find(&menus).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error retrieving menus",
		})
	}
	return c.JSON(fiber.Map{
		"data": menus,
	})
}
