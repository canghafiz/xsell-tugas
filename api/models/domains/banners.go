package domains

import "time"

type Banners struct {
	BannerId    int       `gorm:"column:banner_id;primaryKey;autoIncrement"`
	ImageURL    string    `gorm:"column:image_url;not null"`
	Sequence    int       `gorm:"column:sequence;default:1"`
	Link        string    `gorm:"column:link"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Banners) TableName() string {
	return "banners"
}
