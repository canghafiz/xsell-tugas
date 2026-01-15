package member

import "be/models/resources/sub_category"

type SubCategoryServMember interface {
	GetByCategoryParent(categorySlug string) ([]sub_category.Resource, error)
}
