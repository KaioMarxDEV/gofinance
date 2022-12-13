package transaction

import (
	"fmt"

	"github.com/KaioMarxDEV/gofinance/database"
	"github.com/KaioMarxDEV/gofinance/models"
	"github.com/KaioMarxDEV/gofinance/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func All(c *fiber.Ctx) error {
	var (
		db           = database.DB
		transactions []models.Transaction
		user_id      uuid.UUID
		err          error
	)
	user_id = uuid.MustParse(fmt.Sprint(c.Locals("user_id")))

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

	fmt.Println(transaction)

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
