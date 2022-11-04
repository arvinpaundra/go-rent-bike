package usecase

import (
	"time"

	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"github.com/arvinpaundra/go-rent-bike/internal/repository/gormdb"
	"github.com/google/uuid"
)

type BikeUsecase interface {
	CreateNewBike(bikeDTO dto.BikeDTO) error
	FindAllBikes(bikeName string) (*[]model.Bike, error)
	FindByIdBike(bikeId string) (*model.Bike, error)
	FindBikesByRenter(renterId string) (map[string]interface{}, error)
	FindBikesByCategory(categoryId string) (map[string]interface{}, error)
	UpdateBike(bikeId string, bikeDTO dto.BikeDTO) error
	DeleteBike(bikeId string) error
}

type bikeUsecase struct {
	bikeRepository     repository.BikeRepository
	renterRepository   gormdb.RenterRepository
	categoryRepository gormdb.CategoryRepository
}

func (u bikeUsecase) CreateNewBike(bikeDTO dto.BikeDTO) error {
	renterId := bikeDTO.RenterId
	categoryId := bikeDTO.CategoryId

	if _, err := u.renterRepository.FindById(renterId); err != nil {
		return err
	}

	if _, err := u.categoryRepository.FindById(categoryId); err != nil {
		return err
	}

	bike := model.Bike{
		ID:           uuid.NewString(),
		RenterId:     renterId,
		CategoryId:   categoryId,
		Name:         bikeDTO.Name,
		PricePerHour: bikeDTO.PricePerHour,
		Condition:    bikeDTO.Condition,
		Description:  bikeDTO.Description,
		IsAvailable:  bikeDTO.IsAvailable,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := u.bikeRepository.Create(bike)

	if err != nil {
		return err
	}

	return nil
}

func (u bikeUsecase) FindAllBikes(bikeName string) (*[]model.Bike, error) {
	bikes, err := u.bikeRepository.FindAll(bikeName)

	if err != nil {
		return nil, err
	}

	return bikes, nil
}

func (u bikeUsecase) FindByIdBike(bikeId string) (*model.Bike, error) {
	bike, err := u.bikeRepository.FindById(bikeId)

	if err != nil {
		return nil, err
	}

	return bike, nil
}

func (u bikeUsecase) FindBikesByRenter(renterId string) (map[string]interface{}, error) {
	var err error

	renter := &model.Renter{}
	renter, err = u.renterRepository.FindById(renterId)

	if err != nil {
		return nil, err
	}

	bikes := &[]model.Bike{}
	bikes, err = u.bikeRepository.FindByIdRenter(renterId)

	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{
		"rental_name": renter.RentName,
		"bikes":       bikes,
	}

	return data, nil
}

func (u bikeUsecase) FindBikesByCategory(categoryId string) (map[string]interface{}, error) {
	var err error

	category := &model.Category{}
	category, err = u.categoryRepository.FindById(categoryId)

	if err != nil {
		return nil, err
	}

	bikes := &[]model.Bike{}
	bikes, err = u.bikeRepository.FindByIdCategory(categoryId)

	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{
		"category": category.Name,
		"bikes":    bikes,
	}

	return data, nil
}

func (u bikeUsecase) UpdateBike(bikeId string, bikeDTO dto.BikeDTO) error {
	var err error
	bike := &model.Bike{}

	bike, err = u.bikeRepository.FindById(bikeId)

	if err != nil {
		return err
	}

	renterId := bikeDTO.RenterId
	if _, err := u.renterRepository.FindById(renterId); err != nil {
		return err
	}

	categoryId := bikeDTO.CategoryId
	if _, err := u.categoryRepository.FindById(categoryId); err != nil {
		return err
	}

	updatedBike := model.Bike{
		ID:           bikeId,
		RenterId:     renterId,
		CategoryId:   categoryId,
		Name:         bikeDTO.Name,
		PricePerHour: bikeDTO.PricePerHour,
		Condition:    bikeDTO.Condition,
		Description:  bikeDTO.Description,
		IsAvailable:  bikeDTO.IsAvailable,
		CreatedAt:    bike.CreatedAt,
		UpdatedAt:    time.Now(),
	}

	err = u.bikeRepository.Update(bikeId, updatedBike)

	if err != nil {
		return err
	}

	return nil
}

func (u bikeUsecase) DeleteBike(bikeId string) error {
	var err error

	_, err = u.bikeRepository.FindById(bikeId)

	if err != nil {
		return err
	}

	err = u.bikeRepository.Delete(bikeId)

	if err != nil {
		return err
	}

	return nil
}

func NewBikeUsecase(bikeRepo repository.BikeRepository) BikeUsecase {
	return bikeUsecase{bikeRepository: bikeRepo}
}
