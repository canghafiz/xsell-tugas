package wishlist

import (
	"be/models/domains"
	"time"
)

type Resource struct {
	ProductId int                   `json:"product_id"`
	Slug      string                `json:"slug"`
	Title     string                `json:"title"`
	ViewCount int                   `json:"view_count"`
	MainImage string                `json:"main_image"`
	TotalLike int                   `json:"total_like"`
	Price     float64               `json:"price"`
	Status    domains.ProductStatus `json:"status"`
	CreatedAt time.Time             `json:"created_at"`
}

func ToResource(wishlist domains.WishlistWithTotalLike) Resource {

	var mainImage string
	images := wishlist.Wishlist.Product.ProductImages

	if len(images) > 0 {
		mainImage = images[0].ImageUrl
	}

	product := wishlist.Wishlist.Product

	return Resource{
		ProductId: product.ProductId,
		Slug:      product.ProductSlug,
		Title:     product.Title,
		ViewCount: product.ViewCount,
		MainImage: mainImage,
		TotalLike: wishlist.TotalWishlist,
		Price:     product.Price,
		Status:    product.Status,
		CreatedAt: wishlist.Wishlist.CreatedAt,
	}
}

func ToResources(wishlists []domains.WishlistWithTotalLike) []Resource {
	resources := make([]Resource, 0, len(wishlists))

	for _, wishlist := range wishlists {
		resources = append(resources, ToResource(wishlist))
	}

	return resources
}
