package impl

import (
	"be/exceptions"
	"be/helpers"
	"be/models/requests/user"
	"be/models/services/member"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserContMemberImpl struct {
	UserService member.UserServMember
}

func NewUserContMemberImpl(userService member.UserServMember) *UserContMemberImpl {
	return &UserContMemberImpl{UserService: userService}
}

func (cont *UserContMemberImpl) Register(context *gin.Context) {
	// Parse Request Body
	request := user.RegisterRequest{}
	errParse := helpers.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exceptions.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.UserService.Register(request)
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

func (cont *UserContMemberImpl) Login(context *gin.Context) {
	// Parse Request Body
	request := user.LoginRequest{}
	errParse := helpers.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exceptions.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	result, errServ := cont.UserService.Login(request)
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

func (cont *UserContMemberImpl) Logout(context *gin.Context) {
	email := context.Query("email")

	// Call Service
	errServ := cont.UserService.Logout(email)
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

func (cont *UserContMemberImpl) UpdateData(context *gin.Context) {
	userIdStr := context.Param("userId")
	userId, _ := strconv.Atoi(userIdStr)

	// Parse Request Body
	request := user.UpdateDataRequest{
		UserId: userId,
	}
	errParse := helpers.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exceptions.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	result, errServ := cont.UserService.UpdateData(request)
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
