package dependencies

import (
	"be/models/repositories/admin"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)
import implRepo "be/models/repositories/admin/impl"
import serv "be/models/services/admin"
import implServ "be/models/services/admin/impl"
import cont "be/controllers/admin"
import implCont "be/controllers/admin/impl"

type AdminDependency struct {
	CategoryRepo admin.CategoryRepoAdmin
	CategoryServ serv.CategoryServAdmin
	CategoryCont cont.CategoryContAdmin
}

func NewAdminDependency(Db *gorm.DB, Validator *validator.Validate) *AdminDependency {
	// Repo
	categoryRepo := implRepo.NewCategoryRepoAdminImpl()
	// Serv
	categoryServ := implServ.NewCategoryServAdminImpl(Db, categoryRepo, Validator)
	// Cont
	categoryCont := implCont.NewCategoryContAdminImpl(categoryServ)

	return &AdminDependency{
		CategoryRepo: categoryRepo,
		CategoryServ: categoryServ,
		CategoryCont: categoryCont,
	}
}
