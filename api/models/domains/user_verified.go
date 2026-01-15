package domains

import "time"

type UserVerified struct {
	VerifiedId int       `gorm:"primary_key;column:verified_id;auto_increment"`
	UserId     int       `gorm:"column:user_id"`
	Email      string    `gorm:"column:email"`
	Verified   bool      `gorm:"column:verified"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (UserVerified) TableName() string {
	return "userverified"
}
