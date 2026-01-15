package admin

import "be/models/requests/admin/category"

type CategoryServAdmin interface {
	Create(request category.CreateRequest) error
	Update(request category.UpdateRequest) error
	Delete(categoryId int) error
}
