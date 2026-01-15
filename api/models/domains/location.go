package domains

import "time"

type Location struct {
	LocationId int       `gorm:"primary_key;column:location_id;auto_increment"`
	ProductId  int       `gorm:"column:product_id"`
	Latitude   float64   `gorm:"column:latitude"`
	Longitude  float64   `gorm:"column:longitude"`
	Address    string    `gorm:"column:address"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (r *Location) TableName() string {
	return "location"
}
