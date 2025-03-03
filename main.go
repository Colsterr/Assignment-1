package main

import (
	"fmt"
	"log"
	"net/http"

	"Assignment-1/handler" // Ensure this matches your project name
)

func main() {
	// Register API endpoints
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/countryinfo/v1/info/", handler.InfoHandler)
	http.HandleFunc("/countryinfo/v1/population/", handler.PopulationHandler)
	http.HandleFunc("/countryinfo/v1/status/", handler.StatusHandler) // New status endpoint

	// Start the server
	port := "8080"
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
