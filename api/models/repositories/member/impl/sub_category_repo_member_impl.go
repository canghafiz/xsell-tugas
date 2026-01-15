package impl

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type SubCategoryRepoMemberImpl struct {
}

func NewSubCategoryRepoMemberImpl() *SubCategoryRepoMemberImpl {
	return &SubCategoryRepoMemberImpl{}
}

func (repo *SubCategoryRepoMemberImpl) GetByCategoryParent(db *gorm.DB, categorySlug string) ([]domains.SubCategories, error) {
	var subCategories []domains.SubCategories

	err := db.Joins("JOIN categories ON categories.category_id = subcategories.parent_category_id").
		Where("categories.category_slug = ?", categorySlug).
		Find(&subCategories).Error

	if err != nil {
		return nil, err
	}

	return subCategories, nil
}
