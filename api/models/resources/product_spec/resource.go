package product_spec

import "be/models/domains"

type Resource struct {
	Id        int    `json:"id"`
	MainTitle string `json:"main_title"`
	Name      string `json:"name"`
}

func ToResource(model domains.CategoryProductSpecs) Resource {
	return Resource{
		Id:        model.CategoryProductSpecId,
		MainTitle: model.SpecTypeTitle,
		Name:      model.SpecName,
	}
}

func ToResources(models []domains.CategoryProductSpecs) []Resource {
	var resources []Resource
	for _, model := range models {
		resources = append(resources, ToResource(model))
	}
	return resources
}
