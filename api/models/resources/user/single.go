package user

import (
	"be/models/domains"
	"time"
)

type SingleResource struct {
	UserId       int              `json:"user_id"`
	Email        string           `json:"email"`
	Role         domains.UserRole `json:"role"`
	FirstName    string           `json:"first_name"`
	LastName     *string          `json:"last_name"`
	PhotoProfile *string          `json:"photo_profile"`
	CreatedAt    time.Time        `json:"created_at"`
}

func ToSingleResource(user *domains.Users) *SingleResource {
	return &SingleResource{
		UserId:       user.UserId,
		Email:        user.Email,
		Role:         user.Role,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		PhotoProfile: user.PhotoProfileUrl,
		CreatedAt:    user.CreatedAt,
	}
}
