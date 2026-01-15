package impl

import (
	"be/models/repositories"
	"be/models/resources/meta"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type MetaSeoServImpl struct {
	Db       *gorm.DB
	MetaRepo repositories.MetaSeoRepo
}

func NewMetaSeoServImpl(db *gorm.DB, metaRepo repositories.MetaSeoRepo) *MetaSeoServImpl {
	return &MetaSeoServImpl{Db: db, MetaRepo: metaRepo}
}

func (serv *MetaSeoServImpl) GetByPageKey(pageKey string) ([]meta.Resource, error) {
	// Call serv
	result, err := serv.MetaRepo.GetByPageKey(serv.Db, pageKey)
	if err != nil {
		log.Printf("[MetaSeoRepo.GetByPageKey] error: %v", err)
		return nil, fmt.Errorf("failed to get meta by page key, please try again later")
	}

	return meta.ToResources(result), nil
}
