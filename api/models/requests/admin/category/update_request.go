package category

import "be/models/domains"

type UpdateRequest struct {
	CategoryId   int    `json:"category_id" validate:"required"`
	CategoryName string `json:"category_name" validate:"required"`
	Description  string `json:"description" validate:"required"`
	Icon         string `json:"icon" validate:"required"`
}

func UpdateRequestToDomain(request UpdateRequest) domains.Categories {
	return domains.Categories{
		CategoryId:   request.CategoryId,
		CategoryName: request.CategoryName,
		Description:  request.Description,
		Icon:         request.Icon,
	}
}
