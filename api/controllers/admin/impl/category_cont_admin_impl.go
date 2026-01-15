package impl

import (
	"be/exceptions"
	"be/helpers"
	"be/models/requests/admin/category"
	"be/models/services/admin"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryContAdminImpl struct {
	CategoryServ admin.CategoryServAdmin
}

func NewCategoryContAdminImpl(categoryServ admin.CategoryServAdmin) *CategoryContAdminImpl {
	return &CategoryContAdminImpl{CategoryServ: categoryServ}
}

func (cont *CategoryContAdminImpl) Create(context *gin.Context) {
	// Parse Request Body
	request := category.CreateRequest{}
	errParse := helpers.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exceptions.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.CategoryServ.Create(request)
	if errServ != nil {
		exceptions.ErrorHandler(context, errServ)
		return
	}

	// Response
	response := helpers.ApiResponse{
		Success: true,
		Code:    200,
		Data:    nil,
	}

	errResponse := helpers.WriteToResponseBody(context, response.Code, response)
	if errResponse != nil {
		exceptions.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *CategoryContAdminImpl) Update(context *gin.Context) {
	categoryIdStr := context.Param("categoryId")
	categoryId, _ := strconv.Atoi(categoryIdStr)

	// Parse Request Body
	request := category.UpdateRequest{
		CategoryId: categoryId,
	}
	errParse := helpers.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exceptions.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.CategoryServ.Update(request)
	if errServ != nil {
		exceptions.ErrorHandler(context, errServ)
		return
	}

	// Response
	response := helpers.ApiResponse{
		Success: true,
		Code:    200,
		Data:    nil,
	}

	errResponse := helpers.WriteToResponseBody(context, response.Code, response)
	if errResponse != nil {
		exceptions.ErrorHandler(context, errResponse)
		return
	}
}

func (cont *CategoryContAdminImpl) Delete(context *gin.Context) {
	categoryIdStr := context.Param("categoryId")
	categoryId, _ := strconv.Atoi(categoryIdStr)

	// Call Service
	errServ := cont.CategoryServ.Delete(categoryId)
	if errServ != nil {
		exceptions.ErrorHandler(context, errServ)
		return
	}

	// Response
	response := helpers.ApiResponse{
		Success: true,
		Code:    200,
		Data:    nil,
	}

	errResponse := helpers.WriteToResponseBody(context, response.Code, response)
	if errResponse != nil {
		exceptions.ErrorHandler(context, errResponse)
		return
	}
}
