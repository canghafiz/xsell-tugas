package member

import (
	"be/models/domains"
	"be/models/requests/member/product"

	"gorm.io/gorm"
)

type WishlistRepoMember interface {
	Update(db *gorm.DB, wishlist domains.Wishlist) error
	CheckProductOnWishlist(db *gorm.DB, wishlist domains.Wishlist) (bool, error)
	GetWishlistsByUserId(db *gorm.DB, userId int, filter product.FilterMyAds) ([]domains.WishlistWithTotalLike, error)
}
