package impl

import (
	"be/exceptions"
	"be/helpers"
	_map "be/models/requests/map"
	"be/models/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MapContImpl struct {
	MapServ services.MapServ
}

func NewMapContImpl(mapServ services.MapServ) *MapContImpl {
	return &MapContImpl{MapServ: mapServ}
}

func (cont *MapContImpl) GetAddress(context *gin.Context) {
	lat := context.Query("lat")
	lon := context.Query("lon")

	// Call Service
	result, errServ := cont.MapServ.GetAddress(_map.GetAddressRequest{
		Lat: lat,
		Lon: lon,
	})
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

func (cont *MapContImpl) AutoComplete(context *gin.Context) {
	q := context.Query("q")
	limitStr := context.DefaultQuery("limit", "5")
	limit, _ := strconv.Atoi(limitStr)

	// Call Service
	result, errServ := cont.MapServ.AutoComplete(_map.AutocompleteRequest{
		Q:     q,
		Limit: limit,
	})
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
