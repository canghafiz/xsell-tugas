package domains

import "time"

type PageLayout struct {
	LayoutID  int             `gorm:"primaryKey;column:layout_id;autoIncrement"`
	PageKey   string          `gorm:"column:page_key;not null;index"`
	SectionID int             `gorm:"column:section_id;not null"`
	SortOrder int             `gorm:"column:sort_order;not null;default:0"`
	IsActive  bool            `gorm:"column:is_active;default:true"`
	CreatedAt time.Time       `gorm:"column:created_at"`
	UpdatedAt time.Time       `gorm:"column:updated_at"`
	Section   *ContentSection `gorm:"foreignKey:SectionID;references:SectionID"`
}

func (PageLayout) TableName() string {
	return "page_layouts"
}
