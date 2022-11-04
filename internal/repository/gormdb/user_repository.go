package gormdb

import (
	"errors"
	"time"

	"github.com/arvinpaundra/go-rent-bike/database"
	"github.com/arvinpaundra/go-rent-bike/helper"
	"github.com/arvinpaundra/go-rent-bike/internal"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (u UserRepository) Create(userUC model.User) error {
	user := model.User{}

	_ = database.DB.Model(&model.User{}).Where("email = ?", userUC.Email).Take(&user)

	if user.ID != "" {
		return internal.ErrDataAlreadyExist
	}

	err := database.DB.Model(&model.User{}).Create(&userUC).Error

	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) FindByEmailAndPassword(email string, password string) (*model.User, error) {
	user := &model.User{}

	err := database.DB.Model(&model.User{}).Where("email = ?", email).Take(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, internal.ErrRecordNotFound
	}

	ok := helper.ComparePassword(user.Password, password)

	if !ok {
		return nil, internal.ErrRecordNotFound
	}

	return user, nil
}

func (u UserRepository) FindAll() (*[]model.User, error) {
	users := &[]model.User{}

	err := database.DB.Model(&model.User{}).Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u UserRepository) FindById(userId string) (*model.User, error) {
	user := &model.User{}

	err := database.DB.Model(&model.User{}).Where("id = ?", userId).Take(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, internal.ErrRecordNotFound
		}

		return nil, err
	}

	return user, nil
}

func (u UserRepository) Update(userId string, userUC model.User) error {
	user := model.User{}

	if msg := database.DB.Model(&model.User{}).Where("id = ?", userId).Take(&user).Error; msg != nil {
		if errors.Is(msg, gorm.ErrRecordNotFound) {
			return internal.ErrRecordNotFound
		}

		return msg
	}

	updatedUser := model.User{
		ID:        userId,
		Fullname:  userUC.Fullname,
		Phone:     userUC.Phone,
		Address:   userUC.Address,
		Role:      user.Role,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: time.Now(),
	}

	err := database.DB.Model(&model.User{}).Where("id = ?", userId).Save(&updatedUser).Error

	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) Delete(userId string) (uint, error) {
	user := model.User{}

	if msg := database.DB.Model(&model.User{}).Where("id = ?", userId).Take(&user).Error; msg != nil {
		if errors.Is(msg, gorm.ErrRecordNotFound) {
			return 0, internal.ErrRecordNotFound
		}

		return 0, msg
	}

	err := database.DB.Model(&model.User{}).Where("id = ?", userId).Delete(&model.User{}).Error

	if err != nil {
		return 0, err
	}

	return 1, nil
}

func NewUserRepositoryGorm(db *gorm.DB) repository.UserRepository {
	return UserRepository{db}
}
