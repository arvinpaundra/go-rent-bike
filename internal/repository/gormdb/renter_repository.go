package gormdb

import (
	"errors"
	"time"

	"github.com/arvinpaundra/go-rent-bike/database"
	"github.com/arvinpaundra/go-rent-bike/internal"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"gorm.io/gorm"
)

type RenterRepository struct {
	DB *gorm.DB
}

func (r RenterRepository) Create(renterUC model.Renter) error {
	err := database.DB.Model(&model.Renter{}).Create(&renterUC).Error

	if err != nil {
		return err
	}

	return nil
}

func (r RenterRepository) FindByIdUser(userId string) (*model.Renter, error) {
	renter := &model.Renter{}

	err := database.DB.Model(&model.Renter{}).Where("user_id", userId).Preload("User", "id = ?", userId).Preload("Bikes").Take(&renter).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, internal.ErrRecordNotFound
		}

		return nil, err
	}

	return renter, nil
}

func (r RenterRepository) FindAll(rentName string) (*[]model.Renter, error) {
	renters := &[]model.Renter{}

	err := database.DB.Model(&model.Renter{}).Preload("User").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Omit("password")
	}).Where("rent_name LIKE ?", "%"+rentName+"%").Find(&renters).Error

	if err != nil {
		return nil, err
	}

	return renters, nil
}

func (r RenterRepository) FindById(renterId string) (*model.Renter, error) {
	renter := &model.Renter{}

	err := database.DB.Model(&model.Renter{}).Where("id = ?", renterId).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Omit("password")
	}).Take(&renter).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, internal.ErrRecordNotFound
		}

		return nil, err
	}

	return renter, nil
}

func (r RenterRepository) Update(renterId string, renterUC model.Renter) error {
	renter := model.Renter{}

	msg := database.DB.Model(&model.Renter{}).Where("id = ?", renterId).Take(&renter).Error

	if msg != nil {
		if errors.Is(msg, gorm.ErrRecordNotFound) {
			return internal.ErrRecordNotFound
		}

		return msg
	}

	updatedRenter := model.Renter{
		ID:          renterId,
		UserId:      renter.UserId,
		RentName:    renterUC.RentName,
		RentAddress: renterUC.RentAddress,
		Description: renterUC.Description,
		CreatedAt:   renter.CreatedAt,
		UpdatedAt:   time.Now(),
	}

	err := database.DB.Model(&model.Renter{}).Where("id = ?", renterId).Save(&updatedRenter).Error

	if err != nil {
		return err
	}

	return nil
}

func (r RenterRepository) Delete(renterId string) (uint, error) {
	renter := model.Renter{}

	msg := database.DB.Model(&model.Renter{}).Where("id = ?", renterId).Take(&renter).Error

	if msg != nil {
		if errors.Is(msg, gorm.ErrRecordNotFound) {
			return 0, internal.ErrRecordNotFound
		}
	}

	err := database.DB.Model(&model.Renter{}).Where("id = ?", renterId).Delete(&model.Renter{}).Error

	if err != nil {
		return 0, err
	}

	return 1, nil
}

func NewRenterRepositoryGorm(db *gorm.DB) repository.RenterRepository {
	return RenterRepository{db}
}
