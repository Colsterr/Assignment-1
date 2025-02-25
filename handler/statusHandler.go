package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Track server start time
var startTime = time.Now()

// StatusHandler provides system diagnostics
func StatusHandler(w http.ResponseWriter, _ *http.Request) { // `_` to avoid unused parameter warning
	// Define API endpoints
	restCountriesAPI := "http://129.241.150.113:8080/v3.1/all"
	countriesNowAPI := "http://129.241.150.113:3500/api/v0.1/countries"

	// Check REST Countries API
	restCountriesStatus := checkAPI(restCountriesAPI)

	// Check CountriesNow API
	countriesNowStatus := checkAPI(countriesNowAPI)

	// Calculate uptime
	uptime := time.Since(startTime).Seconds()

	// Create JSON response
	response := map[string]interface{}{
		"countriesnowapi":  countriesNowStatus,
		"restcountriesapi": restCountriesStatus,
		"version":          "v1",
		"uptime":           fmt.Sprintf("%.2f seconds", uptime),
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// checkAPI sends a request to an API and returns the status code
func checkAPI(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error reaching API:", url, err) // ✅ Debugging log
		return 0                                     // ✅ Return 0 if the API is unreachable
	}
	defer resp.Body.Close()
	return resp.StatusCode
}
