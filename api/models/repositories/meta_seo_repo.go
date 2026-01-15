package repositories

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type MetaSeoRepo interface {
	GetByPageKey(db *gorm.DB, pageKey string) ([]domains.MetaSeo, error)
}
