package controller

import (
	"github.com/TomGaleano/Golang/database"
	"github.com/TomGaleano/Golang/models"
	"github.com/gofiber/fiber/v2"
)

func CreateBooking(c *fiber.Ctx) error {
	var booking models.Booking
	if err := c.BodyParser(&booking); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Unable to parse request body",
			"error":   err.Error(),
		})
	}
	if err := database.DB_SEC.Create(&booking).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Unable to create booking",
			"error":   err.Error(),
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "Booking created successfully",
		"booking": booking,
	})
}
