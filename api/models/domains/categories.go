package domains

import (
	"be/helpers"
	"time"

	"gorm.io/gorm"
)

type Categories struct {
	CategoryId    int             `gorm:"primary_key;column:category_id;auto_increment"`
	CategoryName  string          `gorm:"column:category_name"`
	CategorySlug  string          `gorm:"column:category_slug"`
	Description   string          `gorm:"column:description"`
	Icon          string          `gorm:"column:icon"`
	CreatedAt     time.Time       `gorm:"column:created_at"`
	UpdatedAt     time.Time       `gorm:"column:updated_at"`
	SubCategories []SubCategories `gorm:"foreignKey:parent_category_id;references:category_id"`
}

func (c *Categories) BeforeCreate(tx *gorm.DB) error {
	c.CategorySlug = helpers.GenerateSlug(c.CategoryName)
	return nil
}

func (c *Categories) BeforeUpdate(tx *gorm.DB) error {
	c.CategorySlug = helpers.GenerateSlug(c.CategoryName)
	return nil
}

func (Categories) TableName() string {
	return "categories"
}
