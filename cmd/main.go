package main

import (
	"fmt"
	"net/http"
	"recurrente_service/internal/handler"
)

func main() {
	http.HandleFunc("/health", handler.HealthCheck)
	fmt.Println("Recurrent service running on :8000")
	http.ListenAndServe(":8000", nil)
}
