package handler

import (
	"errors"
	"time"

	"github.com/KaioMarxDEV/gofinance/cmd/database"
	"github.com/KaioMarxDEV/gofinance/cmd/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CheckPassword(pass, hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))

	return err == nil
}

func GetUserByUserName(username string) (*model.User, error) {
	db := database.DB
	var user model.User

	if err := db.Table("users").Where("Username=?", username).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	type User struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var (
		input            = new(LoginInput)
		user             User
		err              error
		userFromUserName *model.User
	)

	// gets the request body and bind to input variable filling it with data
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Success: false,
			Message: "Error in getting your input",
			Data:    nil,
		})
	}

	// verifies if inputs passed on request body is empty returning an error in cause true
	if input.Username == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: "Error: input cannot be empty",
			Data:    nil,
		})
	}

	// call the database to validate the login username exists
	userFromUserName, err = GetUserByUserName(input.Username)

	// database query error handler
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Success: false,
			Message: "Internal proccess error, sorry it's our fault",
			Data:    nil,
		})
	}

	// verifies if record with username from request body exists and if don't return error
	if userFromUserName == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Success: false,
			Message: "User does not exist, you need to register first to login",
			Data:    nil,
		})
	}

	// if record exists we fill user var with database data from user found by username
	if userFromUserName != nil {
		user = User{
			ID:       userFromUserName.ID.String(),
			Username: userFromUserName.Username,
			Email:    userFromUserName.Email,
			Password: userFromUserName.Password,
		}
	}

	if !CheckPassword(input.Password, user.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: "password incorrect",
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
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Success: false,
			Message: "JWT token creation failed",
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(ResponseHTTP{
		Success: true,
		Message: "JWT token created",
		Data:    tokenString,
	})
}
