package handlers

import(
	"monitoramento/internal/services"
)
func HandlerPing() bool{
	result:= services.GetStatusResponse("youtube.com/")
	return result
}