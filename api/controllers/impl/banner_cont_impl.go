package impl

import (
	"be/exceptions"
	"be/helpers"
	"be/models/services"

	"github.com/gin-gonic/gin"
)

type BannerContImpl struct {
	BannerServ services.BannerServ
}

func NewBannerContImpl(bannerServ services.BannerServ) *BannerContImpl {
	return &BannerContImpl{BannerServ: bannerServ}
}

func (cont *BannerContImpl) GetBanners(context *gin.Context) {
	// Call Service
	result, errServ := cont.BannerServ.GetBanners()
	if errServ != nil {
		exceptions.ErrorHandler(context, errServ)
		return
	}

	// Response
	response := helpers.ApiResponse{
		Success: true,
		Code:    200,
		Data:    result,
	}

	errResponse := helpers.WriteToResponseBody(context, response.Code, response)
	if errResponse != nil {
		exceptions.ErrorHandler(context, errResponse)
		return
	}
}
