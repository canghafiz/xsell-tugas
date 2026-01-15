package impl

import (
	"be/exceptions"
	"be/helpers"
	"be/models/requests/member/otp"
	"be/models/services"

	"github.com/gin-gonic/gin"
)

type OtpContImpl struct {
	OtpServ services.OtpServ
}

func NewOtpContImpl(otpServ services.OtpServ) *OtpContImpl {
	return &OtpContImpl{OtpServ: otpServ}
}

func (cont *OtpContImpl) SendEmailVerification(context *gin.Context) {
	// Parse Request Body
	request := otp.SendRequest{}
	errParse := helpers.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exceptions.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.OtpServ.SendEmailVerification(request)
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

func (cont *OtpContImpl) SendPasswordReset(context *gin.Context) {
	// Parse Request Body
	request := otp.SendRequest{}
	errParse := helpers.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exceptions.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.OtpServ.SendPasswordReset(request)
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

func (cont *OtpContImpl) CheckOtpPassword(context *gin.Context) {
	// Parse Request Body
	request := otp.CheckRequest{}
	errParse := helpers.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exceptions.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	result, errServ := cont.OtpServ.CheckOtpPassword(request)
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

func (cont *OtpContImpl) CheckOtp(context *gin.Context) {
	// Parse Request Body
	request := otp.CheckRequest{}
	errParse := helpers.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exceptions.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.OtpServ.CheckOtp(request)
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
