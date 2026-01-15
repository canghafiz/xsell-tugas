package domains

import (
	"strings"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type SubCategories struct {
	SubCategoryId    int        `gorm:"primary_key;column:sub_category_id;auto_increment"`
	ParentCategoryId int        `gorm:"column:parent_category_id"`
	Name             string     `gorm:"column:sub_category_name"`
	Slug             string     `gorm:"column:sub_category_slug"`
	Description      *string    `gorm:"column:description"`
	Icon             string     `gorm:"column:icon"`
	Category         Categories `gorm:"foreignKey:parent_category_id;references:category_id"`
}

func (p *SubCategories) BeforeCreate(tx *gorm.DB) error {
	p.generateSlug()
	return nil
}

func (p *SubCategories) BeforeUpdate(tx *gorm.DB) error {
	p.generateSlug()
	return nil
}

func (p *SubCategories) generateSlug() {
	if p.Name == "" {
		return
	}
	s := slug.Make(strings.TrimSpace(p.Name))
	p.Slug = s
}

func (SubCategories) TableName() string {
	return "subcategories"
}
