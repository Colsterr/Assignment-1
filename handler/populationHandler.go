package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"Assignment-1/utils"
	"github.com/gorilla/mux"
)

// PopulationHandler handles requests for population data
func PopulationHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("📌 Handling population request...")

	// Extract country code
	vars := mux.Vars(r)
	countryCode := strings.ToUpper(strings.TrimSpace(vars["code"]))

	// Get country name dynamically
	countryName, err := utils.GetCountryName(countryCode)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid country code: %s", countryCode), http.StatusBadRequest)
		return
	}

	// Prepare request to external API
	requestBody, _ := json.Marshal(map[string]string{"country": countryName})
	client := http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("POST", "http://129.241.150.113:3500/api/v0.1/countries/population", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println("❌ Failed to create request:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// Execute API call
	resp, err := client.Do(req)
	if err != nil {
		log.Println("❌ API request timeout:", err)
		http.Error(w, "External API request timeout", http.StatusGatewayTimeout)
		return
	}
	defer resp.Body.Close()

	// Read API response
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Population data not found", http.StatusNotFound)
		return
	}

	// Parse API response
	var populationResponse struct {
		Error bool   `json:"error"`
		Msg   string `json:"msg"`
		Data  struct {
			Country          string `json:"country"`
			PopulationCounts []struct {
				Year  int `json:"year"`
				Value int `json:"value"`
			} `json:"populationCounts"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &populationResponse); err != nil {
		log.Println("❌ Failed to parse API response:", err)
		http.Error(w, "Invalid API response format", http.StatusInternalServerError)
		return
	}

	// Extract & filter population data
	limit := r.URL.Query().Get("limit")
	var startYear, endYear int
	applyLimit := false

	if limit != "" {
		years := strings.Split(limit, "-")
		if len(years) == 2 {
			startYear, _ = strconv.Atoi(years[0])
			endYear, _ = strconv.Atoi(years[1])
			applyLimit = true
		} else {
			http.Error(w, "Invalid limit format. Expected {startYear-endYear}", http.StatusBadRequest)
			return
		}
	}

	var filteredPopulation []map[string]int
	totalPopulation, count := 0, 0

	for _, entry := range populationResponse.Data.PopulationCounts {
		if !applyLimit || (entry.Year >= startYear && entry.Year <= endYear) {
			filteredPopulation = append(filteredPopulation, map[string]int{"year": entry.Year, "value": entry.Value})
			totalPopulation += entry.Value
			count++
		}
	}

	// Sort by year
	sort.Slice(filteredPopulation, func(i, j int) bool {
		return filteredPopulation[i]["year"] < filteredPopulation[j]["year"]
	})

	// Calculate mean
	meanPopulation := 0
	if count > 0 {
		meanPopulation = totalPopulation / count
	}

	// Return formatted response
	response := map[string]interface{}{
		"country": populationResponse.Data.Country,
		"mean":    meanPopulation,
		"values":  filteredPopulation,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Println("✔ Successfully fetched population data for", countryCode)
}
