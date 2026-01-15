package admin

import "github.com/gin-gonic/gin"

type CategoryContAdmin interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}
