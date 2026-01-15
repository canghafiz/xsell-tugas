package domains

import (
	"math/rand"
	"strings"
	"time"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Products struct {
	ProductId     int              `gorm:"primary_key;column:product_id;auto_increment"`
	ListingUserId int              `gorm:"column:listing_user_id"`
	SubCategoryId int              `gorm:"column:sub_category_id"`
	Title         string           `gorm:"column:title"`
	ProductSlug   string           `gorm:"column:product_slug;index:idx_product_slug,unique"`
	Description   string           `gorm:"column:description"`
	Price         float64          `gorm:"column:price"`
	Condition     ProductCondition `gorm:"column:condition"`
	Status        ProductStatus    `gorm:"column:status"`
	ViewCount     int              `gorm:"column:view_count"`
	CreatedAt     time.Time        `gorm:"column:created_at"`
	UpdatedAt     time.Time        `gorm:"column:updated_at"`
	SubCategory   SubCategories    `gorm:"foreignKey:sub_category_id;references:sub_category_id"`
	ProductImages []ProductImages  `gorm:"foreignKey:product_id;references:product_id"`
	ProductSpecs  []ProductSpecs   `gorm:"foreignKey:product_id;references:product_id"`
	Location      Location         `gorm:"foreignKey:product_id;references:product_id"`
	Listing       Users            `gorm:"foreignKey:listing_user_id;references:user_id"`
	Wishlists     []Wishlist       `gorm:"foreignKey:product_id;references:product_id"`
}

type ProductWithWishlist struct {
	Product       Products `gorm:"embedded"`
	TotalWishlist int      `gorm:"column:total_wishlist"`
}

func (p *Products) BeforeCreate(tx *gorm.DB) error {
	p.generateSlug()
	return nil
}

func (p *Products) BeforeUpdate(tx *gorm.DB) error {
	if p.ProductSlug != "" {
		return nil
	}

	p.generateSlug()
	return nil
}

func (p *Products) generateSlug() {
	if p.Title == "" {
		return
	}

	baseSlug := slug.Make(strings.TrimSpace(p.Title))
	uniqueCode := generateRandomNumber(6)

	p.ProductSlug = baseSlug + "-" + uniqueCode
}

func generateRandomNumber(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	digits := "0123456789"
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = digits[r.Intn(len(digits))]
	}

	return string(result)
}

type ProductCondition string

const (
	New        ProductCondition = "New"
	LikeNew    ProductCondition = "Like New"
	Good       ProductCondition = "Good"
	GoodQuite  ProductCondition = "Good Quite"
	NeedRepair ProductCondition = "Needs Repair"
)

type ProductStatus string

const (
	Available ProductStatus = "Available"
	SoldOut   ProductStatus = "Sold out"
)

func (Products) TableName() string {
	return "products"
}
