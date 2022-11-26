package handler

import (
	"github.com/KaioMarxDEV/gofinance/cmd/database"
	"github.com/KaioMarxDEV/gofinance/cmd/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Select all users
func GetAllUsers(c *fiber.Ctx) error {
	type User struct {
		ID        uuid.UUID
		Username  string
		Email     string
		CreatedAt string
	}

	db := database.DB // get instance of database

	var users []User // create array to bind and fill with db data

	err := db.Table("users").Omit("password").Find(&users).Error // binds users var to db return

	// search for errors in the find query
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// return what was found on database
	return c.Status(fiber.StatusOK).JSON(ResponseHTTP{
		Success: true,
		Data:    users,
		Message: "all users",
	})
}

// Select user by id
func GetUserByID(c *fiber.Ctx) error {
	db := database.DB
	var user model.User
	id := c.Params("id")

	// SELECT * FROM users WHERE ID equals to id, and allocate that in user var
	err := db.Table("users").Find(&user, "ID = ?", id).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(ResponseHTTP{
		Success: true,
		Data:    user,
		Message: "userFound",
	})
}

// Hashing password method to CreateUser Handler
func hashPassword(pass string) (string, error) {
	newPass, err := bcrypt.GenerateFromPassword([]byte(pass), 10)

	return string(newPass), err
}

// Creates a new user on Database
func CreateUser(c *fiber.Ctx) error {
	var err error
	db := database.DB
	user := new(model.User)

	// GET INFO FROM BODY AND BIND TO USER VARIABLE
	if err = c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// VALIDATE IF BODY DATA IS EMPTY
	if user.Username == "" || user.Password == "" || user.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: "Sorry cannot insert empty fields",
			Data:    nil,
		})
	}

	var userFound model.User

	db.Table("users").Where("Username = ?", user.Username).Or("Email = ?", user.Email).Find(&userFound)

	// VALIDATE IF EMAIL PASSED IS NEW OR ALREADY EXISTS
	if userFound.Email == user.Email {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: "user email already Exists",
			Data:    nil,
		})
	}

	// VALIDATE IF USERNAME PASSED IS NEW OR ALREADY EXISTS
	if userFound.Username == user.Username {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: "username already Exists",
			Data:    nil,
		})
	}

	// ENCRYPTS THE PASSWORD
	user.Password, err = hashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Success: false,
			Message: "Could not hash password",
			Data:    nil,
		})
	}

	// INSERT QUERY
	if err = db.Create(&user).Error; err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": fiber.ErrBadRequest.Message,
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}
