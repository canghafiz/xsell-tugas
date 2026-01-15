package repositories

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type PageLayoutRepo interface {
	GetPageLayout(db *gorm.DB, pageKey string) ([]domains.PageLayout, error)
}
