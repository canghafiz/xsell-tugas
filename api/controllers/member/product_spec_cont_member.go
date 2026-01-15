package member

import "github.com/gin-gonic/gin"

type ProductSpecContMember interface {
	GetProductSpecBySubId(context *gin.Context)
}
