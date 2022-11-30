package routes

import (
	"github.com/KaioMarxDEV/gofinance/cmd/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// first endpoint to be considered on all subsequent calls "/api/BLA/FOO/BOO"
	api := app.Group("/api", logger.New())
	// routing with hanlder.hello PING type function to verify API health
	api.Get("/", handler.Hello)

	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// TODO: create authentication route and middleware to ensure it

	// TODO: create static route to serve React compiled app

	// user Routes and Handlers
	user := api.Group("/user")
	user.Post("/", handler.CreateUser)
	user.Get("/", handler.GetAllUsers)
	user.Get("/:id", handler.GetUserByID)
}
