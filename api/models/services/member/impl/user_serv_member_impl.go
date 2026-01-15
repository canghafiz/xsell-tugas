package impl

import (
	"be/helpers"
	"be/models/domains"
	"be/models/repositories"
	"be/models/requests/user"
	user2 "be/models/resources/user"
	"fmt"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserServMemberImpl struct {
	UserRepo  repositories.UserRepo
	DB        *gorm.DB
	Validator *validator.Validate
	JwtKey    string
}

func NewUserServMemberImpl(userRepo repositories.UserRepo, DB *gorm.DB, validator *validator.Validate, jwtKey string) *UserServMemberImpl {
	return &UserServMemberImpl{UserRepo: userRepo, DB: DB, Validator: validator, JwtKey: jwtKey}
}

func (serv *UserServMemberImpl) Register(request user.RegisterRequest) error {
	errValidator := helpers.ErrValidator(request, serv.Validator)
	if errValidator != nil {
		return errValidator
	}

	// Find user by email
	findUser, _ := serv.UserRepo.FindByEmail(serv.DB, request.Email)
	if findUser != nil {
		return fmt.Errorf("user already exists")
	}

	// Define model
	model := user.RegisterRequestToDomain(request)
	model.Role = domains.RoleMember

	// Call repo
	errCreate := serv.UserRepo.Create(serv.DB, *model)
	if errCreate != nil {
		log.Printf("[UserRepo.Create] error: %v", errCreate)
		return fmt.Errorf("failed to register, please try again later")
	}

	return nil

}

func (serv *UserServMemberImpl) Login(request user.LoginRequest) (*string, error) {
	errValidator := helpers.ErrValidator(request, serv.Validator)
	if errValidator != nil {
		return nil, errValidator
	}

	// Find user by email
	findUser, errFindUser := serv.UserRepo.FindByEmail(serv.DB, request.Email)
	if findUser == nil || errFindUser != nil {
		log.Printf("[UserRepo.FindByEmail] error: %v", errFindUser)
		return nil, fmt.Errorf("user not found")
	}

	// Cek user role
	if findUser.Role != domains.RoleMember {
		return nil, fmt.Errorf("user not found")
	}

	// Check user email is verified
	if !findUser.Verified.Verified {
		return nil, fmt.Errorf("your email is not verified")
	}

	// Define model
	model := user.LoginRequestToDomains(request)

	// Check Password
	valid, errPw := serv.UserRepo.CheckPasswordValid(serv.DB, model)
	if errPw != nil {
		log.Printf("[UserRepo.CheckPasswordValid] error: %v", errPw)
		return nil, fmt.Errorf("failed to login, please try again later")
	}

	if !valid {
		return nil, fmt.Errorf("password invalid")
	}

	// User Resource
	resource := user2.ToSingleResource(findUser)

	// Create JWT Token
	duration := time.Hour * 24
	jwt, errJwt := helpers.GenerateJWT(serv.JwtKey, duration, &resource)
	if errJwt != nil {
		log.Printf("[GenerateJWT] error: %v", errJwt)
		return nil, fmt.Errorf("failed to login, please try again later")
	}

	// Update user token
	tokenModel := domains.Users{
		Email: request.Email,
		Token: &jwt,
	}
	errUpToken := serv.UserRepo.UpdateToken(serv.DB, tokenModel)
	if errUpToken != nil {
		log.Printf("[UserRepo.UpdateToken] error: %v", errJwt)
		return nil, fmt.Errorf("failed to login, please try again later")
	}

	return &jwt, nil
}

func (serv *UserServMemberImpl) Logout(email string) error {
	// Call repo
	err := serv.UserRepo.ResetToken(serv.DB, email)
	if err != nil {
		log.Printf("[UserRepo.ResetToken] error: %v", err)
		return fmt.Errorf("failed to logout, please try again later")
	}

	return nil
}

func (serv *UserServMemberImpl) UpdateData(request user.UpdateDataRequest) (*user2.SingleResource, error) {
	errValidator := helpers.ErrValidator(request, serv.Validator)
	if errValidator != nil {
		return nil, errValidator
	}

	// Define model
	model := user.UpdateDataRequestToDomain(request)

	// Call repo
	result, err := serv.UserRepo.Update(serv.DB, model)
	if err != nil {
		log.Printf("[UserRepo.Update] error: %v", err)
		return nil, fmt.Errorf("failed to update, please try again later")
	}

	return user2.ToSingleResource(result), nil
}
