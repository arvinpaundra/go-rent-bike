package gormdb

import (
	"errors"
	"github.com/arvinpaundra/go-rent-bike/pkg"

	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"gorm.io/gorm"
)

type BikeRepository struct {
	DB *gorm.DB
}

func (r BikeRepository) Create(bikeUC model.Bike) error {
	err := r.DB.Model(&model.Bike{}).Create(bikeUC).Error

	if err != nil {
		return err
	}

	return nil
}

func (r BikeRepository) FindAll(bikeName string) (*[]model.Bike, error) {
	bikes := &[]model.Bike{}

	err := r.DB.Model(&model.Bike{}).Where("name LIKE ?", "%"+bikeName+"%").Preload("Category").Find(&bikes).Error

	if err != nil {
		return nil, err
	}

	return bikes, nil
}

func (r BikeRepository) FindById(bikeId string) (*model.Bike, error) {
	bike := &model.Bike{}

	err := r.DB.Model(&model.Bike{}).Where("id = ?", bikeId).Preload("Category").Preload("Reviews").Take(&bike).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrRecordNotFound
		}

		return nil, err
	}

	return bike, nil
}

func (r BikeRepository) FindByIdRenter(renterId string) (*[]model.Bike, error) {
	bikes := &[]model.Bike{}

	err := r.DB.Model(&model.Bike{}).Where("renter_id = ?", renterId).Preload("Category").Find(&bikes).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrRecordNotFound
		}

		return nil, err
	}

	return bikes, nil
}

func (r BikeRepository) FindByIdCategory(categoryId string) (*[]model.Bike, error) {
	bikes := &[]model.Bike{}

	err := r.DB.Model(&model.Bike{}).Where("category_id = ?", categoryId).Preload("Category").Find(&bikes).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrRecordNotFound
		}

		return nil, err
	}

	return bikes, nil
}

func (r BikeRepository) Update(bikeId string, bikeUC model.Bike) error {
	err := r.DB.Model(&model.Bike{}).Where("id = ?", bikeId).Updates(&bikeUC).Error

	if err != nil {
		return err
	}

	return nil
}

func (r BikeRepository) Delete(bikeId string) error {
	err := r.DB.Model(&model.Bike{}).Where("id = ?", bikeId).Delete(&model.Bike{}).Error

	if err != nil {
		return err
	}

	return nil
}

func NewBikeRepositoryGorm(db *gorm.DB) repository.BikeRepository {
	return BikeRepository{db}
}
