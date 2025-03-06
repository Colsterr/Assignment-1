package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// GetCountryName fetches the full country name based on its ISO 3166-2 code.
func GetCountryName(code string) (string, error) {
	apiURL := fmt.Sprintf("http://129.241.150.113:8080/v3.1/alpha/%s", code)
	client := http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(apiURL)
	if err != nil {
		log.Println("❌ Error fetching country name:", err)
		return "", err
	}
	defer resp.Body.Close()

	var countryData []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&countryData); err != nil {
		log.Println("❌ Error parsing country data:", err)
		return "", err
	}

	name, ok := countryData[0]["name"].(map[string]interface{})["common"].(string)
	if !ok {
		return "", fmt.Errorf("Invalid country data format for %s", code)
	}

	return name, nil
}
