package main

import (
	"github.com/gofiber/fiber/v2"
	"sgcu65-backend-assignment/database"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Connect to the database
	database.ConnectDB()

	// Listen on PORT 3000
	app.Listen(":3000")
}
