package dependencies

import (
	cont "be/controllers/member"
	implCont "be/controllers/member/impl"
	"be/models/domains"
	"be/models/repositories"
	"be/models/repositories/impl"
	memberRepo "be/models/repositories/member"
	implRepo "be/models/repositories/member/impl"
	"be/models/services"
	implGeneralServ "be/models/services/impl"
	"be/models/services/member"
	implServ "be/models/services/member/impl"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type MemberDependency struct {
	Db              *gorm.DB
	RedisServ       services.RedisService
	UserRepo        repositories.UserRepo
	UserServ        member.UserServMember
	UserCont        cont.UserContMember
	ProductRepo     memberRepo.ProductRepoMember
	ProductServ     member.ProductServMember
	ProductCont     cont.ProductContMember
	PageLayoutRepo  repositories.PageLayoutRepo
	PageLayoutServ  member.PageLayoutServMember
	PageLayoutCont  cont.PageLayoutContMember
	SubCategoryCont cont.SubCategoryContMember
	CategoryCont    cont.CategoryContMember
	ProductSpecCont cont.ProductSpecContMember
	WishlistCont    cont.WishlistContMember
}

func NewMemberDependency(db *gorm.DB, validator *validator.Validate, redisConfig domains.RedisConfig, jwtKey string) *MemberDependency {
	// Repositories
	userRepo := impl.NewUserRepoImpl()
	productRepo := implRepo.NewProductRepoMemberImpl()
	pageLayoutRepo := impl.NewPageLayoutRepoImpl()
	subCategoryRepo := implRepo.NewSubCategoryRepoMemberImpl()
	categoryRepo := implRepo.NewCategoryRepoMemberImpl()
	productSpecRepo := implRepo.NewProductSpecRepoMemberImpl()
	wishlistRepo := implRepo.NewWishlistRepoMemberImpl()

	// Services
	redisServ := implGeneralServ.NewRedisServiceImpl(redisConfig)
	fileServ := implGeneralServ.NewFileServImpl()
	userServ := implServ.NewUserServMemberImpl(userRepo, db, validator, jwtKey)
	productServ := implServ.NewProductServMemberImpl(db, validator, productRepo, redisServ, fileServ)
	pageLayoutServ := implServ.NewPageLayoutServMemberImpl(db, redisServ, pageLayoutRepo)
	subCategoryServ := implServ.NewSubCategorySerMemberImpl(db, subCategoryRepo, redisServ)
	categoryServ := implServ.NewCategoryServMemberImpl(db, categoryRepo, redisServ)
	productSpecServ := implServ.NewProductSpecServMemberImpl(db, productSpecRepo, redisServ)
	wishlistServ := implServ.NewWishlistServMemberImpl(db, wishlistRepo)

	// Controllers
	userCont := implCont.NewUserContMemberImpl(userServ)
	productCont := implCont.NewProductContMemberImpl(productServ)
	pageLayoutCont := implCont.NewPageLayoutContMemberImpl(pageLayoutServ)
	subCategoryCont := implCont.NewSubCategoryContMemberImpl(subCategoryServ)
	categoryCont := implCont.NewCategoryContMemberImpl(categoryServ)
	productSpecCont := implCont.NewProductSpecContMemberImpl(productSpecServ)
	wishlistCont := implCont.NewWishlistContMemberImpl(wishlistServ)

	return &MemberDependency{Db: db, RedisServ: redisServ, UserRepo: userRepo, UserCont: userCont, ProductCont: productCont, PageLayoutCont: pageLayoutCont, SubCategoryCont: subCategoryCont, CategoryCont: categoryCont,
		ProductSpecCont: productSpecCont,
		WishlistCont:    wishlistCont,
	}
}
