package impl

import (
	"be/helpers"
	"be/models/repositories/member"
	"be/models/resources/sub_category"
	"be/models/services"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type SubCategorySerMemberImpl struct {
	Db              *gorm.DB
	SubCategoryRepo member.SubCategoryRepoMember
	RedisServ       services.RedisService
}

func NewSubCategorySerMemberImpl(db *gorm.DB, subCategoryRepo member.SubCategoryRepoMember, redisServ services.RedisService) *SubCategorySerMemberImpl {
	return &SubCategorySerMemberImpl{Db: db, SubCategoryRepo: subCategoryRepo, RedisServ: redisServ}
}

func (serv *SubCategorySerMemberImpl) GetByCategoryParent(categorySlug string) ([]sub_category.Resource, error) {
	cacheKey := fmt.Sprintf("sub_category:by:category:%v", categorySlug)

	// Get From Cache
	var results []sub_category.Resource
	cachedResult, err := helpers.GetFromCache(serv.RedisServ.GetData, cacheKey, &results)
	if err == nil && cachedResult != nil {
		if result, ok := cachedResult.([]sub_category.Resource); ok {
			return result, nil
		}
	}

	// Call repo
	result, errRepo := serv.SubCategoryRepo.GetByCategoryParent(serv.Db, categorySlug)
	if errRepo != nil {
		log.Printf("[SubCategoryRepoMember.GetByCategoryParent] error: %v", errRepo)
		return nil, fmt.Errorf("failed to get by category parent, please try again later")
	}

	// Save to cache
	if errCache := helpers.SetToCache(serv.RedisServ.SetData, cacheKey, sub_category.ToResources(result), 1440); errCache != nil {
		log.Printf("Failed to set related by category for key %s: %v", cacheKey, errCache)
	}

	return sub_category.ToResources(result), nil
}
