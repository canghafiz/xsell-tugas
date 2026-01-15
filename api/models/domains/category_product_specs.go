package domains

import "time"

type CategoryProductSpecs struct {
	CategoryProductSpecId int       `gorm:"primary_key;column:category_product_spec_id;auto_increment"`
	SubCategoryId         int       `gorm:"column:sub_category_id"`
	IsMainSpec            bool      `gorm:"column:is_main_spec"`
	SpecTypeTitle         string    `gorm:"column:spec_type_title"`
	SpecName              string    `gorm:"column:spec_name"`
	CreatedAt             time.Time `gorm:"column:created_at"`
	UpdatedAt             time.Time `gorm:"column:update_at"`
}

func (CategoryProductSpecs) TableName() string {
	return "categoryproductspecs"
}
