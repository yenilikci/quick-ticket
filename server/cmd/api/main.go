package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yenilikci/quick-ticket/handlers"
	"github.com/yenilikci/quick-ticket/repositories"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "quickticket",
		ServerHeader: "Fiber",
	})

	eventRepository := repositories.NewEventRepository(nil)

	server := app.Group("/api")

	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(":3000")
}
