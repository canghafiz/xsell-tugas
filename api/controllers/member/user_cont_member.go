package member

import "github.com/gin-gonic/gin"

type UserContMember interface {
	Register(context *gin.Context)
	Login(context *gin.Context)
	Logout(context *gin.Context)
	UpdateData(context *gin.Context)
}
