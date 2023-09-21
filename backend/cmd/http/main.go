package main

import (
	// "github.com/gofiber/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"log"
	"os"
)

// @title Quick Note
// @version 1.0
// @description A Browser extension to take quick notes
// @contact.name Binij Shrestha
// @license.name AGPL
// @BasePath /
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the value of the PORT variable from the environment
	PORT := os.Getenv("PORT")

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	// app.Get("/swagger/*", swagger.HandlerDefault)

	go func() {
		app.Listen("0.0.0.0:" + PORT)
	}()
}
