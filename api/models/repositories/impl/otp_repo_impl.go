package impl

import (
	"be/models/domains"
	"fmt"

	"gorm.io/gorm"
)

type OtpRepoImpl struct {
}

func NewOtpRepoImpl() *OtpRepoImpl {
	return &OtpRepoImpl{}
}

func (repo *OtpRepoImpl) SendOtp(db *gorm.DB, otp domains.Otp) error {
	var count int64
	if err := db.Model(&domains.Users{}).Where("email = ?", otp.Email).Count(&count).Error; err != nil {
		return fmt.Errorf("database error: %w", err)
	}

	if count == 0 {
		return fmt.Errorf("email not exist")
	}

	if err := db.Where("email = ? AND purpose = ?", otp.Email, otp.Purpose).Delete(&domains.Otp{}).Error; err != nil {
		return fmt.Errorf("failed to delete old OTP: %w", err)
	}

	if err := db.Create(&otp).Error; err != nil {
		return fmt.Errorf("failed to save OTP: %w", err)
	}

	return nil
}

func (repo *OtpRepoImpl) CheckOtp(db *gorm.DB, otp domains.Otp) (bool, error) {
	var storedOtp domains.Otp

	err := db.
		Where("email = ? AND code = ? AND purpose = ?", otp.Email, otp.Code, otp.Purpose).
		Where("expire_at > NOW()").
		First(&storedOtp).Error

	if err != nil {
		return false, err
	}

	if err := db.Delete(&storedOtp).Error; err != nil {
		return false, fmt.Errorf("failed to delete OTP: %w", err)
	}

	return true, nil
}
