package services

type SmtpServ interface {
	SendEmailOtp(email string, code string) error
}
