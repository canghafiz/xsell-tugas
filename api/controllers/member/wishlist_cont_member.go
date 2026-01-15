package member

import "github.com/gin-gonic/gin"

type WishlistContMember interface {
	Update(context *gin.Context)
	CheckProductOnWishlist(context *gin.Context)
	GetWishlistsByUserId(context *gin.Context)
}
