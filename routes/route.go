package routes

import (
	"github.com/TomGaleano/Golang/controller"
	"github.com/TomGaleano/Golang/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Setup(app *fiber.App) {
	app.Use(cors.New(cors.Config{
        AllowOrigins:     "http://localhost:5173, https://fictional-space-memory-9qrr64w7v67cgxg-5173.app.github.dev, https://alejoparrilla-backend-temp-c472e01a4c9d.herokuapp.com",
        AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
        AllowHeaders:     "Origin, Content-Type, Accept",
        AllowCredentials: true,
    }))

	//AlejoParrilla APIs
	app.Get("/api/allmenu", controller.AllMenu)
	app.Get("/api/filtermenu", controller.FilterMenu)
	app.Post("/api/createbooking", controller.CreateBooking)
	app.Post("/api/collectemail", controller.CollectEmail)

	//Cooler APIs
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	app.Use(middleware.IsAuthenticated)

	app.Post("/api/createpost", controller.CreatePost)
	app.Get("/api/allpost", controller.AllPost)
	app.Get("/api/detailpost/:id", controller.DetailPost)
	app.Put("/api/updatepost/:id", controller.UpdatePost)
	app.Get("/api/uniquepost", controller.UniquePost)
	app.Delete("/api/deletepost/:id", controller.DeletePost)
	app.Post("/api/upload", controller.Upload)
	app.Static("/api/uploads", "./uploads")
}
