package useCases

import (
	"encoding/json"

	"github.com/KaioMarxDEV/gofinance/src/User/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUserController(c *fiber.Ctx) error {
	var newUser models.User

	// parsing the user input data from request.body
	err := json.Unmarshal(c.Body(), &newUser)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Error": "JSON parsing failed sorry it's our fault",
		})
	}

	// FIXME: laking on Schema validation, add better validation daddy
	if newUser.Name == "" || newUser.Email == "" {
		return c.Status(400).JSON(fiber.Map{
			"Error": "You're missing data or insuficient",
		})
	}

	err = createUserUseCase(&newUser)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"Error": "Failed to store on database",
		})
	}

	return c.SendStatus(201)
}
