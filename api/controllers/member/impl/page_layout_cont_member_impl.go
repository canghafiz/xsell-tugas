package impl

import (
	"be/exceptions"
	"be/helpers"
	"be/models/requests/member/page"
	"be/models/services/member"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PageLayoutContMemberImpl struct {
	PageLayoutServ member.PageLayoutServMember
}

func NewPageLayoutContMemberImpl(pageLayoutServ member.PageLayoutServMember) *PageLayoutContMemberImpl {
	return &PageLayoutContMemberImpl{PageLayoutServ: pageLayoutServ}
}

func (cont *PageLayoutContMemberImpl) GetHomeLayouts(context *gin.Context) {
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "100"))
	lat := context.Query("latitude")
	lng := context.Query("longitude")

	req := page.HomeLayoutRequest{
		Latitude:  lat,
		Longitude: lng,
		Limit:     limit,
	}

	// Call serv
	result, err := cont.PageLayoutServ.GetHomeLayouts(req)
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

func (cont *PageLayoutContMemberImpl) GetProductDetailLayouts(context *gin.Context) {
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "100"))
	exceptId, _ := strconv.Atoi(context.Param("exceptId"))
	lat := context.Query("latitude")
	lng := context.Query("longitude")

	req := page.ProductDetailLayoutRequest{
		ExceptId:  exceptId,
		Latitude:  lat,
		Longitude: lng,
		Limit:     limit,
	}

	// Call Serv
	result, err := cont.PageLayoutServ.GetProductDetailLayouts(req)
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
