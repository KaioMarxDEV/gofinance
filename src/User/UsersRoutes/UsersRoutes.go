package UsersRoutes

import (
	"github.com/KaioMarxDEV/gofinance/src/User/useCases"
	"github.com/gofiber/fiber/v2"
)

func Router(api fiber.Router) {
	api.Post("/create", useCases.CreateUserController).Name("CreateUser")

	api.Get("/find/:username", useCases.FindUserByNameController).Name("FindUserByName")
}
