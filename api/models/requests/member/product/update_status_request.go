package product

import "be/models/domains"

type UpdateStatusRequest struct {
	ProductId int                   `json:"product_id" validate:"required"`
	Status    domains.ProductStatus `json:"status" validate:"required,oneof=Available 'Sold out'"`
}

func UpdateStatusRequestToDomain(req UpdateStatusRequest) domains.Products {
	return domains.Products{
		ProductId: req.ProductId,
		Status:    req.Status,
	}
}
