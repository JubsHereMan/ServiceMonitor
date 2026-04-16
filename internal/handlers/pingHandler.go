package handlers

import (
	"monitoramento/internal/services"
	"github.com/gofiber/fiber/v3"
)
func HandlerPing(c fiber.Ctx) error{
	result:= services.GetLastResults()
	return c.JSON(result)

}