package otp

import "be/models/domains"

type CheckRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func CheckRequestToDomains(request CheckRequest) domains.Otp {
	return domains.Otp{
		Email: request.Email,
		Code:  request.Code,
	}
}
