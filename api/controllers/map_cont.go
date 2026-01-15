package controllers

import "github.com/gin-gonic/gin"

type MapCont interface {
	GetAddress(context *gin.Context)
	AutoComplete(context *gin.Context)
}
