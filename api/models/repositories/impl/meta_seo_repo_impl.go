package impl

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type MetaSeoRepoImpl struct {
}

func NewMetaSeoRepoImpl() *MetaSeoRepoImpl {
	return &MetaSeoRepoImpl{}
}

func (repo *MetaSeoRepoImpl) GetByPageKey(db *gorm.DB, pageKey string) ([]domains.MetaSeo, error) {
	var metaSeo []domains.MetaSeo
	err := db.Model(&domains.MetaSeo{}).Where("page_key = ?", pageKey).Find(&metaSeo).Error
	if err != nil {
		return nil, err
	}

	return metaSeo, nil
}
