package domains

import "time"

type ProductImages struct {
	ImageId       int       `gorm:"primary_key;column:image_id;auto_increment"`
	ProductId     int       `gorm:"column:product_id"`
	ImageUrl      string    `gorm:"column:image_url"`
	IsPrimary     bool      `gorm:"column:is_primary"`
	OrderSequence int       `gorm:"column:order_sequence"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}

func (ProductImages) TableName() string {
	return "productimages"
}
