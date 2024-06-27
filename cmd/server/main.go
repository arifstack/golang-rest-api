package main

import (
	"fmt"
	"log"

	"golang-rest-api/config"
	"golang-rest-api/internal/app"
)

func main() {
	// Menggunakan LoadDevConfig untuk memuat konfigurasi development
	cfg := config.LoadDevConfig()

	// Initialize router and server
	router := app.SetupRouter()
	server := app.NewServer(cfg.Server.Port, router)

	// Start the server
	fmt.Printf("Starting server on port %d...\n", cfg.Server.Port)
	if err := server.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
