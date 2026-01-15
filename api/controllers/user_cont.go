package controllers

import "github.com/gin-gonic/gin"

type UserCont interface {
	ChangePw(context *gin.Context)
	GetByUserId(context *gin.Context)
}
