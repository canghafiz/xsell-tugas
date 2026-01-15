package member

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type SubCategoryRepoMember interface {
	GetByCategoryParent(db *gorm.DB, categorySlug string) ([]domains.SubCategories, error)
}
