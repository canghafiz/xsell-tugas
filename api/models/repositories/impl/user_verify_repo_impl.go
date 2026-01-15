package impl

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type UserVerifyRepoImpl struct {
}

func NewUserVerifyRepoImpl() *UserVerifyRepoImpl {
	return &UserVerifyRepoImpl{}
}

func (repo *UserVerifyRepoImpl) UpdateToVerified(db *gorm.DB, verified domains.UserVerified) error {
	err := db.
		Model(&verified).
		Where("email = ?", verified.Email).
		Updates(map[string]interface{}{
			"verified": true,
		}).Error

	if err != nil {
		return err
	}

	return nil
}
