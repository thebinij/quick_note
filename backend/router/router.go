package router

import (
	"quick-note/internal/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var app *fiber.App

func InitRouter(userHandler *user.Handler) {
	app = fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	route := app.Group("/api/v1")

	// Public Routes
	route.Get("/ping", HealthCheck)

	route.Post("/signup", userHandler.CreateUser)
	route.Post("/login", userHandler.Login)
	route.Get("/logout", userHandler.Logout)

	// Private Routes

	// Swagger Docs Route
	swaggerRoute := app.Group("/api-docs")
	app.Static("/", "./swaggerui")
	swaggerRoute.Get("/", SwaggerDocs)

}

func SwaggerDocs(c *fiber.Ctx) error {
	return c.SendFile("./swaggerui/index.html")
}

func Start(addr string) error {
	return app.Listen(addr)
}

func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("Server is up and running!")
}
