package product

import (
	"be/models/domains"
	"time"
)

type SingleProductResource struct {
	ProductId   int                 `json:"product_id"`
	Title       string              `json:"title"`
	Slug        string              `json:"slug"`
	Description string              `json:"description"`
	Price       float64             `json:"price"`
	Condition   string              `json:"condition"`
	Status      string              `json:"status"`
	ViewCount   int                 `json:"view_count"`
	Category    SubCategoryResource `json:"sub_category"`
	Images      []ImageResource     `json:"images"`
	Specs       []SpecResource      `json:"specs"`
	Location    LocationResource    `json:"location"`
	Listing     ListingResource     `json:"listing"`
	CreatedAt   string              `json:"created_at"`
	UpdatedAt   string              `json:"updated_at"`
}

type SubCategoryResource struct {
	SubCategoryId   int              `json:"sub_category_id"`
	SubCategoryName string           `json:"sub_category_name"`
	Slug            string           `json:"slug"`
	Icon            string           `json:"icon"`
	Category        CategoryResource `json:"category"`
}

type CategoryResource struct {
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	CategorySlug string `json:"category_slug"`
	Icon         string `json:"icon"`
}

type ImageResource struct {
	ImageId   int    `json:"image_id"`
	URL       string `json:"url"`
	IsPrimary bool   `json:"is_primary"`
	OrderSeq  int    `json:"order_seq"`
}

type SpecResource struct {
	SpecId        int    `json:"spec_id"`
	SpecTypeTitle string `json:"spec_type_title"`
	Name          string `json:"name"`
	Value         string `json:"value"`
	Category      string `json:"category,omitempty"`
}

type LocationResource struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   string  `json:"address"`
}

type ListingResource struct {
	UserId       int     `json:"user_id"`
	Email        string  `json:"email"`
	FirstName    string  `json:"first_name"`
	LastName     *string `json:"last_name"`
	PhotoProfile *string `json:"photo_profile"`
}

func ToSingleProductResource(product *domains.Products) *SingleProductResource {
	if product == nil {
		return nil
	}

	res := &SingleProductResource{
		ProductId:   product.ProductId,
		Title:       product.Title,
		Slug:        product.ProductSlug,
		Description: product.Description,
		Price:       product.Price,
		Condition:   string(product.Condition),
		Status:      string(product.Status),
		ViewCount:   product.ViewCount,
		Location: LocationResource{
			Latitude:  product.Location.Latitude,
			Longitude: product.Location.Longitude,
			Address:   product.Location.Address,
		},
		Listing: ListingResource{
			UserId:       product.Listing.UserId,
			Email:        product.Listing.Email,
			FirstName:    product.Listing.FirstName,
			LastName:     product.Listing.LastName,
			PhotoProfile: product.Listing.PhotoProfileUrl,
		},
		CreatedAt: product.CreatedAt.Format(time.RFC3339),
		UpdatedAt: product.UpdatedAt.Format(time.RFC3339),
	}

	// Category
	res.Category = SubCategoryResource{
		SubCategoryId:   product.SubCategory.SubCategoryId,
		SubCategoryName: product.SubCategory.Name,
		Slug:            product.SubCategory.Slug,
		Icon:            product.SubCategory.Icon,
		Category: CategoryResource{
			CategoryId:   product.SubCategory.Category.CategoryId,
			CategoryName: product.SubCategory.Category.CategoryName,
			CategorySlug: product.SubCategory.Category.CategorySlug,
			Icon:         product.SubCategory.Category.Icon,
		},
	}

	// Images
	for _, img := range product.ProductImages {
		res.Images = append(res.Images, ImageResource{
			ImageId:   img.ImageId,
			URL:       img.ImageUrl,
			OrderSeq:  img.OrderSequence,
			IsPrimary: img.IsPrimary,
		})
	}

	// Specs
	for _, spec := range product.ProductSpecs {
		specRes := SpecResource{
			SpecId:        spec.ProductSpecId,
			SpecTypeTitle: spec.CategoryProductSpec.SpecTypeTitle,
			Name:          spec.CategoryProductSpec.SpecName,
			Value:         spec.SpecValue,
		}
		if spec.CategoryProductSpec.CategoryProductSpecId != 0 {
			specRes.Category = spec.CategoryProductSpec.SpecName
		}
		res.Specs = append(res.Specs, specRes)
	}

	return res
}
