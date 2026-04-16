package handlers

import (
	"monitoramento/internal/services"

	"github.com/gofiber/fiber/v3"
)
func HandlerPing(c fiber.Ctx) error{
	result:= services.GetStatusResponse("youtube.com/")
	
	return c.JSON(fiber.Map{
		"service": "youtube.com",
		"status": result,
	})
}