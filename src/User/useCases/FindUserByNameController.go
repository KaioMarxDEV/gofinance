package useCases

import (
	"github.com/gofiber/fiber/v2"
)

func FindUserByNameController(c *fiber.Ctx) error {

	return c.SendStatus(200)
}
