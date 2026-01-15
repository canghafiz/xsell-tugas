package controllers

import "github.com/gin-gonic/gin"

type BannerCont interface {
	GetBanners(context *gin.Context)
}
