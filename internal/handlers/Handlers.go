package handlers

import (
	"monitoramento/internal/services"
	"github.com/gofiber/fiber/v3"
)
func HandlerPing(c fiber.Ctx) error{
	result:= services.GetLastResults()
	return c.JSON(result)

}


func HandlerPage(c fiber.Ctx) error {
	results := services.GetLastResults()
	return c.Render("index", results)
}