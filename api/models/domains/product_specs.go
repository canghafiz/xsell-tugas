package domains

import "time"

type ProductSpecs struct {
	ProductSpecId         int                  `gorm:"primary_key;column:product_spec_id;auto_increment"`
	CategoryProductSpecId int                  `gorm:"column:category_product_spec_id"`
	ProductId             int                  `gorm:"column:product_id"`
	SpecValue             string               `gorm:"column:spec_value"`
	CreatedAt             time.Time            `gorm:"column:created_at"`
	UpdatedAt             time.Time            `gorm:"column:updated_at"`
	CategoryProductSpec   CategoryProductSpecs `gorm:"foreignKey:category_product_spec_id;references:category_product_spec_id"`
}

func (ProductSpecs) TableName() string {
	return "productspecs"
}
