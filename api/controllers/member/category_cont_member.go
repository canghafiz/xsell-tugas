package member

import "github.com/gin-gonic/gin"

type CategoryContMember interface {
	GetWithSub(context *gin.Context)
}
