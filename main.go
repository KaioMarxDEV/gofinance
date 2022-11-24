package main

import (
	"log"

	"github.com/KaioMarxDEV/gofinance/database"
	"github.com/KaioMarxDEV/gofinance/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// connect to Docker infra using .env values
	var err error
	for err != nil {
		err = database.ConnectDB()
	}

	app := fiber.New()  // server instace creation by gofiber
	app.Use(cors.New()) // enabling CORS for cross platform conections

	app.Use(recover.New()) // recover from panic calls at any point of handlers context

	routes.SetupRoutes(app) // give the route concert to routes pkg passing app instance

	log.Fatal(app.Listen(":3000")) // shutting down the app with fatal function if listening fails
}
