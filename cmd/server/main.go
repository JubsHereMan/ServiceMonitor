package main

import (
	"monitoramento/internal/handlers"
	"monitoramento/internal/scheduler"

	"github.com/gofiber/fiber/v3"
)

func main() {

	scheduler.Start()
	app:= fiber.New()

	app.Get("/", handlers.HandlerPing)

	app.Listen(":3000")
}
