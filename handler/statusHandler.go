package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Track server start time
var startTime = time.Now()

// StatusHandler provides system diagnostics and lists available API endpoints
func StatusHandler(w http.ResponseWriter, _ *http.Request) { // `_` to avoid unused parameter warning
	// Define external API endpoints
	restCountriesAPI := "http://129.241.150.113:8080/v3.1/all"
	countriesNowAPI := "http://129.241.150.113:3500/api/v0.1/countries"

	// Check REST Countries API status
	restCountriesStatus := checkAPI(restCountriesAPI)

	// Check CountriesNow API status
	countriesNowStatus := checkAPI(countriesNowAPI)

	// Calculate uptime
	uptime := time.Since(startTime).Seconds()

	// List available API endpoints
	apiEndpoints := map[string]string{
		"/api/v1/status":                    "Check API health status",
		"/api/v1/population?country={name}": "Retrieve population data for a given country",
		"/api/v1/info/{code}":               "Retrieve country details by ISO code",
	}

	// Create JSON response
	response := map[string]interface{}{
		"message":          "API is running successfully",
		"countriesnowapi":  countriesNowStatus,
		"restcountriesapi": restCountriesStatus,
		"version":          "v1",
		"uptime":           fmt.Sprintf("%.2f seconds", uptime),
		"endpoints":        apiEndpoints,
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// checkAPI sends a request to an API and returns the HTTP status code
func checkAPI(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error reaching API:", url, err)
		return 0 // Return 0 to indicate the API is unreachable
	}
	defer resp.Body.Close()
	return resp.StatusCode
}
