package database

import (
	"log"
	"os"

	"github.com/TomGaleano/Golang/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DB_SEC *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
	dsn := os.Getenv("DSN")
	database, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to the database.")
	} else {
		log.Println("Connected succesfully to GoBlog database.")
	}
	DB = database
	database.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)
}
func ConnectSec() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
	dsn_sec := os.Getenv("DSN_SEC")
	database_sec, _ := gorm.Open(mysql.Open(dsn_sec), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to the database.")
	} else {
		log.Println("Connected succesfully to AlejoParrilla database.")
	}
	DB_SEC = database_sec
	database_sec.AutoMigrate(
		&models.Menu{},
	)
}
