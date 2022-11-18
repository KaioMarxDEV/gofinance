package useCases

import (
	"github.com/KaioMarxDEV/gofinance/src/User/repositories"
	"github.com/gofiber/fiber/v2"
)

func FindUserByNameController(c *fiber.Ctx) error {
	name := c.Params("username")

	user, err := repositories.FindByName(name)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Error": err,
		})
	}

	return c.Status(200).JSON(user)
}
