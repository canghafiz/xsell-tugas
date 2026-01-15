package impl

import (
	"be/helpers"
	"be/models/repositories"
	"be/models/requests/user"
	user2 "be/models/resources/user"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserServImpl struct {
	Db        *gorm.DB
	Validator *validator.Validate
	UserRepo  repositories.UserRepo
	JwtKey    string
}

func NewUserServImpl(db *gorm.DB, validator *validator.Validate, userRepo repositories.UserRepo, jwtKey string) *UserServImpl {
	return &UserServImpl{Db: db, Validator: validator, UserRepo: userRepo, JwtKey: jwtKey}
}

func (serv *UserServImpl) ChangePw(request user.ChangePwRequest, jwtValue string) error {
	errValidator := helpers.ErrValidator(request, serv.Validator)
	if errValidator != nil {
		return errValidator
	}

	// Decode Token
	result, errDecode := helpers.DecodeJWT(jwtValue, serv.JwtKey)
	if errDecode != nil {
		log.Printf("Decode jwt error: %v", errDecode)
		return fmt.Errorf("failed to change pw, please try again later")
	}

	// Extract user id
	userId, ok := result["user_id"].(float64)
	if !ok || userId == 0 {
		log.Printf("Decode jwt error: %v", errDecode)
		return fmt.Errorf("failed to change pw, please try again later")
	}

	// Define Model
	model := user.ChangePwToDomain(request)
	model.UserId = int(userId)

	// Check confirmation password
	if request.Password != request.ConfirmPassword {
		return fmt.Errorf("the password does not match")
	}

	// Call repo
	errRepo := serv.UserRepo.ChangePassword(serv.Db, model)
	if errRepo != nil {
		log.Printf("[UserRepo.ChangePassword] error: %v", errDecode)
		return fmt.Errorf("failed to change pw, please try again later")
	}

	return nil
}

func (serv *UserServImpl) GetByUserId(userId int) (*user2.SingleResource, error) {
	// Call repo
	result, errRepo := serv.UserRepo.GetByUserId(serv.Db, userId)
	if errRepo != nil {
		log.Printf("[UserRepo.GetByUserId] error: %v", errRepo)
		return nil, fmt.Errorf("failed to get user by user id, please try again later")
	}
	return user2.ToSingleResource(result), nil
}
