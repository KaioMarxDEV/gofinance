package transaction

import (
	"errors"
	"fmt"

	"github.com/KaioMarxDEV/gofinance/database"
	"github.com/KaioMarxDEV/gofinance/models"
	"github.com/KaioMarxDEV/gofinance/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TEST: test to edge cases and bugs
func All(c *fiber.Ctx) error {
	var (
		db           = database.DB
		transactions []models.Transaction
		err          error
	)
	user_id := fmt.Sprint(c.Locals("user_id"))

	err = db.Table("transactions").Where("user_id = ?", user_id).Find(&transactions).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ResponseHTTP{
			Success: false,
			Message: "Failed to query transactions",
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(types.ResponseHTTP{
		Success: true,
		Data:    transactions,
		Message: "all transactions",
	})
}

// TEST: test to edge cases and bugs
func Add(c *fiber.Ctx) error {
	var (
		db          = database.DB
		transaction *models.Transaction
		err         error
	)

	if err = c.BodyParser(&transaction); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ResponseHTTP{
			Success: false,
			Message: "Failed to parse request input",
			Data:    nil,
		})
	}

	user_id := fmt.Sprint(c.Locals("user_id"))

	transaction.UserID = uuid.MustParse(user_id)

	if err = db.Table("transactions").Create(transaction).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ResponseHTTP{
			Success: false,
			Message: "Failed to create transaction",
			Data:    err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(types.ResponseHTTP{
		Success: true,
		Message: "Transaction created!",
		Data:    transaction,
	})
}

// TEST: test to edge cases and bugs
func Delete(c *fiber.Ctx) error {
	var (
		db            = database.DB
		transactionID = c.Params("id")
		transaction   models.Transaction
		err           error
	)

	err = db.Table("transactions").Where("id = ?", transactionID).Find(&transaction).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusBadRequest).JSON(types.ResponseHTTP{
			Success: false,
			Message: "cannot delete unexistent transaction",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ResponseHTTP{
			Success: false,
			Message: "There was an error in processing deletion",
			Data:    nil,
		})
	}

	err = db.Table("transactions").Where("ID = ?", transaction.ID).Delete(&transaction).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ResponseHTTP{
			Success: false,
			Message: "Failed to delete transaction",
		})
	}

	return c.Status(fiber.StatusOK).JSON(types.ResponseHTTP{
		Success: true,
		Message: "Deleted successfully!",
		Data:    transaction.ID,
	})
}

func Search(c *fiber.Ctx) error {
	var (
		db           = database.DB
		query        = c.Query("q")
		transactions []models.Transaction

		err error
	)

	// TODO: NEED TO ADD AUTH MIDDLEWARE TO ENSURE USER VALID
	user_id := fmt.Sprint(c.Locals("user_id"))

	err = db.Table("transactions").Where("description like ? AND user_id = ?", "%"+query+"%", user_id).Or("category like ?", "%"+query+"%").Find(&transactions).Error

	if err != nil {
		return c.Status(fiber.StatusOK).JSON(types.ResponseHTTP{
			Success: false,
			Message: "Failed to query on database",
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(types.ResponseHTTP{
		Success: true,
		Message: "Searched successfully",
		Data:    transactions,
	})
}
