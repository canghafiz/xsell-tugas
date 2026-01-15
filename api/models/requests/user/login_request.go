package user

import (
	"be/models/domains"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func LoginRequestToDomains(request LoginRequest) domains.Users {
	return domains.Users{
		Email:    request.Email,
		Password: request.Password,
	}
}
