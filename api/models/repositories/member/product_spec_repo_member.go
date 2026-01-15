package member

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type ProductSpecRepoMember interface {
	GetProductSpecBySubId(db *gorm.DB, subId int) ([]domains.CategoryProductSpecs, error)
}
