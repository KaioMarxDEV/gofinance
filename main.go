package main

import (
	"log"

	"github.com/KaioMarxDEV/gofinance/cmd/database"
	"github.com/KaioMarxDEV/gofinance/cmd/routes"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func InitDatabase() {
	// tries until database docker is ready...this is a work around of Docker
	// and distributed system issue with "waiting to someone"
	connected, try := false, 0
	for !connected {
		if try > 10 {
			log.Fatal("Something is wrong, impossible to connect on Database")
		}
		connected = database.ConnectDB()
		try++
	}
}

func main() {
	InitDatabase() // initiate database connection

	// server instace by gofiber with alternative JSON enconder/decoder
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	app.Use(cors.New())

	// Prepare a static middleware to serve the built React files.
	app.Static("/", "./web/build")

	// Prepare a fallback route to always serve the 'index.html', had there not be any matching routes.
	app.Static("*", "./web/build/index.html")

	// recover from panic calls at any point of handlers context
	app.Use(recover.New())

	// initialize routing groups and endpoints
	routes.SetupRoutes(app)

	// shutting down the app with fatal function if listening fails
	log.Fatal(app.Listen(":3000"))
}
