package handler

import "github.com/gofiber/fiber/v2"

// PING type function to see if its everything alright...
func Hello(c *fiber.Ctx) error {
	return c.JSON(ResponseHTTP{
		Success: false,
		Message: "Hello Stranger! so you found this api huh.",
		Data:    nil,
	})
}
