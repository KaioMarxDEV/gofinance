package handler

import (
	"errors"
	"time"

	"github.com/KaioMarxDEV/gofinance/cmd/database"
	"github.com/KaioMarxDEV/gofinance/cmd/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CheckPasswordHash(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}

func GetUserByUsername(identity string) (*model.User, error) {
	db := database.DB

	var user model.User
	if err := db.Table("users").Where("Username = ?", identity).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// get input, verifies the authenticity, create a valid jwt.
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}

	type User struct {
		ID       uuid.UUID
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	input := new(LoginInput)
	var userData User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: "Parsing body data on request failed",
			Data:    nil,
		})
	}

	identity := input.Identity
	pass := input.Password
	var userFromUsername *model.User
	var err error

	userFromUsername, err = GetUserByUsername(identity)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Success: false,
			Message: "Failed to login with username passed",
			Data:    nil,
		})
	}

	// if identity inputted is not email, neither username, there is an error, login should not occur
	if userFromUsername == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseHTTP{
			Success: false,
			Message: "Email or Username not valid",
			Data:    nil,
		})
	}

	// if username passed is valid, create a userData struct to serve as struct on next jwt object
	if userFromUsername != nil {
		userData = User{
			ID:       userFromUsername.ID,
			Email:    userFromUsername.Email,
			Username: userFromUsername.Username,
			Password: userFromUsername.Password,
		}
	}

	if !CheckPasswordHash(pass, userData.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: "Password incorrect",
			Data:    userData,
		})
	}

	claims := jwt.MapClaims{
		"username": userData.Username,
		"user_id":  userData.ID,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create jwt object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create jwt token with secret string
	// secret should be env variable and not declared imperatively here
	// but for the sake of the project simplicity it will live here for now.
	tokenString, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Login was finished successfully",
		Data:    tokenString,
	})
}
