package usecase

import (
	"time"

	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"github.com/google/uuid"
)

type BikeUsecase interface {
	CreateNewBike(bikeDTO dto.BikeDTO) error
	CreateNewBikeReview(bikeId string, reviewDTO dto.ReviewDTO) error
	FindAllBikes(bikeName string) (*[]model.Bike, error)
	FindByIdBike(bikeId string) (*model.Bike, error)
	FindBikesByRenter(renterId string) (*[]model.Bike, error)
	FindBikesByCategory(categoryId string) (*[]model.Bike, error)
	UpdateBike(bikeId string, bikeDTO dto.BikeDTO) error
	DeleteBike(bikeId string) error
}

type bikeUsecase struct {
	bikeRepository     repository.BikeRepository
	renterRepository   repository.RenterRepository
	categoryRepository repository.CategoryRepository
	userRepository     repository.UserRepository
	reviewRepository   repository.ReviewRepository
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

func (u bikeUsecase) FindBikesByRenter(renterId string) (*[]model.Bike, error) {
	if _, err := u.renterRepository.FindById(renterId); err != nil {
		return nil, err
	}

	bikes, err := u.bikeRepository.FindByIdRenter(renterId)

	if err != nil {
		return nil, err
	}

	return bikes, nil
}

func (u bikeUsecase) FindBikesByCategory(categoryId string) (*[]model.Bike, error) {
	if _, err := u.categoryRepository.FindById(categoryId); err != nil {
		return nil, err
	}

	bikes, err := u.bikeRepository.FindByIdCategory(categoryId)

	if err != nil {
		return nil, err
	}

	return bikes, nil
}

func (u bikeUsecase) UpdateBike(bikeId string, bikeDTO dto.BikeDTO) error {
	var err error
	_, err = u.bikeRepository.FindById(bikeId)

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
		CategoryId:   categoryId,
		Name:         bikeDTO.Name,
		PricePerHour: bikeDTO.PricePerHour,
		Condition:    bikeDTO.Condition,
		Description:  bikeDTO.Description,
		IsAvailable:  bikeDTO.IsAvailable,
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

func (u bikeUsecase) CreateNewBikeReview(bikeId string, reviewDTO dto.ReviewDTO) error {
	if _, err := u.bikeRepository.FindById(bikeId); err != nil {
		return err
	}

	if _, err := u.userRepository.FindById(reviewDTO.UserId); err != nil {
		return err
	}

	review := model.Review{
		ID:          uuid.NewString(),
		BikeId:      bikeId,
		UserId:      reviewDTO.UserId,
		Rating:      reviewDTO.Rating,
		Description: reviewDTO.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := u.reviewRepository.Create(review)

	if err != nil {
		return err
	}

	return nil
}

func NewBikeUsecase(
	bikeRepo repository.BikeRepository,
	renterRepo repository.RenterRepository,
	categoryRepo repository.CategoryRepository,
	userRepo repository.UserRepository,
	reviewRepo repository.ReviewRepository,
) BikeUsecase {
	return bikeUsecase{
		bikeRepository:     bikeRepo,
		renterRepository:   renterRepo,
		categoryRepository: categoryRepo,
		userRepository:     userRepo,
		reviewRepository:   reviewRepo,
	}
}
