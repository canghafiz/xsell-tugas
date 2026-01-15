package impl

import (
	"be/exceptions"
	"be/helpers"
	"be/models/requests/member/product"
	res "be/models/resources/product"
	"be/models/services/member"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProductContMemberImpl struct {
	ProductServ member.ProductServMember
}

func NewProductContMemberImpl(productServ member.ProductServMember) *ProductContMemberImpl {
	return &ProductContMemberImpl{ProductServ: productServ}
}

func (cont *ProductContMemberImpl) Create(context *gin.Context) {
	// Parse Request Body
	request := product.CreateProductRequest{}
	errParse := helpers.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exceptions.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.ProductServ.Create(request)
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

func (cont *ProductContMemberImpl) Update(context *gin.Context) {
	productIdStr := context.Param("productId")
	productId, _ := strconv.Atoi(productIdStr)

	// Parse Request Body
	request := product.UpdateProductRequest{
		ProductId: productId,
	}
	errParse := helpers.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exceptions.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.ProductServ.Update(request)
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

func (cont *ProductContMemberImpl) GetSingleBySlug(context *gin.Context) {
	productSlug := context.Param("productSlug")

	// Call Service
	result, errServ := cont.ProductServ.GetSingleBySlug(productSlug)
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

func (cont *ProductContMemberImpl) GetRelatedByCategory(context *gin.Context) {
	// Parse categoryIds as comma-separated string → []int
	categoryIdsStr := strings.Split(context.Query("categoryIds"), ",")
	var categoryIds []int
	for _, idStr := range categoryIdsStr {
		if idStr == "" {
			continue
		}
		id, err := strconv.Atoi(strings.TrimSpace(idStr))
		if err != nil || id <= 0 {
			continue // skip invalid IDs
		}
		categoryIds = append(categoryIds, id)
	}

	// Except id
	exceptIdStr := context.Param("exceptId")
	exceptId, _ := strconv.Atoi(exceptIdStr)

	// Parse limit with default fallback
	limit := 10 // default limit
	if limitStr := context.Query("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	latStr := context.Query("latitude")
	lat, _ := strconv.ParseFloat(latStr, 64)
	lngStr := context.Query("longitude")
	lng, _ := strconv.ParseFloat(lngStr, 64)

	if len(categoryIds) == 0 {
		response := helpers.ApiResponse{
			Success: true,
			Code:    200,
			Data:    []res.GeneralResource{},
		}
		_ = helpers.WriteToResponseBody(context, response.Code, response)
		return
	}

	// Call Service
	result, errServ := cont.ProductServ.GetRelatedByCategory(categoryIds, exceptId, limit, lat, lng)
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

func (cont *ProductContMemberImpl) GetByCategory(context *gin.Context) {
	// Parse categorySlug (required)
	categorySlug := context.Query("categorySlug")
	if categorySlug == "" {
		exceptions.ErrorHandler(context, fmt.Errorf("categorySlug is required"))
		return
	}

	// Parse sortBy (use "latest" as default to match repo logic)
	sortBy := context.DefaultQuery("sortBy", "latest")

	// Parse minPrice (default 0)
	minPrice := 0.0
	if minPriceStr := context.Query("minPrice"); minPriceStr != "" {
		if v, err := strconv.ParseFloat(minPriceStr, 64); err == nil {
			minPrice = v
		}
	}
	if minPrice < 0 {
		minPrice = 0
	}

	// Parse maxPrice (default: 999,999,999)
	maxPrice := 999999999.0
	if maxPriceStr := context.Query("maxPrice"); maxPriceStr != "" {
		if v, err := strconv.ParseFloat(maxPriceStr, 64); err == nil {
			if v > 0 {
				maxPrice = v
			}
		}
	}

	// Parse limit (default 20, max 100)
	limit := 20
	if limitStr := context.Query("limit"); limitStr != "" {
		if v, err := strconv.Atoi(limitStr); err == nil && v > 0 {
			if v > 100 {
				limit = 100
			} else {
				limit = v
			}
		}
	}

	latStr := context.Query("latitude")
	lat, _ := strconv.ParseFloat(latStr, 64)
	lngStr := context.Query("longitude")
	lng, _ := strconv.ParseFloat(lngStr, 64)

	// Parse optional subCategorySlug (supports multiple: ?subCategorySlug=phone&subCategorySlug=laptop)
	subCategorySlugs := context.QueryArray("subCategorySlug")
	// Note: QueryArray returns empty slice if not provided → safe for FilterModel

	// Build filter model
	model := &product.FilterModel{
		CategorySlug:    categorySlug,
		SubCategorySlug: subCategorySlugs, // can be empty
		SortBy:          sortBy,
		MinPrice:        minPrice,
		MaxPrice:        maxPrice,
		Limit:           limit,
		Lat:             lat,
		Lng:             lng,
	}

	// Call service
	result, errServ := cont.ProductServ.GetByCategory(model)
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

func (cont *ProductContMemberImpl) GetBySectionKey(context *gin.Context) {
	sectionKey := context.Param("sectionKey")

	categorySlug := context.Query("categorySlug")
	if categorySlug == "all" {
		categorySlug = ""
	}

	subCategorySlugs := context.QueryArray("subCategorySlug")

	if len(subCategorySlugs) == 1 && subCategorySlugs[0] == "all" {
		subCategorySlugs = []string{}
	}

	var filtered []string
	for _, sub := range subCategorySlugs {
		if sub != "all" && sub != "" {
			filtered = append(filtered, sub)
		}
	}
	subCategorySlugs = filtered

	// Parse sortBy
	sortBy := context.DefaultQuery("sortBy", "latest")

	// Parse minPrice
	minPrice := 0.0
	if minPriceStr := context.Query("minPrice"); minPriceStr != "" {
		if v, err := strconv.ParseFloat(minPriceStr, 64); err == nil && v > 0 {
			minPrice = v
		}
	}

	// Parse maxPrice
	maxPrice := 999999999.0
	if maxPriceStr := context.Query("maxPrice"); maxPriceStr != "" {
		if v, err := strconv.ParseFloat(maxPriceStr, 64); err == nil && v > 0 {
			maxPrice = v
		}
	}

	// Parse limit
	limit := 20
	if limitStr := context.Query("limit"); limitStr != "" {
		if v, err := strconv.Atoi(limitStr); err == nil && v > 0 {
			if v > 100 {
				limit = 100
			} else {
				limit = v
			}
		}
	}

	// Parse coordinates
	lat, _ := strconv.ParseFloat(context.Query("latitude"), 64)
	lng, _ := strconv.ParseFloat(context.Query("longitude"), 64)

	// Build filter model
	model := &product.FilterModel{
		CategorySlug:    categorySlug,
		SubCategorySlug: subCategorySlugs,
		SortBy:          sortBy,
		MinPrice:        minPrice,
		MaxPrice:        maxPrice,
		Limit:           limit,
		Lat:             lat,
		Lng:             lng,
	}

	// Call service
	result, errServ := cont.ProductServ.GetBySectionKey(sectionKey, model)
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

func (cont *ProductContMemberImpl) Search(context *gin.Context) {
	title := context.DefaultQuery("title", "")

	categorySlug := context.Query("categorySlug")
	if categorySlug == "all" {
		categorySlug = ""
	}

	subCategorySlugs := context.QueryArray("subCategorySlug")

	if len(subCategorySlugs) == 1 && subCategorySlugs[0] == "all" {
		subCategorySlugs = []string{}
	}

	var filtered []string
	for _, sub := range subCategorySlugs {
		if sub != "all" && sub != "" {
			filtered = append(filtered, sub)
		}
	}
	subCategorySlugs = filtered

	// Parse sortBy
	sortBy := context.DefaultQuery("sortBy", "latest")

	// Parse minPrice
	minPrice := 0.0
	if minPriceStr := context.Query("minPrice"); minPriceStr != "" {
		if v, err := strconv.ParseFloat(minPriceStr, 64); err == nil && v > 0 {
			minPrice = v
		}
	}

	// Parse maxPrice
	maxPrice := 999999999.0
	if maxPriceStr := context.Query("maxPrice"); maxPriceStr != "" {
		if v, err := strconv.ParseFloat(maxPriceStr, 64); err == nil && v > 0 {
			maxPrice = v
		}
	}

	// Parse limit
	limit := 20
	if limitStr := context.Query("limit"); limitStr != "" {
		if v, err := strconv.Atoi(limitStr); err == nil && v > 0 {
			if v > 100 {
				limit = 100
			} else {
				limit = v
			}
		}
	}

	// Parse coordinates
	lat, _ := strconv.ParseFloat(context.Query("latitude"), 64)
	lng, _ := strconv.ParseFloat(context.Query("longitude"), 64)

	// Build filter model
	model := &product.FilterModel{
		CategorySlug:    categorySlug,
		SubCategorySlug: subCategorySlugs,
		SortBy:          sortBy,
		MinPrice:        minPrice,
		MaxPrice:        maxPrice,
		Limit:           limit,
		Lat:             lat,
		Lng:             lng,
	}

	// Call service
	result, errServ := cont.ProductServ.Search(title, model)
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

func (cont *ProductContMemberImpl) GetProductsByUserId(context *gin.Context) {
	userIdStr := context.Param("userId")
	userId, _ := strconv.Atoi(userIdStr)

	sortBy := context.Query("sortBy")
	model := product.FilterMyAds{
		SortBy: sortBy,
	}

	// Call service
	result, errServ := cont.ProductServ.GetProductsByUserId(userId, model)
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

func (cont *ProductContMemberImpl) UpdateStatus(context *gin.Context) {
	productIdStr := context.Param("productId")
	productId, _ := strconv.Atoi(productIdStr)

	// Parse Request Body
	request := product.UpdateStatusRequest{}
	request.ProductId = productId
	errParse := helpers.ReadFromRequestBody(context, &request)
	if errParse != nil {
		exceptions.ErrorHandler(context, errParse)
		return
	}

	// Call Service
	errServ := cont.ProductServ.UpdateStatus(request)
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

func (cont *ProductContMemberImpl) UpdateViewCount(context *gin.Context) {
	productIdStr := context.Param("productId")
	productId, _ := strconv.Atoi(productIdStr)

	// Call Service
	errServ := cont.ProductServ.UpdateViewCount(productId)
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

func (cont *ProductContMemberImpl) Delete(context *gin.Context) {
	productIdStr := context.Param("productId")
	productId, _ := strconv.Atoi(productIdStr)

	// Call Service
	errServ := cont.ProductServ.Delete(productId)
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
