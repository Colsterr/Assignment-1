package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"Assignment-1/entities" // Make sure this matches your go.mod module name
)

// InfoHandler fetches country information based on the country code
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	// Extract country code from URL and ensure it's lowercase
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 5 {
		http.Error(w, "Country code is missing in the request.", http.StatusBadRequest)
		return
	}
	countryCode := strings.TrimSpace(strings.ToLower(pathParts[4])) // Removes spaces & newlines

	// Use the correct API endpoint: /alpha/{code}
	apiURL := fmt.Sprintf("http://129.241.150.113:8080/v3.1/alpha/%s", countryCode)
	fmt.Println("Making request to API:", apiURL) // Debugging output

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Failed to reach API:", err)
		http.Error(w, "Failed to fetch country data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Debug: Print the response status
	fmt.Println("Response status:", resp.StatusCode)

	// Read and print response body for debugging
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("API Response Body:", string(body))

	// If API request fails, return an error
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Country not found", http.StatusNotFound)
		return
	}

	// Decode JSON response
	var countries []entities.CountryInfo
	if err := json.Unmarshal(body, &countries); err != nil || len(countries) == 0 {
		http.Error(w, "Error processing country data", http.StatusInternalServerError)
		return
	}
	country := countries[0] // API returns an array, take the first element

	// Create JSON response
	response := map[string]interface{}{
		"name":       country.Name.Common,
		"continent":  country.Region,
		"population": country.Population,
		"languages":  country.Languages,
		"borders":    country.Borders,
		"flag":       country.Flags.Png,
		"capital":    country.Capital,
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// HomeHandler provides a welcome message
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the Country Info API!\nUse /countryinfo/v1/info/{code} for country details."))
}
