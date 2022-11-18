package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// server instace creation by gofiber docs
	server := fiber.New()

	server.Use(cors.New(cors.ConfigDefault))

	server.Use("/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello Gin Runtime",
		})
	})

	// if the instance of the server not run well, send a log on terminal
	log.Fatal(server.Listen(":5000"))
}
