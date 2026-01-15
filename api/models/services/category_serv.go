package services

import "be/models/resources/category"

type CategoryServ interface {
	GetCategories() ([]category.Resource, error)
}
