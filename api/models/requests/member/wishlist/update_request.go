package wishlist

import "be/models/domains"

type UpdateRequest struct {
	UserId    int `json:"user_id" validate:"required"`
	ProductId int `json:"product_id" validate:"required"`
}

func UpdateRequestToDomain(request UpdateRequest) domains.Wishlist {
	return domains.Wishlist{
		UserId:    request.UserId,
		ProductId: request.ProductId,
	}
}
