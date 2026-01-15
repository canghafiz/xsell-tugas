package product

import "be/models/domains"

type ImageRequest struct {
	ImageURL      string `json:"image_url" validate:"required,url"`
	IsPrimary     bool   `json:"is_primary" validate:"required"`
	OrderSequence int    `json:"order_sequence" validate:"required"`
}

func ImageRequestToDomain(request ImageRequest) domains.ProductImages {
	return domains.ProductImages{
		ImageUrl:      request.ImageURL,
		IsPrimary:     request.IsPrimary,
		OrderSequence: request.OrderSequence,
	}
}

func ImageRequestToDomains(requests []ImageRequest) []domains.ProductImages {
	var productImages []domains.ProductImages
	for _, request := range requests {
		productImages = append(productImages, ImageRequestToDomain(request))
	}
	return productImages
}
