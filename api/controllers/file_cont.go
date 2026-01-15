package controllers

import "github.com/gin-gonic/gin"

type FileCont interface {
	UploadFiles(context *gin.Context)
	DeleteFiles(context *gin.Context)
}
