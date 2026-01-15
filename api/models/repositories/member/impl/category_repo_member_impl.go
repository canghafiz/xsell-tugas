package impl

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type CategoryRepoMemberImpl struct {
}

func NewCategoryRepoMemberImpl() *CategoryRepoMemberImpl {
	return &CategoryRepoMemberImpl{}
}

func (repo *CategoryRepoMemberImpl) GetWithSub(db *gorm.DB) ([]domains.Categories, error) {
	var categories []domains.Categories
	err := db.Model(&domains.Categories{}).
		Preload("SubCategories").
		Find(&categories).Error
	if err != nil {
		return nil, err
	}
	
	return categories, nil
}
