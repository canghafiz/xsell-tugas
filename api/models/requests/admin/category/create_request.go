package category

import "be/models/domains"

type CreateRequest struct {
	CategoryName string `json:"category_name" validate:"required"`
	Description  string `json:"description" validate:"required"`
	Icon         string `json:"icon" validate:"required"`
}

func CreateRequestToDomain(request CreateRequest) domains.Categories {
	return domains.Categories{
		CategoryName: request.CategoryName,
		Description:  request.Description,
		Icon:         request.Icon,
	}
}
