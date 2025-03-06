package main

import (
	"log"
	"net/http"

	"Assignment-1/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Register API routes
	r.HandleFunc("/countryinfo/v1/info/{code}", handler.InfoHandler).Methods("GET")
	r.HandleFunc("/countryinfo/v1/population/{code}", handler.PopulationHandler).Methods("GET")
	r.HandleFunc("/countryinfo/v1/status", handler.StatusHandler).Methods("GET")

	// Start the server
	log.Println("🚀 Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
