package handler

import (
	"github.com/KaioMarxDEV/gofinance/database"
	"github.com/KaioMarxDEV/gofinance/model"
	"github.com/gofiber/fiber/v2"
)

// Select all users
func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB

	users := new([]model.User)

	if err := db.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": users,
	})
}

// Creates a new user on Database
func CreateUser(c *fiber.Ctx) error {
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
