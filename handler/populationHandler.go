package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"Assignment-1/entities" // Ensure this matches your module name in go.mod
)

// PopulationHandler fetches historical population data for a country
func PopulationHandler(w http.ResponseWriter, r *http.Request) {
	// Extract country name from URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 5 {
		http.Error(w, "Country name is missing in the request.", http.StatusBadRequest)
		return
	}
	countryName := strings.TrimSpace(pathParts[4]) // Keep full country name (not ISO code)

	// Create JSON request payload
	requestBody, _ := json.Marshal(map[string]string{"country": countryName})

	// Use the correct API endpoint
	apiURL := "http://129.241.150.113:3500/api/v0.1/countries/population"

	// Send POST request to external API
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		http.Error(w, "Failed to fetch population data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, _ := io.ReadAll(resp.Body)

	// If API request fails, return an error
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Population data not found", http.StatusNotFound)
		return
	}

	// Decode JSON response
	var populationData entities.PopulationData
	if err := json.Unmarshal(body, &populationData); err != nil {
		http.Error(w, "Error processing population data", http.StatusInternalServerError)
		return
	}

	// Calculate mean population
	totalPopulation := 0
	count := len(populationData.Data.PopulationCounts)
	for _, entry := range populationData.Data.PopulationCounts {
		totalPopulation += entry.Value
	}
	meanPopulation := 0
	if count > 0 {
		meanPopulation = totalPopulation / count
	}

	// Create JSON response
	response := map[string]interface{}{
		"country": populationData.Data.Country,
		"mean":    meanPopulation,
		"values":  populationData.Data.PopulationCounts,
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
