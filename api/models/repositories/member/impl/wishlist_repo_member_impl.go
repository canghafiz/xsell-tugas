package impl

import (
	"be/models/domains"
	"be/models/requests/member/product"
	"errors"

	"gorm.io/gorm"
)

type WishlistRepoMemberImpl struct {
}

func NewWishlistRepoMemberImpl() *WishlistRepoMemberImpl {
	return &WishlistRepoMemberImpl{}
}

func (repo *WishlistRepoMemberImpl) Update(
	db *gorm.DB,
	wishlist domains.Wishlist,
) error {

	if wishlist.UserId == 0 || wishlist.ProductId == 0 {
		return gorm.ErrInvalidValue
	}

	var existing domains.Wishlist

	// 1️⃣ Check if wishlist already exists
	err := db.
		Where("user_id = ? AND product_id = ?", wishlist.UserId, wishlist.ProductId).
		First(&existing).
		Error

	// 2️⃣ If exists → DELETE (toggle off)
	if err == nil {
		return db.
			Where("wishlist_id = ?", existing.WishlistId).
			Delete(&domains.Wishlist{}).
			Error
	}

	// 3️⃣ If not found → INSERT (toggle on)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return db.Create(&wishlist).Error
	}

	// 4️⃣ Other DB error
	return err
}

func (repo *WishlistRepoMemberImpl) CheckProductOnWishlist(
	db *gorm.DB,
	wishlist domains.Wishlist,
) (bool, error) {
	var count int64

	err := db.
		Model(&domains.Wishlist{}).
		Where("user_id = ? AND product_id = ?", wishlist.UserId, wishlist.ProductId).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repo *WishlistRepoMemberImpl) GetWishlistsByUserId(
	db *gorm.DB,
	userId int,
	filter product.FilterMyAds,
) ([]domains.WishlistWithTotalLike, error) {

	var results []domains.WishlistWithTotalLike

	query := db.
		Model(&domains.Wishlist{}).
		Select(`
			wishlists.*,
			COUNT(w2.wishlist_id) AS total_wishlist
		`).
		Joins(`
			LEFT JOIN wishlists w2 
			ON w2.product_id = wishlists.product_id
		`).
		Where("wishlists.user_id = ?", userId).
		Group("wishlists.wishlist_id")

	// sorting
	switch filter.SortBy {
	case "oldest_to_new":
		query = query.Order("wishlists.created_at ASC")

	case "most_liked":
		query = query.Order("total_wishlist DESC, wishlists.created_at DESC")

	default:
		query = query.Order("wishlists.created_at DESC")
	}

	if err := query.
		Preload("Product").
		Preload("Product.ProductImages").
		Find(&results).
		Error; err != nil {
		return nil, err
	}

	return results, nil
}
