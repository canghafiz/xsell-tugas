package wishlist

import "be/models/domains"

type CheckRequest struct {
	UserId    int `json:"user_id" validate:"required"`
	ProductId int `json:"product_id" validate:"required"`
}

func CheckRequestToDomain(request CheckRequest) domains.Wishlist {
	return domains.Wishlist{
		UserId:    request.UserId,
		ProductId: request.ProductId,
	}
}
