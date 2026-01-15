package domains

type MetaSeo struct {
	MetaId    int    `gorm:"primary_key;column:meta_id;auto_increment"`
	PageKey   string `gorm:"column:page_key"`
	MetaName  string `gorm:"column:meta_name"`
	MetaValue string `gorm:"column:meta_value"`
	IsActive  bool   `gorm:"column:is_active"`
}

func (MetaSeo) TableName() string {
	return "content_meta_seo"
}
