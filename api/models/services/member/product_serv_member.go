package member

import (
	"be/models/requests/member/product"
	res "be/models/resources/product"
)

type ProductServMember interface {
	Create(request product.CreateProductRequest) error
	Update(request product.UpdateProductRequest) error
	GetSingleBySlug(productSlug string) (*res.SingleProductResource, error)
	GetRelatedByCategory(categoryIds []int, exceptProductId, limit int, lat, lng float64) ([]res.GeneralResource, error)
	GetByCategory(filter *product.FilterModel) ([]res.GeneralResource, error)
	GetBySectionKey(key string, filter *product.FilterModel) ([]res.GeneralResource, error)
	Search(title string, filter *product.FilterModel) ([]res.GeneralResource, error)
	GetProductsByUserId(userId int, filter product.FilterMyAds) ([]res.MyProductResource, error)
	UpdateStatus(request product.UpdateStatusRequest) error
	UpdateViewCount(productId int) error
	Delete(productId int) error
}
