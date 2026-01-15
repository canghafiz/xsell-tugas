package member

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type CategoryRepoMember interface {
	GetWithSub(db *gorm.DB) ([]domains.Categories, error)
}
