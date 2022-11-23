package main

import (
	"log"

	"github.com/KaioMarxDEV/gofinance/src/database"
	"github.com/KaioMarxDEV/gofinance/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// server instace creation by gofiber - docs
	// enabling CORS for cross platform conections
	app := fiber.New()
	app.Use(cors.New())

	database.ConnectDB() // connect to Docker infra using .env values

	routes.SetupRoutes(app) // give the route concert to routes pkg passing app instance
	// shutting down the app with fatal function if listening fails
	log.Fatal(app.Listen(":3000"))
}
