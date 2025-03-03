package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"Assignment-1/entities"
)

// PopulationHandler fetches historical population data for a country
func PopulationHandler(w http.ResponseWriter, r *http.Request) {
	// Extract country name from query parameter
	countryName := r.URL.Query().Get("country")
	if countryName == "" {
		http.Error(w, "Country name is missing in the request.", http.StatusBadRequest)
		return
	}

	// Prepare request body
	requestBody, _ := json.Marshal(map[string]string{"country": countryName})
	apiURL := "http://129.241.150.113:3500/api/v0.1/countries/population"

	// Send POST request
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		http.Error(w, "Failed to fetch population data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, _ := io.ReadAll(resp.Body)

	// Debugging
	fmt.Println("API Response Body:", string(body))

	// Handle API failure
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
