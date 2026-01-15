package domains

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type ContentSection struct {
	SectionID   int           `gorm:"primaryKey;column:section_id;autoIncrement"`
	SectionKey  string        `gorm:"column:section_key;not null;uniqueIndex"`
	Title       string        `gorm:"column:title;not null"`
	Subtitle    string        `gorm:"column:subtitle"`
	Url         string        `gorm:"column:url;not null"`
	SectionType string        `gorm:"column:section_type;not null"` // dynamic, fixed, predefined
	Config      ContentConfig `gorm:"column:config;type:jsonb;not null"`
	IsActive    bool          `gorm:"column:is_active;default:true"`
	SortOrder   int           `gorm:"column:sort_order;default:0"`
	CreatedAt   time.Time     `gorm:"column:created_at"`
	UpdatedAt   time.Time     `gorm:"column:updated_at"`
}

type ContentConfig struct {
	Limit      int                    `json:"limit,omitempty"`
	Filters    map[string]interface{} `json:"filters,omitempty"`
	ProductIDs []int                  `json:"product_ids,omitempty"`
	Strategy   string                 `json:"strategy,omitempty"`
}

// Value implements driver.Valuer for JSONB storage
func (c ContentConfig) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *ContentConfig) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan JSONB value")
	}
	return json.Unmarshal(bytes, c)
}

func (ContentSection) TableName() string {
	return "content_sections"
}
