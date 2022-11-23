package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// server instace creation by gofiber - docs
	// enabling CORS for cross platform conections
	app := fiber.New()
	app.Use(cors.New())

	// TODO: CREATE DATABASE CONNECTIONC HERE

	// TODO: CALL ROUTER INIT FUNCTION HERE PASSING APP TO IT

	// shutting down the app with fatal function if listening fails
	log.Fatal(app.Listen(":3000"))
}
