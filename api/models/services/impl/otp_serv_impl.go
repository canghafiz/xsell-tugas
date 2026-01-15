package impl

import (
	"be/helpers"
	"be/models/domains"
	"be/models/repositories"
	"be/models/requests/member/otp"
	user2 "be/models/resources/user"
	"be/models/services"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type OtpServImpl struct {
	Db             *gorm.DB
	Validator      *validator.Validate
	OtpRepo        repositories.OtpRepo
	SmtpServ       services.SmtpServ
	UserRepo       repositories.UserRepo
	UserVerifyRepo repositories.UserVerifyRepo
	JwtKey         string
}

func NewOtpServImpl(db *gorm.DB, validator *validator.Validate, otpRepo repositories.OtpRepo, smtpServ services.SmtpServ, userRepo repositories.UserRepo, userVerifyRepo repositories.UserVerifyRepo, jwtKey string) *OtpServImpl {
	return &OtpServImpl{Db: db, Validator: validator, OtpRepo: otpRepo, SmtpServ: smtpServ, UserRepo: userRepo, UserVerifyRepo: userVerifyRepo, JwtKey: jwtKey}
}

func (serv *OtpServImpl) SendEmailVerification(request otp.SendRequest) error {
	errValidator := helpers.ErrValidator(request, serv.Validator)
	if errValidator != nil {
		return errValidator
	}

	// Begin Transaction
	tx := serv.Db.Begin()
	if tx.Error != nil {
		log.Printf("Error starting transaction: %v", tx.Error)
		return fmt.Errorf("failed to send OTP, please try again later")
	}

	// Auto Rollback
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Rollback due to panic: %v", r)
		} else if tx.Error != nil {
			tx.Rollback()
			log.Printf("Transaction rolled back: %v", tx.Error)
		}
	}()

	// Generate
	model := otp.SendRequestToDomain(request)
	code, err := serv.generateNumericOTP(6)
	if err != nil {
		return fmt.Errorf("failed to generate OTP")
	}
	model.Code = code

	// Call repo
	if err := serv.OtpRepo.SendOtp(tx, model); err != nil {
		errorMessage := err.Error()

		if errorMessage == "email not exist" {
			return err
		}

		log.Printf("[OtpRepo.SendOtp] error: %v", err)
		return fmt.Errorf("failed to send OTP, please try again later")
	}

	// Call Smtp
	if err := serv.SmtpServ.SendEmailOtp(model.Email, model.Code); err != nil {
		log.Printf("[SmtpServ.SendEmailOtp] error: %v", err)
		return fmt.Errorf("failed to send OTP, please try again later")
	}

	// Commit
	if err := tx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return fmt.Errorf("OTP sent but failed to finalize, please contact support")
	}

	return nil
}

func (serv *OtpServImpl) SendPasswordReset(request otp.SendRequest) error {
	errValidator := helpers.ErrValidator(request, serv.Validator)
	if errValidator != nil {
		return errValidator
	}

	// Begin Transaction
	tx := serv.Db.Begin()
	if tx.Error != nil {
		log.Printf("Error starting transaction: %v", tx.Error)
		return fmt.Errorf("failed to send OTP, please try again later")
	}

	// Auto Rollback
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Rollback due to panic: %v", r)
		} else if tx.Error != nil {
			tx.Rollback()
			log.Printf("Transaction rolled back: %v", tx.Error)
		}
	}()

	// Generate
	model := otp.SendRequestToDomain(request)
	code, err := serv.generateNumericOTP(6)
	if err != nil {
		return fmt.Errorf("failed to generate OTP")
	}
	model.Code = code
	model.Purpose = "password_reset"

	// Call repo
	if err := serv.OtpRepo.SendOtp(tx, model); err != nil {
		errorMessage := err.Error()

		if errorMessage == "email not exist" {
			return err
		}

		log.Printf("[OtpRepo.SendOtp] error: %v", err)
		return fmt.Errorf("failed to send OTP, please try again later")
	}

	// Call Smtp
	if err := serv.SmtpServ.SendEmailOtp(model.Email, model.Code); err != nil {
		log.Printf("[SmtpServ.SendPasswordReset] error: %v", err)
		return fmt.Errorf("failed to send password reset, please try again later")
	}

	// Commit
	if err := tx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return fmt.Errorf("OTP sent but failed to finalize, please contact support")
	}

	return nil
}

func (serv *OtpServImpl) CheckOtp(request otp.CheckRequest) error {
	errValidator := helpers.ErrValidator(request, serv.Validator)
	if errValidator != nil {
		return errValidator
	}

	model := otp.CheckRequestToDomains(request)
	model.Purpose = "email_verification"

	// Call repo
	result, err := serv.OtpRepo.CheckOtp(serv.Db, model)
	if err != nil {
		log.Printf("[OtpRepo.CheckOtp] error: %v", err)
		return fmt.Errorf("failed to check OTP, please try again later")
	}
	if !result {
		log.Printf("Check otp result is false")
		return fmt.Errorf("failed to check OTP, please try again later")
	}

	// Call Verify Repo
	verifyErr := serv.UserVerifyRepo.UpdateToVerified(serv.Db, domains.UserVerified{
		Email: request.Email,
	})
	if verifyErr != nil {
		log.Printf("[UserVerifyRepo.UpdateToVerified] error: %v", err)
		return fmt.Errorf("failed to check OTP, please try again later")
	}

	return nil
}

func (serv *OtpServImpl) CheckOtpPassword(request otp.CheckRequest) (*string, error) {
	errValidator := helpers.ErrValidator(request, serv.Validator)
	if errValidator != nil {
		return nil, errValidator
	}

	model := otp.CheckRequestToDomains(request)
	model.Purpose = "password_reset"

	// Call repo
	result, err := serv.OtpRepo.CheckOtp(serv.Db, model)
	if err != nil {
		log.Printf("[OtpRepo.CheckOtp] error: %v", err)
		return nil, fmt.Errorf("failed to check OTP, please try again later")
	}
	if !result {
		log.Printf("Check otp result is false")
		return nil, fmt.Errorf("failed to check OTP, please try again later")
	}

	// Find by email
	userByEmail, errEmail := serv.UserRepo.FindByEmail(serv.Db, model.Email)
	if userByEmail == nil || errEmail != nil {
		log.Printf("[OtpRepo.CheckOtp] error: %v", err)
		return nil, fmt.Errorf("failed to check OTP, please try again later")
	}

	// Create JWT Token
	duration := time.Minute * 5
	jwtModel := user2.ToSingleResource(userByEmail)
	jwt, errJwt := helpers.GenerateJWT(serv.JwtKey, duration, &jwtModel)
	if errJwt != nil {
		log.Printf("[GenerateJWT] error: %v", errJwt)
		return nil, fmt.Errorf("failed to check OTP, please try again later")
	}

	return &jwt, nil
}

func (serv *OtpServImpl) generateNumericOTP(length int) (string, error) {
	const digits = "0123456789"
	result := make([]byte, length)
	maxCode := big.NewInt(int64(len(digits)))

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, maxCode)
		if err != nil {
			return "", err
		}
		result[i] = digits[n.Int64()]
	}

	return string(result), nil
}
