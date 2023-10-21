package main

import (
	"log"
	"os"
	"quick-note/db"
	"quick-note/internal/user"
	"quick-note/router"

	"github.com/joho/godotenv"
)

// @title Quick Note
// @version 1.0
// @description A Browser extension to take quick notes
// @contact.name Binij Shrestha
// @license.name AGPL
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	dbConnection, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("failed to initialize of database connection: %s", err)
	}

	userRep := user.NewRepository(dbConnection.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	router.InitRouter(userHandler)

	if err := router.Start("0.0.0.0:" + PORT); err != nil {
		log.Fatalf("Error in starting the server: %v", err)
	}
}

// import (
// 	// "github.com/gofiber/swagger"

// 	// app.Get("/swagger/*", swagger.HandlerDefault)

// 	go func() {
// 		app.Listen("0.0.0.0:" + PORT)
// 	}()
