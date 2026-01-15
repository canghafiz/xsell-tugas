package services

import "be/models/resources/meta"

type MetaSeoServ interface {
	GetByPageKey(pageKey string) ([]meta.Resource, error)
}
