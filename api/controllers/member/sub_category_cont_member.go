package member

import "github.com/gin-gonic/gin"

type SubCategoryContMember interface {
	GetByCategoryParent(context *gin.Context)
}
