package middlewares

import (
	"errors"
	"time"

	"github.com/KaioMarxDEV/gofinance/database"
	"github.com/KaioMarxDEV/gofinance/models"
	"github.com/KaioMarxDEV/gofinance/types"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func checkPassword(pass, hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))

	return err == nil
}

func getUserByEmail(email string) (*models.User, error) {
	db := database.DB
	var user models.User

	if err := db.Table("users").Where("Email=?", email).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type User struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var (
		input         = new(LoginInput)
		user          User
		err           error
		userFromEmail *models.User
	)

	// gets the request body and bind to input variable filling it with data
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ResponseHTTP{
			Success: false,
			Message: "Error in getting your input",
			Data:    nil,
		})
	}

	// verifies if inputs passed on request body is empty returning an error in cause true
	if input.Email == "" || input.Password == "" {
		return c.Status(fiber.StatusOK).JSON(types.ResponseHTTP{
			Success: false,
			Message: "Error: input cannot be empty",
			Data:    nil,
		})
	}

	// call the database to validate the login username exists
	userFromEmail, err = getUserByEmail(input.Email)

	// database query error handler
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ResponseHTTP{
			Success: false,
			Message: "Internal proccess error, sorry it's our fault",
			Data:    nil,
		})
	}

	// verifies if record with username from request body exists and if don't return error
	if userFromEmail == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ResponseHTTP{
			Success: false,
			Message: "User does not exist, you need to register first to login",
			Data:    nil,
		})
	}

	// if record exists we fill user var with database data from user found by username
	if userFromEmail != nil {
		user = User{
			ID:       userFromEmail.ID.String(),
			Username: userFromEmail.Username,
			Email:    userFromEmail.Email,
			Password: userFromEmail.Password,
		}
	}

	if !checkPassword(input.Password, user.Password) {
		return c.Status(fiber.StatusOK).JSON(types.ResponseHTTP{
			Success: false,
			Message: "password incorrect",
			Data:    nil,
		})
	}

	// Create JWT object and payload configurations
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"user_id":  user.ID,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	// Generate a JWT token string using the JWT object and a secret
	tokenString, err := token.SignedString([]byte("SECRET"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ResponseHTTP{
			Success: false,
			Message: "JWT token creation failed",
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(types.ResponseHTTP{
		Success: true,
		Message: "JWT token created",
		Data:    tokenString,
	})
}
