package member

import (
	"be/models/resources/product_spec"
)

type ProductSpecServMember interface {
	GetProductSpecBySubId(subId int) ([]product_spec.Resource, error)
}
