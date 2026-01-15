package repositories

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type BannerRepo interface {
	GetBanners(db *gorm.DB) ([]domains.Banners, error)
}
