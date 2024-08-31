package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yenilikci/quick-ticket/config"
	"github.com/yenilikci/quick-ticket/db"
	"github.com/yenilikci/quick-ticket/handlers"
	"github.com/yenilikci/quick-ticket/middlewares"
	"github.com/yenilikci/quick-ticket/repositories"
	"github.com/yenilikci/quick-ticket/services"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "quickticket",
		ServerHeader: "Fiber",
	})

	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)
	authRepository := repositories.NewAuthRepository(db)

	authService := services.NewAuthService(authRepository)

	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middlewares.AuthProtected(db))

	handlers.NewEventHandler(privateRoutes.Group("/event"), eventRepository)
	handlers.NewTicketHandler(privateRoutes.Group("/ticket"), ticketRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
