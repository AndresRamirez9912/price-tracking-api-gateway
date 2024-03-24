package main

import (
	"log"
	"net/http"
	"os"
	"price-tracking-api-gateway/src/constants"
	"price-tracking-api-gateway/src/handlers"
	"price-tracking-api-gateway/src/middlewares"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Main Router
	r := chi.NewRouter()
	auth := r.Group(nil)
	auth.Use(middlewares.AuthMiddleware)

	// Handler
	r.Post("/api/AddUser", handlers.ForwardingV1)
	r.Post("/api/logIn", handlers.ForwardingV1)
	r.Post("/api/verifyUser", handlers.ForwardingV1)
	auth.Post("/api/*", handlers.ForwardingV1)

	log.Println("Serving API Gateway on port", os.Getenv(constants.PORT))
	err := http.ListenAndServe(os.Getenv(constants.PORT), r)
	if err != nil {
		log.Fatal("Error serving the application", err)
	}
}
