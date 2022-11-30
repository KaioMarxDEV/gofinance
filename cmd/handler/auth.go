package handler

import (
	"errors"
	"net/mail"

	"github.com/KaioMarxDEV/gofinance/cmd/database"
	"github.com/KaioMarxDEV/gofinance/cmd/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CheckPasswordHash(pass string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(hash))
	return err == nil
}

func GetUserByEmail(identity string) (*model.User, error) {
	db := database.DB

	var user model.User
	if err := db.Table("users").Where("Email = ?", identity).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
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

func valid(email string) bool {
	// parser the email to address like and get errors if fails,
	// returning a bool that represent success or not.
	_, err := mail.ParseAddress(email)
	return err == nil
}

// get input, verifies the authenticity, create a valid jwt and call next middleware
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"Identity"`
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
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Success: false,
			Message: "parse input structure failed",
			Data:    nil,
		})
	}

	// verify if identity is
	identity := input.Identity
	pass := input.Password
	userFromUsername, userFromEmail, err := new(model.User), new(model.User), *new(error)

	// input can be a username or email, this if else verifies this situation
	if valid(identity) {
		userFromEmail, err = GetUserByEmail(identity)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
				Success: false,
				Message: "Failed to login with email passed",
				Data:    nil,
			})
		}
	} else {
		userFromUsername, err = GetUserByUsername(identity)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
				Success: false,
				Message: "Failed to login with username passed",
				Data:    nil,
			})
		}
	}

	if userFromEmail == nil || userFromUsername == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseHTTP{
			Success: false,
			Message: "Email or Username not valid",
			Data:    nil,
		})
	}

	if userFromEmail != nil {
		userData = User{
			ID:       userFromEmail.ID,
			Email:    userFromEmail.Email,
			Username: userFromEmail.Username,
			Password: userFromEmail.Password,
		}
	}

	if userFromUsername != nil {
		userData = User{
			ID:       userFromUsername.ID,
			Email:    userFromUsername.Email,
			Username: userFromUsername.Username,
			Password: userFromUsername.Password,
		}
	}

	// TODO: START HERE TODAY
	if !CheckPasswordHash(pass, userData.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: "Password incorrect",
			Data:    nil,
		})
	}

	// Create jwt token here
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": userData.Username,
		"email":    userData.Email,
		"id":       userData.ID,
	})

	// secret should be env variable and not declared imperatively here
	// but for the sake of the project simplicity it will live here for now.
	tokenString, err := token.SignedString("secret")

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Success: false,
			Message: "Failed to create authentication token",
			Data:    nil,
		})
	}

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Login was finished successfully",
		Data:    tokenString,
	})
}
