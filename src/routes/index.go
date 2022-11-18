package routes

import (
	"github.com/KaioMarxDEV/gofinance/src/User/UsersRoutes"
	"github.com/gofiber/fiber/v2"
)

// creating the router central for instances of all prefix inside the API
func MainRouter(api fiber.Router) {
	api.Route("/users", UsersRoutes.Router, "users")
}
