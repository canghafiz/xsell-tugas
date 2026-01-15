package impl

import (
	"be/exceptions"
	"be/helpers"
	"be/models/requests/member/product"
	"be/models/requests/member/wishlist"
	"be/models/services/member"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WishlistContMemberImpl struct {
	WishlistServ member.WishlistServMember
}

func NewWishlistContMemberImpl(wishlistServ member.WishlistServMember) *WishlistContMemberImpl {
	return &WishlistContMemberImpl{WishlistServ: wishlistServ}
}

func (cont *WishlistContMemberImpl) Update(context *gin.Context) {
	// Parse Request Body
	request := wishlist.UpdateRequest{}
	errParse := helpers.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exceptions.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.WishlistServ.Update(request)
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

func (cont *WishlistContMemberImpl) CheckProductOnWishlist(context *gin.Context) {
	userIdStr := context.DefaultQuery("userId", "0")
	userId, _ := strconv.Atoi(userIdStr)
	productIdStr := context.DefaultQuery("productId", "0")
	productId, _ := strconv.Atoi(productIdStr)

	// Parse Request Body
	request := wishlist.CheckRequest{
		UserId:    userId,
		ProductId: productId,
	}

	// Call Service
	result, errServ := cont.WishlistServ.CheckProductOnWishlist(request)
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

func (cont *WishlistContMemberImpl) GetWishlistsByUserId(context *gin.Context) {
	userIdStr := context.Param("userId")
	userId, _ := strconv.Atoi(userIdStr)

	sortBy := context.Query("sortBy")
	model := product.FilterMyAds{
		SortBy: sortBy,
	}

	// Call Service
	result, errServ := cont.WishlistServ.GetWishlistsByUserId(userId, model)
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
