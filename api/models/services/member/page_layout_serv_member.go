package member

import (
	"be/models/requests/member/page"
	"be/models/resources"
)

type PageLayoutServMember interface {
	GetHomeLayouts(request page.HomeLayoutRequest) (resources.PageResource, error)
	GetProductDetailLayouts(request page.ProductDetailLayoutRequest) (resources.PageResource, error)
}
