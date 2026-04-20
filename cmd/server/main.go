package main

import (

	"monitoramento/internal/handlers"
	"monitoramento/internal/scheduler"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {

	scheduler.Start()
	engine := html.New("./views",".html")

	app:= fiber.New(fiber.Config{Views: engine})
	app.Use("/static", static.New("./views/public"))

	app.Get("/api", handlers.HandlerPing)
	app.Get("/", handlers.HandlerPage)

	app.Listen(":3000")
}
