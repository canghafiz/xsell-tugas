package user

import "be/models/domains"

type UpdateDataRequest struct {
	UserId          int     `json:"userId"`
	FirstName       string  `json:"first_name" validate:"required,max=25"`
	LastName        *string `json:"last_name" validate:"omitempty,max=25"`
	PhotoProfileUrl *string `json:"photo_profile_url" validate:"omitempty"`
}

func UpdateDataRequestToDomain(request UpdateDataRequest) domains.Users {
	return domains.Users{
		UserId:          request.UserId,
		FirstName:       request.FirstName,
		LastName:        request.LastName,
		PhotoProfileUrl: request.PhotoProfileUrl,
	}
}
