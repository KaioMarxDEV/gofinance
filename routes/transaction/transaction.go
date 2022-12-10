package transaction

import (
	"github.com/KaioMarxDEV/gofinance/database"
	"github.com/KaioMarxDEV/gofinance/models"
	"github.com/KaioMarxDEV/gofinance/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func All(c *fiber.Ctx) error {
	type Transaction struct {
		ID          uuid.UUID
		Description string `json:"description"`
		Number      int    `json:"number"`
		Category    string `json:"category"`
		Type        string `json:"type"`
	}
	var (
		db           = database.DB
		transactions []models.Transaction
		err          error
	)

	err = db.Table("transactions").Find(&transactions).Error

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
	type TransactionInput struct {
		Description string `json:"description"`
		Number      int    `json:"number"`
		Category    string `json:"category"`
		Type        string `json:"type"`
	}
	var (
		db          = database.DB
		transaction TransactionInput
		err         error
	)

	if err = c.BodyParser(&transaction); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ResponseHTTP{
			Success: false,
			Message: "Failed to parse request input",
			Data:    nil,
		})
	}

	if err = db.Table("transactions").Create(transaction).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(types.ResponseHTTP{
			Success: false,
			Message: "Failed to create transaction",
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(types.ResponseHTTP{
		Success: true,
		Message: "Transaction created!",
		Data:    transaction,
	})
}
