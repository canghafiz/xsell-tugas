package services

import (
	"be/models/requests/user"
	user2 "be/models/resources/user"
)

type UserServ interface {
	ChangePw(request user.ChangePwRequest, jwtValue string) error
	GetByUserId(userId int) (*user2.SingleResource, error)
}
