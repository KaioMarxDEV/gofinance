package handler

import (
	"github.com/KaioMarxDEV/gofinance/database"
	"github.com/KaioMarxDEV/gofinance/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllUsers(c *fiber.Ctx) error {
	// FIXME: implements here
	return nil
}

// Creates a new user on Database
func CreateUser(c *fiber.Ctx) error {
	type NewUser struct {
		ID       uuid.UUID
		username string
		email    string
	}

	db := database.DB
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// TODO:VERIFICATION QUERIES

	// TODO:HASHING PASSWORD

	// INSERT QUERY
	if err := db.Create(&user).Error; err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": fiber.ErrBadRequest.Message,
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}
