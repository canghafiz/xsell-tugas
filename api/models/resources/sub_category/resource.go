package sub_category

import "be/models/domains"

type Resource struct {
	SubCategoryId    int    `json:"sub_category_id"`
	ParentCategoryId int    `json:"parent_category_id"`
	SubCategoryName  string `json:"sub_category_name"`
	SubCategorySlug  string `json:"sub_category_slug"`
}

func ToResource(subCategory domains.SubCategories) *Resource {
	return &Resource{
		SubCategoryId:    subCategory.SubCategoryId,
		ParentCategoryId: subCategory.ParentCategoryId,
		SubCategoryName:  subCategory.Name,
		SubCategorySlug:  subCategory.Slug,
	}
}

func ToResources(subCategories []domains.SubCategories) []Resource {
	var resources []Resource
	for _, subCategory := range subCategories {
		resources = append(resources, *ToResource(subCategory))
	}
	return resources
}

func ToResources2(subCategories []domains.SubCategories) *[]Resource {
	var resources []Resource
	for _, subCategory := range subCategories {
		resources = append(resources, *ToResource(subCategory))
	}
	return &resources
}
