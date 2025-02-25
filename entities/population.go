package entities

// PopulationData represents the population response from the API
type PopulationData struct {
	Data struct {
		Country          string `json:"country"`
		PopulationCounts []struct {
			Year  int `json:"year"`
			Value int `json:"value"`
		} `json:"populationCounts"`
	} `json:"data"`
}
