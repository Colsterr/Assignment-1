package entities

// CountryInfo represents the structure of a country response
type CountryInfo struct {
	Name       NameStruct        `json:"name"`
	Region     string            `json:"region"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flags      FlagStruct        `json:"flags"`
	Capital    []string          `json:"capital"`
}

// NameStruct is used to get the country's common name
type NameStruct struct {
	Common string `json:"common"`
}

// FlagStruct is used to get the flag image URL
type FlagStruct struct {
	Png string `json:"png"`
}
