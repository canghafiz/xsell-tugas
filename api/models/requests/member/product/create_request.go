package product

import "be/models/domains"

type CreateProductRequest struct {
	Title         string                   `json:"title" validate:"required,min=3,max=100"`
	Description   string                   `json:"description" validate:"required"`
	Price         float64                  `json:"price" validate:"required,gt=0,lte=999999999"`
	Condition     domains.ProductCondition `json:"condition" validate:"required,oneof=New 'Like New' Good 'Good Quite' 'Needs Repair'"`
	Status        domains.ProductStatus    `json:"status" validate:"required,oneof=Available 'Sold out'"`
	SubCategoryID int                      `json:"sub_category_id" validate:"required,gt=0"`
	ListingUserId int                      `json:"listing_user_id" validate:"required,gt=0"`
	Images        []ImageRequest           `json:"images,omitempty"`
	Specs         []SpecRequest            `json:"specs,omitempty"`
	Location      LocationRequest          `json:"location"`
}

func CreateProductRequestToDomain(request CreateProductRequest) *domains.Products {
	return &domains.Products{
		Title:         request.Title,
		Description:   request.Description,
		Price:         request.Price,
		Condition:     request.Condition,
		Status:        request.Status,
		SubCategoryId: request.SubCategoryID,
		ListingUserId: request.ListingUserId,
		ProductImages: ImageRequestToDomains(request.Images),
		ProductSpecs:  SpecRequestToDomains(request.Specs),
		Location:      LocationRequestToDomain(request.Location),
	}
}
