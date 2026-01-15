package repositories

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type OtpRepo interface {
	SendOtp(db *gorm.DB, otp domains.Otp) error
	CheckOtp(db *gorm.DB, otp domains.Otp) (bool, error)
}
