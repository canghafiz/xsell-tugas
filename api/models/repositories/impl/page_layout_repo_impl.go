package impl

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type PageLayoutRepoImpl struct {
}

func NewPageLayoutRepoImpl() *PageLayoutRepoImpl {
	return &PageLayoutRepoImpl{}
}

func (repo *PageLayoutRepoImpl) GetPageLayout(db *gorm.DB, pageKey string) ([]domains.PageLayout, error) {
	var layouts []domains.PageLayout

	if err := db.
		Where("page_key = ? AND is_active = ?", pageKey, true).
		Order("sort_order ASC").
		Preload("Section").
		Find(&layouts).Error; err != nil {
		return nil, err
	}

	return layouts, nil
}
