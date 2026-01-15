package member

import (
	"be/models/domains"
	"be/models/requests/member/product"

	"gorm.io/gorm"
)

type ProductRepoMember interface {
	Create(db *gorm.DB, product domains.Products) error
	Update(db *gorm.DB, product domains.Products) error
	GetSingleBySlug(db *gorm.DB, slug string) (*domains.Products, error)
	GetRelatedByCategory(db *gorm.DB, categoryIds []int, exceptProductId, limit int, lat, lng float64) ([]domains.Products, error)
	GetByCategory(db *gorm.DB, filter product.FilterModel) ([]domains.Products, error)
	GetBySectionKey(db *gorm.DB, key string, model product.FilterModel) ([]domains.Products, error)
	Search(db *gorm.DB, title string, filter *product.FilterModel) ([]domains.Products, error)
	GetProductsByUserId(db *gorm.DB, userId int, filter product.FilterMyAds) ([]domains.ProductWithWishlist, error)
	UpdateStatus(db *gorm.DB, product domains.Products) error
	UpdateViewCount(db *gorm.DB, productId int) error
	Delete(db *gorm.DB, product domains.Products) ([]string, error)
}
