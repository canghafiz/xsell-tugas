package impl

import (
	"be/models/domains"

	"gorm.io/gorm"
)

type ProductSpecRepoMemberImpl struct {
}

func NewProductSpecRepoMemberImpl() *ProductSpecRepoMemberImpl {
	return &ProductSpecRepoMemberImpl{}
}

func (p ProductSpecRepoMemberImpl) GetProductSpecBySubId(db *gorm.DB, subId int) ([]domains.CategoryProductSpecs, error) {
	var results []domains.CategoryProductSpecs
	err := db.Where("sub_category_id = ?", subId).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
