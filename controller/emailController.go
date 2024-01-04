package controller

import (
	"github.com/TomGaleano/Golang/database"
	"github.com/TomGaleano/Golang/models"
	"github.com/gofiber/fiber/v2"
)

func CollectEmail(c *fiber.Ctx) error {
	var email models.Email
	if err := c.BodyParser(&email); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Unable to parse request body",
			"error":   err.Error(),
		})
	}
	if err := database.DB_SEC.Create(&email).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Unable to collect email",
			"error":   err.Error(),
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "Email collected successfully",
		"email":   email,
	})
}
