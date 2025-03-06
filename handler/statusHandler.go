package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ServiceStatus holds the API response structure
type ServiceStatus struct {
	CountriesNowAPI  string `json:"countriesnowapi"`
	RestCountriesAPI string `json:"restcountriesapi"`
	Version          string `json:"version"`
	Uptime           int64  `json:"uptime"`
}

// Start time for calculating uptime
var startTime = time.Now()

// StatusHandler handles the API diagnostics
func StatusHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("✅ StatusHandler was called!") // test, fjern etterpå

	// Check the status of external services
	countriesNowStatus := checkServiceStatus("https://countriesnow.space/api/v0.1/countries")
	restCountriesStatus := checkServiceStatus("https://restcountries.com/v3.1/all")

	// Calculate uptime
	uptime := int64(time.Since(startTime).Seconds())

	// Construct response
	response := ServiceStatus{
		CountriesNowAPI:  countriesNowStatus,
		RestCountriesAPI: restCountriesStatus,
		Version:          "v1",
		Uptime:           uptime,
	}

	// Set content type and return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// checkServiceStatus pings an API and returns the HTTP status as a string
func checkServiceStatus(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "Service unavailable"
	}
	defer resp.Body.Close()

	return http.StatusText(resp.StatusCode)
}
