package impl

import (
	"be/helpers"
	"be/models/repositories/member"
	"be/models/resources/category"
	"be/models/services"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type CategoryServMemberImpl struct {
	Db                 *gorm.DB
	CategoryRepoMember member.CategoryRepoMember
	RedisServ          services.RedisService
}

func NewCategoryServMemberImpl(db *gorm.DB, categoryRepoMember member.CategoryRepoMember, redisServ services.RedisService) *CategoryServMemberImpl {
	return &CategoryServMemberImpl{Db: db, CategoryRepoMember: categoryRepoMember, RedisServ: redisServ}
}

func (serv *CategoryServMemberImpl) GetWithSub() ([]category.Resource, error) {
	// Set cache key
	cacheKey := fmt.Sprintf("categories-with-sub")

	// Get From Cache
	var results []category.Resource
	cachedResult, err := helpers.GetFromCache(serv.RedisServ.GetData, cacheKey, &results)
	if err == nil && cachedResult != nil {
		if result, ok := cachedResult.([]category.Resource); ok {
			return result, nil
		}
	}

	// Call repo
	result, errRepo := serv.CategoryRepoMember.GetWithSub(serv.Db)
	if errRepo != nil {
		log.Printf("[CategoryRepoMember.GetCategories] error: %v", errRepo)
		return nil, fmt.Errorf("failed to get categories, please try again later")
	}

	// Save to cache
	if errCache := helpers.SetToCache(serv.RedisServ.SetData, cacheKey, category.ToResources(result), 60); errCache != nil {
		log.Printf("Failed to set related by category for key %s: %v", cacheKey, errCache)
	}

	return category.ToResources(result), nil
}
