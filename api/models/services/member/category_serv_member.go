package member

import "be/models/resources/category"

type CategoryServMember interface {
	GetWithSub() ([]category.Resource, error)
}
