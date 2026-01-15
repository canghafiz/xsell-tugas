package impl

import (
	"be/helpers"
	"be/models/domains"
	"be/models/repositories/member"
	"be/models/requests/member/product"
	res "be/models/resources/product"
	"be/models/services"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProductServMemberImpl struct {
	Db          *gorm.DB
	Validator   *validator.Validate
	ProductRepo member.ProductRepoMember
	RedisServ   services.RedisService
	StorageServ services.FileServ
}

func NewProductServMemberImpl(db *gorm.DB, validator *validator.Validate, productRepo member.ProductRepoMember, redisServ services.RedisService, storageServ services.FileServ) *ProductServMemberImpl {
	return &ProductServMemberImpl{Db: db, Validator: validator, ProductRepo: productRepo, RedisServ: redisServ, StorageServ: storageServ}
}

func (serv *ProductServMemberImpl) Create(request product.CreateProductRequest) error {
	errValidator := helpers.ErrValidator(request, serv.Validator)
	if errValidator != nil {
		return errValidator
	}

	model := product.CreateProductRequestToDomain(request)

	// Call repo
	err := serv.ProductRepo.Create(serv.Db, *model)
	if err != nil {
		log.Printf("[ProductRepoMember.Create] error: %v", err)
		return fmt.Errorf("failed to create product, please try again later")
	}

	return nil
}

func (serv *ProductServMemberImpl) Update(request product.UpdateProductRequest) error {
	errValidator := helpers.ErrValidator(request, serv.Validator)
	if errValidator != nil {
		return errValidator
	}

	model := product.UpdateProductRequestToDomain(request)

	// Call repo
	err := serv.ProductRepo.Update(serv.Db, *model)
	if err != nil {
		log.Printf("[ProductRepoMember.Update] error: %v", err)
		return fmt.Errorf("failed to update product, please try again later")
	}

	return nil
}

func (serv *ProductServMemberImpl) GetSingleBySlug(productSlug string) (*res.SingleProductResource, error) {
	// Call repo
	result, err := serv.ProductRepo.GetSingleBySlug(serv.Db, productSlug)
	if err != nil {
		log.Printf("[ProductRepoMember.GetSingleBySlug] error: %v", err)
		return nil, fmt.Errorf("failed to get single product by slug, please try again later")
	}

	return res.ToSingleProductResource(result), nil
}

func (serv *ProductServMemberImpl) GetRelatedByCategory(categoryIds []int, exceptProductId, limit int, lat, lng float64) ([]res.GeneralResource, error) {
	// Set cache key
	cacheKey := fmt.Sprintf("product-related:exceptId:%d:categoryIds:%d:limit:%d", exceptProductId, categoryIds, limit)

	// Get From Cache
	var results []res.GeneralResource
	cachedResult, err := helpers.GetFromCache(serv.RedisServ.GetData, cacheKey, &results)
	if err == nil && cachedResult != nil {
		if result, ok := cachedResult.([]res.GeneralResource); ok {
			return result, nil
		}
	}

	// Call repo
	result, errRepo := serv.ProductRepo.GetRelatedByCategory(serv.Db, categoryIds, exceptProductId, limit, lat, lng)
	if errRepo != nil {
		log.Printf("[ProductRepoMember.GetRelatedByCategory] error: %v", errRepo)
		return nil, fmt.Errorf("failed to get related products by categories, please try again later")
	}

	// Save to cache
	if errCache := helpers.SetToCache(serv.RedisServ.SetData, cacheKey, res.ToGeneralResources(result), 1440); errCache != nil {
		log.Printf("Failed to set related by category for key %s: %v", cacheKey, errCache)
	}

	return res.ToGeneralResources(result), nil
}

func (serv *ProductServMemberImpl) GetBySectionKey(
	key string,
	filter *product.FilterModel,
) ([]res.GeneralResource, error) {

	if key == "" {
		return []res.GeneralResource{}, nil
	}

	cacheKey := fmt.Sprintf(
		"product-section:%s:lat:%.6f:lng:%.6f",
		key,
		filter.Lat,
		filter.Lng,
	)

	var cached []res.GeneralResource
	cachedResult, err := helpers.GetFromCache(serv.RedisServ.GetData, cacheKey, &cached)
	if err == nil && cachedResult != nil {
		if result, ok := cachedResult.([]res.GeneralResource); ok {
			return result, nil
		}
	}

	products, errRepo := serv.ProductRepo.GetBySectionKey(
		serv.Db,
		key,
		*filter,
	)
	if errRepo != nil {
		log.Printf("[ProductRepoMember.GetBySectionKey] error: %v", errRepo)
		return nil, fmt.Errorf("failed to get products by section")
	}

	resources := res.ToGeneralResources(products)

	if errCache := helpers.SetToCache(
		serv.RedisServ.SetData,
		cacheKey,
		resources,
		60,
	); errCache != nil {
		log.Printf("Failed to cache section %s: %v", key, errCache)
	}

	return resources, nil
}

func (serv *ProductServMemberImpl) GetByCategory(filter *product.FilterModel) ([]res.GeneralResource, error) {
	// Set cache key
	cacheKey := buildCategoryCacheKey(filter)

	// Get From Cache
	var results []res.GeneralResource
	cachedResult, err := helpers.GetFromCache(serv.RedisServ.GetData, cacheKey, &results)
	if err == nil && cachedResult != nil {
		if result, ok := cachedResult.([]res.GeneralResource); ok {
			return result, nil
		}
	}

	// Save to cache
	if errCache := helpers.SetToCache(serv.RedisServ.SetData, cacheKey, results, 1440); errCache != nil {
		log.Printf("Failed to set get by category for key %s: %v", cacheKey, errCache)
	}

	// Call repo
	result, errRepo := serv.ProductRepo.GetByCategory(serv.Db, *filter)
	if errRepo != nil {
		log.Printf("[ProductRepoMember.GetByCategory] error: %v", errRepo)
		return nil, fmt.Errorf("failed to get products by category, please try again later")
	}

	return res.ToGeneralResources(result), nil
}

func (serv *ProductServMemberImpl) Search(title string, filter *product.FilterModel) ([]res.GeneralResource, error) {
	// Build cache key based on search parameters
	cacheKey := fmt.Sprintf(
		"product-search:title:%s:category:%s:subCategory:%v:sortBy:%s:minPrice:%s:maxPrice:%s:limit:%d",
		url.QueryEscape(title),
		filter.CategorySlug,
		filter.SubCategorySlug,
		filter.SortBy,
		filter.MinPrice,
		filter.MaxPrice,
		filter.Limit,
	)

	// Get from cache
	var results []res.GeneralResource
	cachedResult, err := helpers.GetFromCache(serv.RedisServ.GetData, cacheKey, &results)
	if err == nil && cachedResult != nil {
		if result, ok := cachedResult.([]res.GeneralResource); ok {
			return result, nil
		}
	}

	// Call repository
	result, errRepo := serv.ProductRepo.Search(serv.Db, title, filter)
	if errRepo != nil {
		log.Printf("[ProductServMemberImpl.Search] error: %v", errRepo)
		return nil, fmt.Errorf("failed to search products, please try again later")
	}

	// Save to cache (only if results exist to avoid caching empty results indefinitely)
	if len(result) > 0 {
		if errCache := helpers.SetToCache(serv.RedisServ.SetData, cacheKey, res.ToGeneralResources(result), 1440); errCache != nil {
			log.Printf("Failed to set search cache for key %s: %v", cacheKey, errCache)
		}
	}

	return res.ToGeneralResources(result), nil
}

func (serv *ProductServMemberImpl) GetProductsByUserId(userId int, filter product.FilterMyAds) ([]res.MyProductResource, error) {
	// Set cache key
	cacheKey := fmt.Sprintf("product-by-userId:%s", userId)

	// Get From Cache
	var results []res.MyProductResource
	cachedResult, err := helpers.GetFromCache(serv.RedisServ.GetData, cacheKey, &results)
	if err == nil && cachedResult != nil {
		if result, ok := cachedResult.([]res.MyProductResource); ok {
			return result, nil
		}
	}

	// Call repo
	result, errRepo := serv.ProductRepo.GetProductsByUserId(serv.Db, userId, filter)
	if errRepo != nil {
		log.Printf("[ProductRepoMember.GetRelatedByCategory] error: %v", errRepo)
		return nil, fmt.Errorf("failed to get related products by categories, please try again later")
	}

	// Save to cache
	if errCache := helpers.SetToCache(serv.RedisServ.SetData, cacheKey, res.ToMyProductResources(result), 1440); errCache != nil {
		log.Printf("Failed to set related by category for key %s: %v", cacheKey, errCache)
	}

	return res.ToMyProductResources(result), nil
}

func (serv *ProductServMemberImpl) UpdateStatus(request product.UpdateStatusRequest) error {
	model := product.UpdateStatusRequestToDomain(request)

	// Call repo
	err := serv.ProductRepo.UpdateStatus(serv.Db, model)
	if err != nil {
		log.Printf("[ProductRepoMember.UpdateStatus] error: %v", err)
		return fmt.Errorf("failed to update status product, please try again later")
	}
	return nil
}

func (serv *ProductServMemberImpl) UpdateViewCount(productId int) error {
	// Call repo
	err := serv.ProductRepo.UpdateViewCount(serv.Db, productId)
	if err != nil {
		log.Printf("[ProductRepoMember.UpdateViewCount] error: %v", err)
		return fmt.Errorf("failed to update view_count product, please try again later")
	}
	return nil
}

func (serv *ProductServMemberImpl) Delete(productId int) error {
	// Call repo
	images, err := serv.ProductRepo.Delete(serv.Db, domains.Products{
		ProductId: productId,
	})
	if err != nil {
		log.Printf("[ProductRepoMember.Delete] error: %v", err)
		return fmt.Errorf("failed to delete by product id, please try again later")
	}

	// Delete files
	for _, file := range images {
		errDelFile := serv.StorageServ.DeleteFile(file)
		if errDelFile != nil {
			log.Printf("[FileServ.DeleteFile] error: %v", errDelFile)
		}
	}

	return nil
}

func buildCategoryCacheKey(filter *product.FilterModel) string {
	sortBy := strings.ToLower(filter.SortBy)

	cacheKey := fmt.Sprintf(
		"product-filter:category-slug:%s:sortBy:%s:minPrice:%.2f:maxPrice:%.2f:limit:%d",
		filter.CategorySlug,
		sortBy,
		filter.MinPrice,
		filter.MaxPrice,
		filter.Limit,
	)

	return cacheKey
}
