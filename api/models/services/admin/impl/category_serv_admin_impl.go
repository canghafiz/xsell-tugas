package impl

import (
	"be/helpers"
	"be/models/domains"
	"be/models/repositories/admin"
	"be/models/requests/admin/category"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CategoryServAdminImpl struct {
	Db           *gorm.DB
	CategoryRepo admin.CategoryRepoAdmin
	Validator    *validator.Validate
}

func NewCategoryServAdminImpl(db *gorm.DB, categoryRepo admin.CategoryRepoAdmin, validator *validator.Validate) *CategoryServAdminImpl {
	return &CategoryServAdminImpl{Db: db, CategoryRepo: categoryRepo, Validator: validator}
}
func (serv *CategoryServAdminImpl) Create(request category.CreateRequest) error {
	errValidator := helpers.ErrValidator(request, serv.Validator)
	if errValidator != nil {
		return errValidator
	}

	model := category.CreateRequestToDomain(request)

	// Call repo
	err := serv.CategoryRepo.Create(serv.Db, model)
	if err != nil {
		log.Printf("[CategoryAdminRepo.Create] error: %v", err)
		return fmt.Errorf("failed to create category, please try again later")
	}

	return nil
}

func (serv *CategoryServAdminImpl) Update(request category.UpdateRequest) error {
	errValidator := helpers.ErrValidator(request, serv.Validator)
	if errValidator != nil {
		return errValidator
	}

	model := category.UpdateRequestToDomain(request)

	// Call repo
	err := serv.CategoryRepo.Update(serv.Db, model)
	if err != nil {
		log.Printf("[CategoryAdminRepo.Update] error: %v", err)
		return fmt.Errorf("failed to update category, please try again later")
	}

	return nil
}

func (serv *CategoryServAdminImpl) Delete(categoryId int) error {
	model := domains.Categories{
		CategoryId: categoryId,
	}

	// Call repo
	err := serv.CategoryRepo.Delete(serv.Db, model)
	if err != nil {
		log.Printf("[CategoryAdminRepo.Delete] error: %v", err)
		return fmt.Errorf("failed to delete category, please try again later")
	}

	return nil
}
