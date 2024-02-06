package main

import (
	"log"
	"net/http"
	"price-tracking-api-gateway/src/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Main Router
	r := chi.NewRouter()

	// Handler
	r.Post("/api/*", handlers.ForwardingV1)

	log.Println("Serving API Gateway on port :3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal("Error serving the application", err)
	}
}
