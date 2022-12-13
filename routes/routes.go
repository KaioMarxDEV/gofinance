package routes

import (
	"github.com/KaioMarxDEV/gofinance/middlewares"
	"github.com/KaioMarxDEV/gofinance/routes/transaction"
	"github.com/KaioMarxDEV/gofinance/routes/user"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/", Hello)

	auth := api.Group("/login")
	auth.Post("/", middlewares.Login)

	userGroup := api.Group("/user")
	userGroup.Post("/add", user.Add)
	userGroup.Get("/all", middlewares.Authenticate, user.All)
	userGroup.Delete("/delete/:id", middlewares.Authenticate, user.Delete)

	transactionGroup := api.Group("/transaction")
	transactionGroup.Post("/add", middlewares.Authenticate, transaction.Add)
	transactionGroup.Get("/all", middlewares.Authenticate, transaction.All)
	transactionGroup.Delete("/delete/:id", middlewares.Authenticate, transaction.Delete)
	transactionGroup.Get("/search", transaction.Search)
}

// ping method
func Hello(c *fiber.Ctx) error {
	return c.SendString("API up and running")
}
