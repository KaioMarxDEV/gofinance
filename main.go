package main

import (
	"log"

	"github.com/KaioMarxDEV/gofinance/cmd/database"
	"github.com/KaioMarxDEV/gofinance/cmd/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func InitDatabase() {
	// connect to Docker infra using .env values
	// tries until database docker is ready...this is a work around of Docker
	// and distributed system issue with "waiting to someone"
	connected, try := false, 0
	for connected != true {
		if try > 10 {
			log.Fatal("Something is wrong, impossible to connect on Database")
		}
		connected = database.ConnectDB()
		try++
	}
}

func main() {
	InitDatabase()      // initiate database connection
	app := fiber.New()  // server instace creation by gofiber
	app.Use(cors.New()) // enabling CORS for cross platform conections (Frontend)

	app.Use(recover.New()) // recover from panic calls at any point of handlers context

	routes.SetupRoutes(app) // initialize routing groups and endpoints

	log.Fatal(app.Listen(":3000")) // shutting down the app with fatal function if listening fails
}
