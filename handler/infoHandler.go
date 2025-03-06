package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CountryInfo represents the structure of a country response
type CountryInfo struct {
	Name       string            `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Capital    string            `json:"capital"`
	Cities     []string          `json:"cities,omitempty"`
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("InfoHandler called with request") // Debugging log

	vars := mux.Vars(r)
	countryCode := strings.TrimSpace(strings.ToLower(vars["code"]))

	if countryCode == "" {
		http.Error(w, "Country code is missing.", http.StatusBadRequest)
		return
	}

	// Fetch general country info from the REST Countries API
	countryAPI := fmt.Sprintf("http://129.241.150.113:8080/v3.1/alpha/%s", countryCode)
	fmt.Println("Calling external API:", countryAPI)
	resp, err := http.Get(countryAPI)
	if err != nil {
		http.Error(w, "Failed to fetch country data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Country not found", http.StatusNotFound)
		return
	}

	// Decode API response
	var countryData []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&countryData); err != nil {
		http.Error(w, "Error processing JSON data", http.StatusInternalServerError)
		return
	}

	// Extract necessary fields
	country := countryData[0]

	// Convert JSON response into the expected structure
	info := CountryInfo{
		Name:       country["name"].(map[string]interface{})["common"].(string),
		Continents: []string{country["continents"].([]interface{})[0].(string)},
		Population: int(country["population"].(float64)),
		Languages:  convertToStringMap(country["languages"].(map[string]interface{})),
		Borders:    convertToStringArray(country["borders"].([]interface{})),
		Flag:       country["flags"].(map[string]interface{})["png"].(string),
		Capital:    country["capital"].([]interface{})[0].(string),
	}

	// Fetch cities from the CountriesNow API
	citiesAPI := fmt.Sprintf("http://129.241.150.113:3500/api/v0.1/countries/cities")
	citiesReqBody := fmt.Sprintf(`{"country":"%s"}`, info.Name)

	req, err := http.NewRequest("POST", citiesAPI, strings.NewReader(citiesReqBody))
	req.Header.Set("Content-Type", "application/json")

	citiesResp, err := http.DefaultClient.Do(req)
	if err == nil {
		defer citiesResp.Body.Close()

		if citiesResp.StatusCode == http.StatusOK {
			var citiesData map[string]interface{}
			if err := json.NewDecoder(citiesResp.Body).Decode(&citiesData); err == nil {
				info.Cities = convertToStringArray(citiesData["data"].([]interface{}))
			}
		} else {
			fmt.Println("Warning: Could not fetch cities:", citiesResp.Status)
		}
	} else {
		fmt.Println("Warning: Failed to fetch cities:", err)
	}

	// --- Implement the limit parameter ---
	if info.Cities != nil {
		// Sort the cities in ascending alphabetical order
		sort.Strings(info.Cities)

		// Check if the limit query parameter is provided
		limitParam := r.URL.Query().Get("limit")
		if limitParam != "" {
			limit, err := strconv.Atoi(limitParam)
			if err != nil {
				http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
				return
			}
			if limit < len(info.Cities) {
				info.Cities = info.Cities[:limit]
			}
		}
	}
	// --- End of limit implementation ---

	// Respond with formatted JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

// Helper function to convert an interface{} array to a string slice
func convertToStringArray(data []interface{}) []string {
	var result []string
	for _, v := range data {
		result = append(result, v.(string))
	}
	return result
}

// Helper function to convert a map with interface{} values to a string map
func convertToStringMap(data map[string]interface{}) map[string]string {
	result := make(map[string]string)
	for k, v := range data {
		result[k] = v.(string)
	}
	return result
}
