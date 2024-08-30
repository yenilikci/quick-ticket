package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yenilikci/quick-ticket/config"
	"github.com/yenilikci/quick-ticket/db"
	"github.com/yenilikci/quick-ticket/handlers"
	"github.com/yenilikci/quick-ticket/repositories"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "quickticket",
		ServerHeader: "Fiber",
	})

	eventRepository := repositories.NewEventRepository(db)

	server := app.Group("/api")

	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
