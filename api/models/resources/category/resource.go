package category

import (
	"be/models/domains"
	"be/models/resources/sub_category"
)

type Resource struct {
	CategoryId    int                      `json:"category_id"`
	CategoryName  string                   `json:"category_name"`
	CategorySlug  string                   `json:"category_slug"`
	Description   string                   `json:"description"`
	Icon          string                   `json:"icon"`
	SubCategories *[]sub_category.Resource `json:"sub_categories,omitempty"`
}

func ToResource(category domains.Categories) Resource {
	result := Resource{
		CategoryId:   category.CategoryId,
		CategoryName: category.CategoryName,
		CategorySlug: category.CategorySlug,
		Description:  category.Description,
		Icon:         category.Icon,
	}

	if category.SubCategories != nil {
		result.SubCategories = sub_category.ToResources2(category.SubCategories)
	}

	return result
}

func ToResources(categories []domains.Categories) []Resource {
	var resources []Resource
	for _, category := range categories {
		resources = append(resources, ToResource(category))
	}
	return resources
}
