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

	app.Post("/signup", userHandler.CreateUser)
	app.Post("/login", userHandler.Login)
	app.Get("/logout", userHandler.Logout)
}

func Start(addr string) error {
	return app.Listen(addr)
}
