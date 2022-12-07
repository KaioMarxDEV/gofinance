package transaction

import "github.com/gofiber/fiber/v2"

func All(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
func Add(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
