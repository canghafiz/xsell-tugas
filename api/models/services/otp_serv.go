package services

import "be/models/requests/member/otp"

type OtpServ interface {
	SendEmailVerification(request otp.SendRequest) error
	SendPasswordReset(request otp.SendRequest) error
	CheckOtp(request otp.CheckRequest) error
	CheckOtpPassword(request otp.CheckRequest) (*string, error)
}
