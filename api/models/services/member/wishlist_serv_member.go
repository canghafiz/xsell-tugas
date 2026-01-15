package member

import (
	"be/models/requests/member/product"
	"be/models/requests/member/wishlist"
)
import res "be/models/resources/wishlist"

type WishlistServMember interface {
	Update(request wishlist.UpdateRequest) error
	CheckProductOnWishlist(request wishlist.CheckRequest) (bool, error)
	GetWishlistsByUserId(userId int, filter product.FilterMyAds) ([]res.Resource, error)
}
