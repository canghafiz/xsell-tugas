package impl

import (
	"be/exceptions"
	"be/helpers"
	"be/models/services"

	"github.com/gin-gonic/gin"
)

type MetaSeoContImpl struct {
	MetaSeoServ services.MetaSeoServ
}

func NewMetaSeoContImpl(metaSeoServ services.MetaSeoServ) *MetaSeoContImpl {
	return &MetaSeoContImpl{MetaSeoServ: metaSeoServ}
}

func (cont *MetaSeoContImpl) GetByPageKey(context *gin.Context) {
	pageKey := context.Param("pageKey")

	// Call Service
	result, errServ := cont.MetaSeoServ.GetByPageKey(pageKey)
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
