package _map

import (
	"encoding/json"
	"strconv"
)

type AutocompleteGeocodingResponse struct {
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	DisplayName string `json:"display_name"`
}

type AutocompleteResource struct {
	Lat     float64 `json:"latitude"`
	Long    float64 `json:"longitude"`
	Address string  `json:"address"`
}

func ToAutocompleteResources(jsonData []byte) ([]AutocompleteResource, error) {
	var geoResponses []AutocompleteGeocodingResponse

	// Unmarshal JSON array
	err := json.Unmarshal(jsonData, &geoResponses)
	if err != nil {
		return nil, err
	}

	// Convert to slice of AutocompleteResource
	resources := make([]AutocompleteResource, 0, len(geoResponses))

	for _, geoResponse := range geoResponses {
		// Convert string to float64
		lat, err := strconv.ParseFloat(geoResponse.Lat, 64)
		if err != nil {
			continue // Skip invalid entries
		}

		lon, err := strconv.ParseFloat(geoResponse.Lon, 64)
		if err != nil {
			continue // Skip invalid entries
		}

		resources = append(resources, AutocompleteResource{
			Lat:     lat,
			Long:    lon,
			Address: geoResponse.DisplayName,
		})
	}

	return resources, nil
}
