package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/mogilyoy/chess-api/internal/api"
)

func main() {
	// Инициализируем реализацию интерфейса
	serverImpl := &api.Server{}

	// Создаём роутер chi
	r := chi.NewRouter()

	// Регистрируем сгенерированные хэндлеры
	api.RegisterHandlers(r, serverImpl)

	addr := ":8080"
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
