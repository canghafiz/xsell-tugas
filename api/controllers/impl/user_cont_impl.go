package impl

import (
	"be/exceptions"
	"be/helpers"
	"be/models/requests/user"
	"be/models/services"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserContImpl struct {
	UserServ services.UserServ
}

func NewUserContImpl(userServ services.UserServ) *UserContImpl {
	return &UserContImpl{UserServ: userServ}
}

func (cont *UserContImpl) ChangePw(context *gin.Context) {
	header := context.GetHeader("Authorization")
	jwt := strings.TrimPrefix(header, "Bearer ")

	// Parse Request Body
	request := user.ChangePwRequest{}
	errParse := helpers.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exceptions.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.UserServ.ChangePw(request, jwt)
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

func (cont *UserContImpl) GetByUserId(context *gin.Context) {
	userIdStr := context.Param("userId")
	userId, _ := strconv.Atoi(userIdStr)

	// Call Service
	result, errServ := cont.UserServ.GetByUserId(userId)
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
