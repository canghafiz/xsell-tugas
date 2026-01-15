package impl

import (
	"be/helpers"
	"be/models/domains"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserRepoImpl struct {
}

func NewUserRepoImpl() *UserRepoImpl {
	return &UserRepoImpl{}
}

func (repo *UserRepoImpl) Create(db *gorm.DB, user domains.Users) error {
	err := db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepoImpl) UpdateToken(db *gorm.DB, user domains.Users) error {
	err := db.
		Model(&domains.Users{}).
		Where("email = ?", user.Email).
		Updates(map[string]interface{}{
			"token":        user.Token,
			"token_expire": time.Now().Add(24 * time.Hour),
		}).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepoImpl) CheckTokenValid(db *gorm.DB, user domains.Users) bool {
	var count int64

	result := db.
		Model(&user).
		Where("email = ? AND token = ?", user.Email, user.Token).
		Count(&count)

	if result.Error != nil {
		return false
	}

	return int(count) > 0
}

func (repo *UserRepoImpl) CheckUserIdValid(db *gorm.DB, userId int) bool {
	var count int64

	result := db.
		Model(&domains.Users{}).
		Where("user_id = ?", userId).
		Count(&count)

	if result.Error != nil {
		return false
	}

	return int(count) > 0
}

func (repo *UserRepoImpl) Update(db *gorm.DB, user domains.Users) (*domains.Users, error) {
	result := db.
		Model(&domains.Users{}).
		Where("user_id = ?", user.UserId).
		Updates(map[string]interface{}{
			"first_name":        user.FirstName,
			"last_name":         user.LastName,
			"photo_profile_url": user.PhotoProfileUrl,
		})

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user with ID %d not found or no changes were made", user.UserId)
	}

	var updatedData domains.Users
	err := db.Where("user_id = ?", user.UserId).First(&updatedData).Error
	if err != nil {
		return nil, err
	}

	return &updatedData, nil
}

func (repo *UserRepoImpl) ChangePassword(db *gorm.DB, user domains.Users) error {
	result := db.
		Model(&domains.Users{}).
		Where("user_id = ?", user.UserId).
		Updates(map[string]interface{}{
			"password": user.Password,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("user with ID %d not found or no changes were made", user.UserId)
	}

	var updatedData domains.Users
	err := db.Where("user_id = ?", user.UserId).First(&updatedData).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepoImpl) FindByEmail(db *gorm.DB, email string) (*domains.Users, error) {
	var user domains.Users
	result := db.
		Model(&user).
		Where("email = ?", email).
		Preload("Verified").
		First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo *UserRepoImpl) CheckPasswordValid(db *gorm.DB, user domains.Users) (bool, error) {
	result, err := repo.FindByEmail(db, user.Email)
	if err != nil {
		return false, err
	}

	return helpers.CheckPassword(result.Password, user.Password), nil
}

func (repo *UserRepoImpl) ResetToken(db *gorm.DB, email string) error {
	result := db.
		Model(&domains.Users{}).
		Where("email = ?", email).
		Updates(map[string]interface{}{
			"token":        nil,
			"token_expire": nil,
		})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *UserRepoImpl) GetByUserId(
	db *gorm.DB,
	userId int,
) (*domains.Users, error) {

	var user domains.Users

	err := db.
		Preload("Verified").
		Where("user_id = ?", userId).
		First(&user).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // user not found
		}
		return nil, err // other database error
	}

	return &user, nil
}
