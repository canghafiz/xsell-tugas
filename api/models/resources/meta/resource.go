package meta

import "be/models/domains"

type Resource struct {
	MetaId    int    `json:"meta_id"`
	PageKey   string `json:"page_key"`
	MetaName  string `json:"meta_name"`
	MetaValue string `json:"meta_value"`
	IsActive  bool   `json:"is_active"`
}

func ToResource(meta domains.MetaSeo) *Resource {
	return &Resource{
		MetaId:    meta.MetaId,
		PageKey:   meta.PageKey,
		MetaName:  meta.MetaName,
		MetaValue: meta.MetaValue,
		IsActive:  meta.IsActive,
	}
}

func ToResources(metas []domains.MetaSeo) []Resource {
	var resources []Resource
	for _, meta := range metas {
		resources = append(resources, *ToResource(meta))
	}
	return resources
}
