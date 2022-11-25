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
	// tries until database docker is ready...this is a work around of Docker
	// and distributed system issue with "waiting to someone"
	connected := false
	for connected != true {
		connected = database.ConnectDB()
	}

	app := fiber.New()  // server instace creation by gofiber
	app.Use(cors.New()) // enabling CORS for cross platform conections (Frontend)

	app.Use(recover.New()) // recover from panic calls at any point of handlers context

	routes.SetupRoutes(app) // initialize routing groups and endpoints

	log.Fatal(app.Listen(":3000")) // shutting down the app with fatal function if listening fails
}
