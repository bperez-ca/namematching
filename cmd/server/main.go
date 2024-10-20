package main

import (
	http_adapter "NameMatching/internal/adapters/http"
	"NameMatching/internal/app"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Initialize services
	riskService := &app.CustomerValidationService{}

	// Initialize adapters
	httpAdapter := http_adapter.NewHTTPAdapter(riskService)

	// Set up routes
	router := mux.NewRouter()
	router.HandleFunc("/name-match", httpAdapter.NameMatchHandler).Methods("POST")
	router.HandleFunc("/email-match", httpAdapter.EmailMatchHandler).Methods("POST")

	// Start the HTTP server
	log.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
