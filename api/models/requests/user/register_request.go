package user

import (
	"be/helpers"
	"be/models/domains"
)

type RegisterRequest struct {
	FirstName string  `json:"first_name" validate:"required,max=25"`
	LastName  *string `json:"last_name" validate:"omitempty,max=25"`
	Email     string  `json:"email" validate:"required,email,max=100"`
	Password  string  `json:"password" validate:"required,min=6,max=100"`
}

func RegisterRequestToDomain(request RegisterRequest) *domains.Users {
	return &domains.Users{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  helpers.HashedPassword(request.Password),
		Verified: domains.UserVerified{
			Email: request.Email,
		},
	}
}
