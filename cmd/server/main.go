package main

import (
	"monitoramento/internal/handlers"
	"github.com/gofiber/fiber/v3"
)

func main() {

	app:= fiber.New()

	app.Get("/", handlers.HandlerPing)

	app.Listen(":3000")
}
