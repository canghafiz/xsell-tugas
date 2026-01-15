package impl

import (
	"be/exceptions"
	"be/helpers"
	"be/models/services"

	"github.com/gin-gonic/gin"
)

type CategoryContImpl struct {
	CategoryServ services.CategoryServ
}

func NewCategoryContImpl(categoryServ services.CategoryServ) *CategoryContImpl {
	return &CategoryContImpl{CategoryServ: categoryServ}
}

func (cont *CategoryContImpl) GetCategories(context *gin.Context) {
	// Call Service
	result, errServ := cont.CategoryServ.GetCategories()
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
