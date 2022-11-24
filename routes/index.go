package routes

import (
	"github.com/KaioMarxDEV/gofinance/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// user routes
	user := api.Group("/user")
	user.Post("/", handler.CreateUser)
}
