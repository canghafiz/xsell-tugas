package dependencies

import (
	"be/controllers"
	contImpl "be/controllers/impl"
	"be/models/domains"
	repo "be/models/repositories/impl"
	"be/models/services/impl"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Dependency struct {
	FileCont     controllers.FileCont
	OtpCont      controllers.OtpCont
	BannerCont   controllers.BannerCont
	CategoryCont controllers.CategoryCont
	MetaSeoCont  controllers.MetaSeoCont
	UserCont     controllers.UserCont
	MapCont      controllers.MapCont
}

func NewDependency(db *gorm.DB, validator *validator.Validate, smtp domains.Smtp, appName string, jwtKey, mapApiEndpoint, mapApiKey string) *Dependency {
	// Repo
	otpRepo := repo.NewOtpRepoImpl()
	bannerRepo := repo.NewBannerRepoImpl()
	categoryRepo := repo.NewCategoryRepoImpl()
	metaRepo := repo.NewMetaSeoRepoImpl()
	userRepo := repo.NewUserRepoImpl()
	userVerifyRepo := repo.NewUserVerifyRepoImpl()

	// Serv
	fileServ := impl.NewFileServImpl()
	smtpServ := impl.NewSmtpServImpl(smtp, appName)
	otpServ := impl.NewOtpServImpl(db, validator, otpRepo, smtpServ, userRepo, userVerifyRepo, jwtKey)
	bannerServ := impl.NewBannerServImpl(db, bannerRepo)
	categoryServ := impl.NewCategoryServImpl(db, categoryRepo)
	metaServ := impl.NewMetaSeoServImpl(db, metaRepo)
	userServ := impl.NewUserServImpl(db, validator, userRepo, jwtKey)
	mapServ := impl.NewMapServImpl(mapApiEndpoint, mapApiKey)

	// Cont
	fileCont := contImpl.NewFileContImpl(fileServ)
	otpCont := contImpl.NewOtpContImpl(otpServ)
	bannerCont := contImpl.NewBannerContImpl(bannerServ)
	categoryCont := contImpl.NewCategoryContImpl(categoryServ)
	metaCont := contImpl.NewMetaSeoContImpl(metaServ)
	userCont := contImpl.NewUserContImpl(userServ)
	mapCont := contImpl.NewMapContImpl(mapServ)

	return &Dependency{
		FileCont:     fileCont,
		OtpCont:      otpCont,
		BannerCont:   bannerCont,
		CategoryCont: categoryCont,
		MetaSeoCont:  metaCont,
		UserCont:     userCont,
		MapCont:      mapCont,
	}
}
