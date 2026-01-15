package otp

import "be/models/domains"

type SendRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func SendRequestToDomain(request SendRequest) domains.Otp {
	return domains.Otp{
		Email:   request.Email,
		Purpose: "email_verification",
	}
}
