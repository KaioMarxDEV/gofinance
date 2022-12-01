package middleware

import (
	"errors"
	"strings"

	"github.com/KaioMarxDEV/gofinance/cmd/database"
	"github.com/KaioMarxDEV/gofinance/cmd/model"
	"github.com/gofiber/fiber/v2"
	"github.com/montanaflynn/jsonwebtoken"
	"gorm.io/gorm"
)

func GetUserByID(id interface{}) (*model.User, error) {
	db := database.DB

	var user model.User

	if err := db.Table("users").Where("ID=?", id).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &user, nil
	}

	return &user, nil
}

func Authenticate(c *fiber.Ctx) error {
	bearerToken := c.Get("Authorization")

	if bearerToken == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "missing auth token",
		})
	}

	_, string, _ := strings.Cut(bearerToken, "Bearer ")
	tokenString := strings.TrimSpace(string)

	if tokenString == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "auth token passed is invalid",
		})
	}

	tokenClaims, err := jsonwebtoken.DecodeHMAC(tokenString, []byte("SECRET"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "authentication failed",
			"log":     err,
		})
	}

	// verifies if ID on payload is valid
	_, err = GetUserByID(tokenClaims["user_id"])

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "payload data in token is invalid",
		})
	}

	// set header "user_id" to receive in all subsequent routes the user_id value
	c.Locals("user_id", tokenClaims["user_id"])
	return c.Next()
}
