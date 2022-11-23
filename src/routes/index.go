package routes

import (
	"github.com/KaioMarxDEV/gofinance/src/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	api.Get("/", handler.Hello)

	// TODO:  CREATE ROUTES FOR USERS RESOURCES
}
