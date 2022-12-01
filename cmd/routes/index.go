package routes

import (
	"github.com/KaioMarxDEV/gofinance/cmd/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// first endpoint to be considered on all subsequent calls "/api/BLA/FOO/BOO"
	api := app.Group("/api", logger.New())

	// authentication route
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// user Routes and Handlers
	user := api.Group("/user")
	user.Post("/", handler.CreateUser)
	// TODO: implement admin authenticate middleware to all below routes
	user.Get("/", handler.GetAllUsers)
	user.Get("/:id", handler.GetUserByID)
}
