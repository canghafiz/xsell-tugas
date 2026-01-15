package admin

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type CategoryRepoAdmin interface {
	Create(db *gorm.DB, category domains.Categories) error
	Update(db *gorm.DB, category domains.Categories) error
	Delete(db *gorm.DB, categoryId domains.Categories) error
}
