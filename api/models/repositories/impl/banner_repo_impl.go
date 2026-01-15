package impl

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type BannerRepoImpl struct {
}

func NewBannerRepoImpl() *BannerRepoImpl {
	return &BannerRepoImpl{}
}

func (repo *BannerRepoImpl) GetBanners(db *gorm.DB) ([]domains.Banners, error) {
	var results []domains.Banners
	err := db.Model(&domains.Banners{}).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
