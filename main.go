package main

import (
	"log"
	"net/http"

	"Assignment-1/handler" // Ensure this matches your module name
)

func main() {
	// Register API routes
	http.HandleFunc("/", handler.InfoHandler)
	http.HandleFunc("/api/v1/status", handler.StatusHandler)
	http.HandleFunc("/api/v1/population", handler.PopulationHandler)

	// Start server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
