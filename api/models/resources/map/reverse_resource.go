package _map

import (
	"encoding/json"
	"strconv"
)

type ReverseGeocodingResponse struct {
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	DisplayName string `json:"display_name"`
}

type ReverseResource struct {
	Lat     float64 `json:"latitude"`
	Long    float64 `json:"longitude"`
	Address string  `json:"address"`
}

func ToReverseResource(jsonData []byte) (*ReverseResource, error) {
	var geoResponse ReverseGeocodingResponse

	// Unmarshal JSON
	err := json.Unmarshal(jsonData, &geoResponse)
	if err != nil {
		return nil, err
	}

	// Convert string to float64
	lat, errLat := strconv.ParseFloat(geoResponse.Lat, 64)
	if errLat != nil {
		return nil, errLat
	}

	lon, errLong := strconv.ParseFloat(geoResponse.Lon, 64)
	if errLong != nil {
		return nil, errLong
	}

	return &ReverseResource{
		Lat:     lat,
		Long:    lon,
		Address: geoResponse.DisplayName,
	}, nil
}
