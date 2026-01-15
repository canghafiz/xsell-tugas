package member

import (
	"be/models/requests/user"
	user2 "be/models/resources/user"
)

type UserServMember interface {
	Register(request user.RegisterRequest) error
	Login(request user.LoginRequest) (*string, error)
	Logout(email string) error
	UpdateData(request user.UpdateDataRequest) (*user2.SingleResource, error)
}
