package controllers

import "github.com/gin-gonic/gin"

type CategoryCont interface {
	GetCategories(context *gin.Context)
}
