package impl

import (
	"be/models/repositories"
	"be/models/resources/category"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type CategoryServImpl struct {
	Db           *gorm.DB
	CategoryRepo repositories.CategoryRepo
}

func NewCategoryServImpl(db *gorm.DB, categoryRepo repositories.CategoryRepo) *CategoryServImpl {
	return &CategoryServImpl{Db: db, CategoryRepo: categoryRepo}
}

func (serv *CategoryServImpl) GetCategories() ([]category.Resource, error) {
	// Call repo
	result, err := serv.CategoryRepo.GetCategories(serv.Db)
	if err != nil {
		log.Printf("[CategoryRepo.GetCategories] error: %v", err)
		return nil, fmt.Errorf("failed to get categories, please try again later")
	}

	return category.ToResources(result), nil
}
