package impl

import (
	"be/helpers"
	"be/models/repositories/member"
	"be/models/resources/product_spec"
	"be/models/services"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type ProductSpecServMemberImpl struct {
	Db              *gorm.DB
	ProductSpecRepo member.ProductSpecRepoMember
	RedisServ       services.RedisService
}

func NewProductSpecServMemberImpl(db *gorm.DB, productSpecRepo member.ProductSpecRepoMember, redisServ services.RedisService) *ProductSpecServMemberImpl {
	return &ProductSpecServMemberImpl{Db: db, ProductSpecRepo: productSpecRepo, RedisServ: redisServ}
}

func (serv *ProductSpecServMemberImpl) GetProductSpecBySubId(subId int) ([]product_spec.Resource, error) {
	// Set cache key
	cacheKey := fmt.Sprintf("product-spec:subId:%d:", subId)

	// Get From Cache
	var results []product_spec.Resource
	cachedResult, err := helpers.GetFromCache(serv.RedisServ.GetData, cacheKey, &results)
	if err == nil && cachedResult != nil {
		if result, ok := cachedResult.([]product_spec.Resource); ok {
			return result, nil
		}
	}

	// Call repo
	result, errRepo := serv.ProductSpecRepo.GetProductSpecBySubId(serv.Db, subId)
	if errRepo != nil {
		log.Printf("[ProductSpecRepoMember.GetProductSpecBySubId] error: %v", errRepo)
		return nil, fmt.Errorf("failed to get product spec by sub id, please try again later")
	}

	// Save to cache
	if errCache := helpers.SetToCache(serv.RedisServ.SetData, cacheKey, product_spec.ToResources(result), 1440); errCache != nil {
		log.Printf("failed to get product spec by sub id for key %s: %v", cacheKey, errCache)
	}

	return product_spec.ToResources(result), nil
}
