package main

import (
	"log"

	"github.com/KaioMarxDEV/gofinance/src/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// server instace creation by gofiber docs
	server := fiber.New()

	// enabling CORS for cross platform conections
	// server.Use(cors.New())

	server.Route("/api", routes.MainRouter, "mainRouter")
	// shutting down the server with fatal function
	// if the instance of the server not run well
	// also send a log on terminal
	log.Fatal(server.Listen(":3000"))
}
