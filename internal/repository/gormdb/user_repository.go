package gormdb

import (
	"errors"
	"github.com/arvinpaundra/go-rent-bike/pkg"

	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r UserRepository) Create(userUC model.User) error {
	err := r.DB.Model(&model.User{}).Create(&userUC).Error

	if err != nil {
		return err
	}

	return nil
}

func (r UserRepository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}

	err := r.DB.Model(&model.User{}).Where("email = ?", email).Take(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, pkg.ErrRecordNotFound
	}

	return user, nil
}

func (r UserRepository) FindAll() (*[]model.User, error) {
	users := &[]model.User{}

	err := r.DB.Model(&model.User{}).Omit("password").Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r UserRepository) FindById(userId string) (*model.User, error) {
	user := &model.User{}

	err := r.DB.Model(&model.User{}).Where("id = ?", userId).Omit("password").Take(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrRecordNotFound
		}

		return nil, err
	}

	return user, nil
}

func (r UserRepository) Update(userId string, userUC model.User) error {
	err := r.DB.Model(&model.User{}).Where("id = ?", userId).Updates(&userUC).Error

	if err != nil {
		return err
	}

	return nil
}

func (r UserRepository) Delete(userId string) error {
	err := r.DB.Model(&model.User{}).Where("id = ?", userId).Delete(&model.User{}).Error

	if err != nil {
		return err
	}

	return nil
}

func NewUserRepositoryGorm(db *gorm.DB) repository.UserRepository {
	return UserRepository{db}
}
