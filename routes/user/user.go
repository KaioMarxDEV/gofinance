package user

import (
	"errors"

	"github.com/KaioMarxDEV/gofinance/database"
	"github.com/KaioMarxDEV/gofinance/models"
	"github.com/KaioMarxDEV/gofinance/types"
	"github.com/KaioMarxDEV/gofinance/util/validate"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// helpers functions
func hashPassword(pass string) (string, error) {
	newPass, err := bcrypt.GenerateFromPassword([]byte(pass), 10)

	return string(newPass), err
}

func All(c *fiber.Ctx) error {
	type User struct {
		ID        uuid.UUID
		Username  string
		Email     string
		CreatedAt string
	}
	var (
		db    = database.DB // get instance of database
		users []User        // create array to bind and fill with db data
		err   error
	)

	err = db.Table("users").Omit("password").Find(&users).Error // binds users var to db return

	// search for errors in the find query
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// return what was found on database
	return c.Status(fiber.StatusOK).JSON(types.ResponseHTTP{
		Success: true,
		Data:    users,
		Message: "all users",
	})
}
func Add(c *fiber.Ctx) error {
	var (
		db        = database.DB
		user      = new(models.User)
		userFound models.User
		err       error
	)

	// GET INFO FROM BODY AND BIND TO USER VARIABLE
	if err = c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// VALIDATE IF BODY DATA IS EMPTY
	if user.Username == "" || user.Password == "" || user.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(types.ResponseHTTP{
			Success: false,
			Message: "Sorry cannot insert empty fields",
			Data:    nil,
		})
	}

	// VALIDATE BY STRUCT DEFINITIONS USING UTIL FUNCTION
	errors := validate.Struct(*user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(types.ResponseHTTP{
			Success: false,
			Data:    errors,
			Message: "Given data type is not supported or is wrong",
		})
	}

	// DATABASE QUERY WITH CREATE DATA INPUTTED
	db.Table("users").Where("Username = ?", user.Username).Or("Email = ?", user.Email).Find(&userFound)

	// VALIDATE IF EMAIL PASSED IS NEW OR ALREADY EXISTS
	if userFound.Email == user.Email {
		return c.Status(fiber.StatusBadRequest).JSON(types.ResponseHTTP{
			Success: false,
			Message: "user email already Exists",
			Data:    nil,
		})
	}

	// VALIDATE IF USERNAME PASSED IS NEW OR ALREADY EXISTS
	if userFound.Username == user.Username {
		return c.Status(fiber.StatusBadRequest).JSON(types.ResponseHTTP{
			Success: false,
			Message: "username already Exists",
			Data:    nil,
		})
	}

	// ENCRYPTS THE PASSWORD
	user.Password, err = hashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ResponseHTTP{
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

	return c.Status(fiber.StatusCreated).JSON(types.ResponseHTTP{
		Success: true,
		Data:    user.ID.String(),
		Message: "Created successfully!",
	})
}
func Delete(c *fiber.Ctx) error {
	var (
		db       = database.DB
		id       string
		idParsed uuid.UUID
		err      error
	)

	id = c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusOK).JSON(types.ResponseHTTP{
			Success: false,
			Message: "Id cannot be empty",
			Data:    nil,
		})
	}

	idParsed, err = uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ResponseHTTP{
			Success: false,
			Message: "Failed to parse id string to uuid",
			Data:    nil,
		})
	}

	err = db.Table("users").Delete(&models.User{}, "ID=?", idParsed).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusOK).JSON(types.ResponseHTTP{
			Success: false,
			Message: "trying to delete an non existent line",
			Data:    nil,
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ResponseHTTP{
			Success: false,
			Message: "Failed delete query:",
			Data:    nil,
		})
	}

	return c.Status(200).JSON(types.ResponseHTTP{
		Success: true,
		Message: "Deleted successfully!",
		Data:    id,
	})
}
