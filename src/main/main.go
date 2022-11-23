package main

import (
	"log"

	"github.com/KaioMarxDEV/gofinance/src/database"
	"github.com/KaioMarxDEV/gofinance/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectDB() // connect to Docker infra using .env values

	app := fiber.New()  // server instace creation by gofiber
	app.Use(cors.New()) // enabling CORS for cross platform conections

	routes.SetupRoutes(app) // give the route concert to routes pkg passing app instance

	log.Fatal(app.Listen(":3000")) // shutting down the app with fatal function if listening fails
}
