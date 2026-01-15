package impl

import (
	"be/models/repositories"
	"be/models/resources/banner"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type BannerServImpl struct {
	Db         *gorm.DB
	BannerRepo repositories.BannerRepo
}

func NewBannerServImpl(db *gorm.DB, bannerRepo repositories.BannerRepo) *BannerServImpl {
	return &BannerServImpl{Db: db, BannerRepo: bannerRepo}
}

func (serv *BannerServImpl) GetBanners() ([]banner.Resource, error) {
	// Call repo
	result, err := serv.BannerRepo.GetBanners(serv.Db)
	if err != nil {
		log.Printf("[BannerRepo.GetBanners] error: %v", err)
		return nil, fmt.Errorf("failed to get banners, please try again later")
	}

	return banner.ToResources(result), nil
}
