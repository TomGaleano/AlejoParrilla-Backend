package main

import (
	"log"
	"os"

	"github.com/TomGaleano/Golang/database"
	"github.com/TomGaleano/Golang/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	database.ConnectSec()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files.")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":" + port)
}
