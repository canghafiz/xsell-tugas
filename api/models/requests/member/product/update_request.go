package product

import "be/models/domains"

type UpdateProductRequest struct {
	ProductId     int                       `json:"product_id" validate:"required"`
	Title         *string                   `json:"title,omitempty"`
	Description   *string                   `json:"description,omitempty"`
	Price         *float64                  `json:"price,omitempty"`
	Condition     *domains.ProductCondition `json:"condition" validate:"required,oneof=New 'Like New' Good 'Good Quite' 'Needs Repair'"`
	Status        *domains.ProductStatus    `json:"status" validate:"required,oneof=Available 'Sold out'"`
	SubCategoryID *int                      `json:"sub_category_id,omitempty"`
	Images        *[]ImageRequest           `json:"images,omitempty"`
	Specs         *[]SpecRequest            `json:"specs,omitempty"`
	Location      *LocationRequest          `json:"location,omitempty"`
}

func UpdateProductRequestToDomain(request UpdateProductRequest) *domains.Products {
	domain := &domains.Products{
		ProductId: request.ProductId,
	}

	if request.Title != nil {
		domain.Title = *request.Title
	}
	if request.Description != nil {
		domain.Description = *request.Description
	}
	if request.Price != nil {
		domain.Price = *request.Price
	}
	if request.Condition != nil {
		domain.Condition = *request.Condition
	}
	if request.Status != nil {
		domain.Status = *request.Status
	}
	if request.SubCategoryID != nil {
		domain.SubCategoryId = *request.SubCategoryID
	}
	if request.Images != nil {
		domain.ProductImages = ImageRequestToDomains(*request.Images)
	}
	if request.Specs != nil {
		domain.ProductSpecs = SpecRequestToDomains(*request.Specs)
	}
	if request.Location != nil {
		domain.Location = LocationRequestToDomain(*request.Location)
	}

	return domain
}
