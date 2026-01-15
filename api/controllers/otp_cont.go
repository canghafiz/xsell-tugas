package controllers

import "github.com/gin-gonic/gin"

type OtpCont interface {
	SendEmailVerification(context *gin.Context)
	SendPasswordReset(context *gin.Context)
	CheckOtp(context *gin.Context)
	CheckOtpPassword(context *gin.Context)
}
