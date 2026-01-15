package member

import "github.com/gin-gonic/gin"

type PageLayoutContMember interface {
	GetHomeLayouts(context *gin.Context)
	GetProductDetailLayouts(context *gin.Context)
}
