package handler

import (
	"github.com/KaioMarxDEV/gofinance/src/database"
	"github.com/KaioMarxDEV/gofinance/src/model"
	"github.com/gofiber/fiber/v2"
)

// Creates a new user on Database
func CreateUser(c *fiber.Ctx) error {
	type NewUser struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	db := database.DB
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "review your input",
			"data":    err,
		})
	}

	// password hashing and validation

	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "couldn't create user",
			"data":    err,
		})
	}

	createdUser := NewUser{
		Username: user.Email,
		Email:    user.Email,
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "Created user",
		"data":    createdUser,
	})
}
