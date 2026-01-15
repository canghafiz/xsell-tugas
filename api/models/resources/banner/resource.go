package banner

import "be/models/domains"

type Resource struct {
	BannerId    int    `json:"banner_id"`
	Sequence    int    `json:"sequence"`
	ImageUrl    string `json:"image_url"`
	Link        string `json:"link"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func ToResource(banner domains.Banners) Resource {
	return Resource{
		BannerId:    banner.BannerId,
		Sequence:    banner.Sequence,
		ImageUrl:    banner.ImageURL,
		Link:        banner.Link,
		Title:       banner.Title,
		Description: banner.Description,
	}
}

func ToResources(banners []domains.Banners) []Resource {
	var resources []Resource
	for _, banner := range banners {
		resources = append(resources, ToResource(banner))
	}
	return resources
}
