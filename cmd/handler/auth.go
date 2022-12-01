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
	input := new(LoginInput)
	var user User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Success: false,
			Message: "Error in getting your input",
			Data:    nil,
		})
	}

	if input.Username == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: "Error: input cannot be empty",
			Data:    nil,
		})
	}

	userFromUserName, err := GetUserByUserName(input.Username)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Success: false,
			Message: "Internal proccess error, sorry it's our fault",
			Data:    nil,
		})
	}

	if userFromUserName == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Success: false,
			Message: "User does not exist, you need to register first to login",
			Data:    nil,
		})
	}

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

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

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
