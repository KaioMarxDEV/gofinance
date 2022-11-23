package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// server instace creation by gofiber docs
	server := fiber.New()

	// enabling CORS for cross platform conections
	server.Use(cors.New())

	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Docker")
	})
	// shutting down the server with fatal function
	// if the instance of the server not run well
	// also send a log on terminal
	log.Fatal(server.Listen(":3000"))
}
