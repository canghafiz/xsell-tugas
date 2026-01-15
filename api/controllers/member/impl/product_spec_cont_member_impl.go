package impl

import (
	"be/exceptions"
	"be/helpers"
	"be/models/services/member"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductSpecContMemberImpl struct {
	ProductSpecServMember member.ProductSpecServMember
}

func NewProductSpecContMemberImpl(productSpecServMember member.ProductSpecServMember) *ProductSpecContMemberImpl {
	return &ProductSpecContMemberImpl{ProductSpecServMember: productSpecServMember}
}

func (cont *ProductSpecContMemberImpl) GetProductSpecBySubId(context *gin.Context) {
	subIdStr := context.Param("subId")
	subId, _ := strconv.Atoi(subIdStr)

	// Call service
	result, errServ := cont.ProductSpecServMember.GetProductSpecBySubId(subId)
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
