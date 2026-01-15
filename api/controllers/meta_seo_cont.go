package controllers

import "github.com/gin-gonic/gin"

type MetaSeoCont interface {
	GetByPageKey(context *gin.Context)
}
