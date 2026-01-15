package product

import (
	"be/models/domains"
	"time"
)

type MyProductResource struct {
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

func ToMyProductResource(product domains.ProductWithWishlist) MyProductResource {
	return MyProductResource{
		ProductId: product.Product.ProductId,
		Slug:      product.Product.ProductSlug,
		Title:     product.Product.Title,
		ViewCount: product.Product.ViewCount,
		MainImage: product.Product.ProductImages[0].ImageUrl,
		TotalLike: product.TotalWishlist,
		Price:     product.Product.Price,
		Status:    product.Product.Status,
		CreatedAt: product.Product.CreatedAt,
	}
}

func ToMyProductResources(products []domains.ProductWithWishlist) []MyProductResource {
	var resources []MyProductResource
	for _, product := range products {
		resources = append(resources, ToMyProductResource(product))
	}
	return resources
}
