package domains

import "time"

type Users struct {
	UserId          int          `gorm:"primary_key;column:user_id;auto_increment"`
	Role            UserRole     `gorm:"column:role;not null;default:'Member'"`
	FirstName       string       `gorm:"column:first_name;not null"`
	LastName        *string      `gorm:"column:last_name;null"`
	Email           string       `gorm:"column:email;not null;unique"`
	Password        string       `gorm:"column:password;not null"`
	PhotoProfileUrl *string      `gorm:"column:photo_profile_url;null"`
	Token           *string      `gorm:"column:token;null"`
	TokenExpire     *time.Time   `gorm:"column:token_expire;null;default:null"`
	CreatedAt       time.Time    `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       time.Time    `gorm:"column:updated_at;autoUpdateTime"`
	Verified        UserVerified `gorm:"foreignKey:user_id;references:user_id"`
}

type UserRole string

const (
	RoleAdmin  UserRole = "Admin"
	RoleMember UserRole = "Member"
)

func (Users) TableName() string {
	return "users"
}
