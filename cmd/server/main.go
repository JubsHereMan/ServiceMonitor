package main

import (
	"fmt"
	"monitoramento/internal/handlers"
)

func main() {
	result := handlers.HandlerPing()
	fmt.Println(result)
}
