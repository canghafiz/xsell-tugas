package repositories

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type UserVerifyRepo interface {
	UpdateToVerified(db *gorm.DB, verified domains.UserVerified) error
}
