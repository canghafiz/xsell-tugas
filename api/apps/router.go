package apps

import (
	"be/dependencies"
	"be/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	AdminDependency  *dependencies.AdminDependency
	MemberDependency *dependencies.MemberDependency
	Dependency       *dependencies.Dependency

	JwtKey string
	Engine *gin.Engine
}

func NewRouter(r Router) *Router {
	authMiddleware := middlewares.AuthMiddleware(r.MemberDependency.Db, r.MemberDependency.UserRepo, r.JwtKey)
	pwMiddleware := middlewares.PwMiddleware(r.MemberDependency.Db, r.MemberDependency.UserRepo, r.JwtKey)

	r.Engine.StaticFS("/assets", http.Dir("./assets"))
	// General
	generalGroup := r.Engine.Group("api/v1/")
	{
		storageGroup := generalGroup.Group("storage")
		{
			storageGroup.POST("/uploadFiles", r.Dependency.FileCont.UploadFiles)
			storageGroup.DELETE("/deleteFile", r.Dependency.FileCont.DeleteFiles)
		}

		bannerGroup := generalGroup.Group("banners")
		{
			bannerGroup.GET("/", r.Dependency.BannerCont.GetBanners)
		}

		categoryGroup := generalGroup.Group("categories")
		{
			categoryGroup.GET("/", r.Dependency.CategoryCont.GetCategories)
		}

		metaSeoGroup := generalGroup.Group("meta")
		{
			metaSeoGroup.GET("/page/:pageKey", r.Dependency.MetaSeoCont.GetByPageKey)
		}

		userGroup := generalGroup.Group("user")
		{
			userGroup.POST("/sendPasswordReset", r.Dependency.OtpCont.SendPasswordReset)
			userGroup.POST("/checkOtpPassword", r.Dependency.OtpCont.CheckOtpPassword)
			userGroup.GET("/:userId", r.Dependency.UserCont.GetByUserId)

			middleware := userGroup.Use(pwMiddleware)
			{
				middleware.PUT("/changePassword", r.Dependency.UserCont.ChangePw)
			}
		}

		mapGroup := generalGroup.Group("map")
		{
			mapGroup.GET("/getAddress", r.Dependency.MapCont.GetAddress)
			mapGroup.GET("/autoComplete", r.Dependency.MapCont.AutoComplete)
		}
	}

	// Admin
	adminGroup := r.Engine.Group("api/v1/admin")
	{
		categoryGroup := adminGroup.Group("/category")
		{
			categoryGroup.POST("/", r.AdminDependency.CategoryCont.Create)
			categoryGroup.PUT("/:categoryId", r.AdminDependency.CategoryCont.Update)
			categoryGroup.DELETE("/:categoryId", r.AdminDependency.CategoryCont.Delete)
		}
	}

	// Member Route
	memberGroup := r.Engine.Group("api/v1/member")
	{
		userGroup := memberGroup.Group("/user")
		{
			userGroup.POST("/register", r.MemberDependency.UserCont.Register)
			userGroup.POST("/login", r.MemberDependency.UserCont.Login)

			authMiddleware := userGroup.Use(authMiddleware)
			{
				authMiddleware.PUT("/:userId", r.MemberDependency.UserCont.UpdateData)
				authMiddleware.DELETE("/logout", r.MemberDependency.UserCont.Logout)
			}
		}

		productGroup := memberGroup.Group("/product")
		{
			productGroup.GET("/:productSlug", r.MemberDependency.ProductCont.GetSingleBySlug)
			productGroup.GET("/relatedByCategories/:exceptId", r.MemberDependency.ProductCont.GetRelatedByCategory)
			productGroup.GET("/category", r.MemberDependency.ProductCont.GetByCategory)
			productGroup.GET("/section/:sectionKey", r.MemberDependency.ProductCont.GetBySectionKey)
			productGroup.GET("/search", r.MemberDependency.ProductCont.Search)
			productGroup.GET("/user/:userId", r.MemberDependency.ProductCont.GetProductsByUserId)
			productGroup.PATCH("/viewCount/:productId", r.MemberDependency.ProductCont.UpdateViewCount)

			authMiddleware := productGroup.Use(authMiddleware)
			{
				authMiddleware.POST("/", r.MemberDependency.ProductCont.Create)
				authMiddleware.PUT("/:productId", r.MemberDependency.ProductCont.Update)
				authMiddleware.PATCH("/:productId/status", r.MemberDependency.ProductCont.UpdateStatus)
				authMiddleware.DELETE("/:productId", r.MemberDependency.ProductCont.Delete)
			}
		}

		pageGroup := memberGroup.Group("/page")
		{
			pageGroup.GET("/home", r.MemberDependency.PageLayoutCont.GetHomeLayouts)
			pageGroup.GET("/detail/:exceptId", r.MemberDependency.PageLayoutCont.GetProductDetailLayouts)
		}

		otpGroup := memberGroup.Group("/otp")
		{
			otpGroup.POST("/sendEmail", r.Dependency.OtpCont.SendEmailVerification)
			otpGroup.POST("/verifyEmail", r.Dependency.OtpCont.CheckOtp)
		}

		categoryGroup := memberGroup.Group("/category")
		{
			categoryGroup.GET("/withSub", r.MemberDependency.CategoryCont.GetWithSub)
			// Sub
			categoryGroup.GET("/:categorySlug/subCategories", r.MemberDependency.SubCategoryCont.GetByCategoryParent)
		}

		productSpecGroup := memberGroup.Group("/productSpec")
		{
			productSpecGroup.GET("/sub/:subId", r.MemberDependency.ProductSpecCont.GetProductSpecBySubId)
		}

		wishlistGroup := memberGroup.Group("/wishlist")
		{
			authMiddleware := wishlistGroup.Use(authMiddleware)
			{
				authMiddleware.POST("/", r.MemberDependency.WishlistCont.Update)
				authMiddleware.GET("/check", r.MemberDependency.WishlistCont.CheckProductOnWishlist)
				authMiddleware.GET("/user/:userId", r.MemberDependency.WishlistCont.GetWishlistsByUserId)
			}
		}
	}

	return &Router{
		AdminDependency:  r.AdminDependency,
		MemberDependency: r.MemberDependency,
		Dependency:       r.Dependency,

		JwtKey: r.JwtKey,
		Engine: r.Engine,
	}
}
