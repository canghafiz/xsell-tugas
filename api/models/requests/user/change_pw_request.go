package user

import (
	"be/helpers"
	"be/models/domains"
)

type ChangePwRequest struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func ChangePwToDomain(request ChangePwRequest) domains.Users {
	return domains.Users{
		Password: helpers.HashedPassword(request.Password),
	}
}
