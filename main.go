package main

import (
	"log"

	"github.com/KaioMarxDEV/gofinance/database"
	"github.com/KaioMarxDEV/gofinance/routes"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Init() {
	connected := false

	for connected != true {
		connected = database.ConnectDB()
	}
}

func main() {
	Init()

	// server instace by gofiber with alternative JSON enconder/decoder
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	app.Use(cors.New())

	app.Static("/", "./web/build")
	app.Static("*", "./web/build/index.html")

	app.Use(recover.New())

	routes.SetUpRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
