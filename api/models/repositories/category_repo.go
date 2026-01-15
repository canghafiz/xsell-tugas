package repositories

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type CategoryRepo interface {
	GetCategories(db *gorm.DB) ([]domains.Categories, error)
}
