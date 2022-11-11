package gormdb

import (
	"errors"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"github.com/arvinpaundra/go-rent-bike/pkg"
	"gorm.io/gorm"
)

type RenterRepository struct {
	DB *gorm.DB
}

func (r RenterRepository) Create(renterUC model.Renter) error {
	err := r.DB.Model(&model.Renter{}).Create(&renterUC).Error

	if err != nil {
		return err
	}

	return nil
}

func (r RenterRepository) FindByIdUser(userId string) (*model.Renter, error) {
	renter := &model.Renter{}

	err := r.DB.Model(&model.Renter{}).Where("user_id", userId).Preload("User").Preload("Bikes").Take(&renter).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrRecordNotFound
		}

		return nil, err
	}

	return renter, nil
}

func (r RenterRepository) FindAll(rentName string) (*[]model.Renter, error) {
	renters := &[]model.Renter{}

	err := r.DB.Model(&model.Renter{}).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Omit("password")
	}).Where("rent_name LIKE ?", "%"+rentName+"%").Find(&renters).Error

	if err != nil {
		return nil, err
	}

	return renters, nil
}

func (r RenterRepository) FindById(renterId string) (*model.Renter, error) {
	renter := &model.Renter{}

	err := r.DB.Model(&model.Renter{}).Where("id = ?", renterId).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Omit("password")
	}).Take(&renter).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrRecordNotFound
		}

		return nil, err
	}

	return renter, nil
}

func (r RenterRepository) Update(renterId string, renterUC model.Renter) error {
	err := r.DB.Model(&model.Renter{}).Where("id = ?", renterId).Updates(&renterUC).Error

	if err != nil {
		return err
	}

	return nil
}

func (r RenterRepository) Delete(renterId string) error {
	err := r.DB.Model(&model.Renter{}).Where("id = ?", renterId).Delete(&model.Renter{}).Error

	if err != nil {
		return err
	}

	return nil
}

func NewRenterRepositoryGorm(db *gorm.DB) repository.RenterRepository {
	return RenterRepository{db}
}
