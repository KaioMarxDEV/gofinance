package middlewares

import (
	"errors"
	"strings"

	"github.com/KaioMarxDEV/gofinance/database"
	"github.com/KaioMarxDEV/gofinance/models"
	"github.com/KaioMarxDEV/gofinance/types"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/montanaflynn/jsonwebtoken"
	"gorm.io/gorm"
)

// --- REPOSITORY FUNCTIONS START ---
func getUserByID(id interface{}) (*models.User, error) {
	db := database.DB

	var user models.User

	if err := db.Table("users").Where("ID=?", id).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return &user, nil
	}

	return &user, nil
}

// --- REPOSITORY FUNCTIONS END ---

// --- MIDDLEWARES FUNCTIONS START ---
func Authenticate(c *fiber.Ctx) error {
	var (
		bearerToken = c.Get("Authorization")

		jwtString   string
		tokenString string
		err         error
		tokenClaims jwt.MapClaims
		userExists  *models.User
	)

	if bearerToken == "" {
		return c.Status(fiber.StatusOK).JSON(types.ResponseHTTP{
			Success: false,
			Message: "missing auth token",
			Data:    nil,
		})
	}

	_, jwtString, _ = strings.Cut(bearerToken, "Bearer ")
	tokenString = strings.TrimSpace(jwtString)

	if tokenString == "" {
		return c.Status(fiber.StatusBadRequest).JSON(types.ResponseHTTP{
			Success: false,
			Message: "auth token passed is invalid",
			Data:    nil,
		})
	}

	tokenClaims, err = jsonwebtoken.DecodeHMAC(tokenString, []byte("SECRET"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ResponseHTTP{
			Success: false,
			Message: "authentication failed",
			Data:    nil,
		})
	}

	// verifies if ID on payload is valid
	userExists, err = getUserByID(tokenClaims["user_id"])

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(types.ResponseHTTP{
			Success: false,
			Message: "payload data in token is invalid",
			Data:    nil,
		})
	}

	// set header "user_id" to receive in all subsequent routes the user_id value
	c.Locals("user_id", userExists.ID)
	// call the next route
	return c.Next()
}

// --- MIDDLEWARES FUNCTIONS END ---
