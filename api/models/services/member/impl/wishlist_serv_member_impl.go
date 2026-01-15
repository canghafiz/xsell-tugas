package impl

import (
	"be/models/repositories/member"
	"be/models/requests/member/product"
	"be/models/requests/member/wishlist"
	res "be/models/resources/wishlist"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type WishlistServMemberImpl struct {
	Db           *gorm.DB
	WishlistRepo member.WishlistRepoMember
}

func NewWishlistServMemberImpl(db *gorm.DB, wishlistRepo member.WishlistRepoMember) *WishlistServMemberImpl {
	return &WishlistServMemberImpl{Db: db, WishlistRepo: wishlistRepo}
}

func (serv *WishlistServMemberImpl) Update(request wishlist.UpdateRequest) error {
	model := wishlist.UpdateRequestToDomain(request)

	// Call repo
	err := serv.WishlistRepo.Update(serv.Db, model)
	if err != nil {
		log.Printf("[WishlistRepoMember.Update] error: %v", err)
		return fmt.Errorf("failed to update wishlist, please try again later")
	}

	return nil
}

func (serv *WishlistServMemberImpl) CheckProductOnWishlist(request wishlist.CheckRequest) (bool, error) {
	model := wishlist.CheckRequestToDomain(request)

	// Call repo
	result, err := serv.WishlistRepo.CheckProductOnWishlist(serv.Db, model)
	if err != nil {
		log.Printf("[WishlistRepoMember.CheckProductOnWishlist] error: %v", err)
		return false, fmt.Errorf("failed to check product on wishlist, please try again later")
	}

	return result, nil
}

func (serv *WishlistServMemberImpl) GetWishlistsByUserId(userId int, filter product.FilterMyAds) ([]res.Resource, error) {
	// Call repo
	result, err := serv.WishlistRepo.GetWishlistsByUserId(serv.Db, userId, filter)
	if err != nil {
		log.Printf("[WishlistRepoMember.GetWishlistsByUserId] error: %v", err)
		return nil, fmt.Errorf("failed to get wishlists by user id, please try again later")
	}

	return res.ToResources(result), nil
}
