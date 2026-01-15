package member

import "github.com/gin-gonic/gin"

type ProductContMember interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	GetSingleBySlug(context *gin.Context)
	GetRelatedByCategory(context *gin.Context)
	GetByCategory(context *gin.Context)
	GetBySectionKey(context *gin.Context)
	Search(context *gin.Context)
	GetProductsByUserId(context *gin.Context)
	UpdateStatus(context *gin.Context)
	UpdateViewCount(context *gin.Context)
	Delete(context *gin.Context)
}
