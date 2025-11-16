package main

import (
	"log"
	"net/http"

	"github.com/mogilyoy/chess-api/internal/api"
)

func main() {
	// Инициализируем реализацию интерфейса
	serverImpl := &api.Server{}

	// Создаём роутер chi
	// r := chi.NewRouter()

	// Регистрируем сгенерированные хэндлеры
	apiHandler := api.HandlerWithOptions(serverImpl, api.ChiServerOptions{
		BaseURL: "",
	})

	mux := http.NewServeMux()
	mux.Handle("/openapi.yaml", http.FileServer(http.Dir("./api")))
	mux.Handle("/", apiHandler)

	addr := ":8080"
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
