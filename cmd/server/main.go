package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/nathanfabio/schedule-saas/config"
)

func main() {
	// Connect to the database
	config.ConnectDB()

	// Create router
	r := chi.NewRouter()

	// Server port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	fmt.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}