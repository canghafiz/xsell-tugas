package product

import "be/models/domains"

type GeneralResource struct {
	ProductId   int              `json:"product_id"`
	ProductSlug string           `json:"product_slug"`
	Title       string           `json:"title"`
	Price       float64          `json:"price"`
	Condition   string           `json:"condition"`
	Images      []ImageResource  `json:"images"`
	Location    LocationResource `json:"location"`
	Listing     ListingResource  `json:"listing"`
}

func ToGeneralResource(product domains.Products) *GeneralResource {
	res := &GeneralResource{
		ProductId:   product.ProductId,
		ProductSlug: product.ProductSlug,
		Title:       product.Title,
		Price:       product.Price,
		Condition:   string(product.Condition),
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

	return res
}

func ToGeneralResources(products []domains.Products) []GeneralResource {
	var res []GeneralResource
	for _, product := range products {
		res = append(res, *ToGeneralResource(product))
	}
	return res
}
