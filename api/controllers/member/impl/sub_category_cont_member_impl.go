package impl

import (
	"be/exceptions"
	"be/helpers"
	"be/models/services/member"

	"github.com/gin-gonic/gin"
)

type SubCategoryContMemberImpl struct {
	SubCategoryServ member.SubCategoryServMember
}

func NewSubCategoryContMemberImpl(subCategoryServ member.SubCategoryServMember) *SubCategoryContMemberImpl {
	return &SubCategoryContMemberImpl{SubCategoryServ: subCategoryServ}
}

func (cont SubCategoryContMemberImpl) GetByCategoryParent(context *gin.Context) {
	categorySlug := context.Param("categorySlug")
	
	// Call service
	result, errServ := cont.SubCategoryServ.GetByCategoryParent(categorySlug)
	if errServ != nil {
		exceptions.ErrorHandler(context, errServ)
		return
	}

	// Send success response
	response := helpers.ApiResponse{
		Success: true,
		Code:    200,
		Data:    result,
	}

	if err := helpers.WriteToResponseBody(context, response.Code, response); err != nil {
		exceptions.ErrorHandler(context, err)
		return
	}
}
