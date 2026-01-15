package services

import "be/models/resources/banner"

type BannerServ interface {
	GetBanners() ([]banner.Resource, error)
}
