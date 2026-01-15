package section

import (
	"be/models/domains"
	"be/models/resources/product"
)

type Resource struct {
	SectionId  int                       `json:"section_id"`
	SectionKey string                    `json:"section_key"`
	Title      string                    `json:"title"`
	SubTitle   string                    `json:"subtitle"`
	Url        string                    `json:"url"`
	Products   []product.GeneralResource `json:"products"`
}

func ToResource(section domains.ContentSection, products []domains.Products) *Resource {
	return &Resource{
		SectionId:  section.SectionID,
		SectionKey: section.SectionKey,
		Title:      section.Title,
		SubTitle:   section.Subtitle,
		Url:        section.Url,
		Products:   product.ToGeneralResources(products),
	}
}
