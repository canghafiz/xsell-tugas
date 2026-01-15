package impl

import (
	"be/exceptions"
	"be/helpers"
	"be/models/services/member"

	"github.com/gin-gonic/gin"
)

type CategoryContMemberImpl struct {
	CategoryServ member.CategoryServMember
}

func NewCategoryContMemberImpl(categoryServ member.CategoryServMember) *CategoryContMemberImpl {
	return &CategoryContMemberImpl{CategoryServ: categoryServ}
}

func (cont *CategoryContMemberImpl) GetWithSub(context *gin.Context) {
	// Call Serv
	result, err := cont.CategoryServ.GetWithSub()
	if err != nil {
		exceptions.ErrorHandler(context, err)
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
