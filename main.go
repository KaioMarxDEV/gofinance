package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type User struct {
	Name string
	Age  int
}

func main() {
	// instace of database in memory
	database := make(map[int]*User)
	// create user for sake of GET method testing
	database[1] = &User{Name: "kaio", Age: 20}

	// server instace creation by gofiber docs
	server := fiber.New()
	// server cors enabling process to debug from postman.com
	server.Use(cors.New())
	// http method GET to request information from the API
	server.Get("/:name?", func(c *fiber.Ctx) error {
		// process to guarantee params is not nil
		if name := c.Params("name"); name != "" {
			// database.find() type of process below:
			for _, v := range database {
				if v.Name == name {
					return c.JSON(v)
				} else {
					return fiber.NewError(404, "user requested not found")
				}
			}
		}
		// if everything failed there's something wrong
		return fiber.NewError(500, "something went wrong, it's our fault")
	})

	// if the instance of the server not run well, send a log on terminal
	log.Fatal(server.Listen(":3000"))
}
