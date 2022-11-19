package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/KaioMarxDEV/gofinance/src/database"
	"github.com/KaioMarxDEV/gofinance/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "github.com/lib/pq"
)

func main() {
	// server instace creation by gofiber docs
	server := fiber.New()

	// enabling CORS for cross platform conections
	server.Use(cors.New())

	config, err := database.LoadConfig("../../.")

	if err != nil {
		log.Fatal(err)
	}

	connStr := fmt.Sprintf("user=%v dbname=%v sslmode=disable", config.DB_USER, config.DB_NAME)
	store, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%v", store.Stats())

	server.Route("/api", routes.MainRouter, "mainRouter")
	// shutting down the server with fatal function
	// if the instance of the server not run well
	// also send a log on terminal
	log.Fatal(server.Listen(":3000"))
}
