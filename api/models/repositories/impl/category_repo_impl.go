package impl

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type CategoryRepoImpl struct {
}

func NewCategoryRepoImpl() *CategoryRepoImpl {
	return &CategoryRepoImpl{}
}

func (repo *CategoryRepoImpl) GetCategories(db *gorm.DB) ([]domains.Categories, error) {
	var results []domains.Categories
	err := db.Model(&domains.Categories{}).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
