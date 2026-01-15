package product

import "be/models/domains"

type SpecRequest struct {
	CategoryProductSpecID int    `json:"category_product_spec_id" validate:"required,gt=0"`
	SpecValue             string `json:"spec_value" validate:"required,min=1,max=255"`
}

func SpecRequestToDomain(request SpecRequest) domains.ProductSpecs {
	return domains.ProductSpecs{
		CategoryProductSpecId: request.CategoryProductSpecID,
		SpecValue:             request.SpecValue,
	}
}

func SpecRequestToDomains(requests []SpecRequest) []domains.ProductSpecs {
	var productSpecs []domains.ProductSpecs
	for _, request := range requests {
		productSpecs = append(productSpecs, SpecRequestToDomain(request))
	}
	return productSpecs
}
