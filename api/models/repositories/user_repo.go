package repositories

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type UserRepo interface {
	Create(db *gorm.DB, user domains.Users) error
	UpdateToken(db *gorm.DB, user domains.Users) error
	CheckTokenValid(db *gorm.DB, user domains.Users) bool
	CheckUserIdValid(db *gorm.DB, userId int) bool
	Update(db *gorm.DB, user domains.Users) (*domains.Users, error)
	ChangePassword(db *gorm.DB, user domains.Users) error
	FindByEmail(db *gorm.DB, email string) (*domains.Users, error)
	CheckPasswordValid(db *gorm.DB, user domains.Users) (bool, error)
	GetByUserId(db *gorm.DB, userId int) (*domains.Users, error)
	ResetToken(db *gorm.DB, email string) error
}
