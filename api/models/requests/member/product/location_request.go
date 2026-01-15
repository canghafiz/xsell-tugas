package product

import "be/models/domains"

type LocationRequest struct {
	Latitude  float64 `json:"latitude" validate:"required,gt=-90,lte=90"`
	Longitude float64 `json:"longitude" validate:"required,gt=-180,lte=180"`
	Address   string  `json:"address" validate:"required"`
}

func LocationRequestToDomain(request LocationRequest) domains.Location {
	return domains.Location{
		Latitude:  request.Latitude,
		Longitude: request.Longitude,
		Address:   request.Address,
	}
}
