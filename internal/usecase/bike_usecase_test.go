package usecase

import (
	"testing"
	"time"

	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	repomock "github.com/arvinpaundra/go-rent-bike/internal/repository/gormdb/mock"
	"github.com/arvinpaundra/go-rent-bike/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var bikeRepository = repomock.BikeRepositoryMock{Mock: mock.Mock{}}
var bikeUsecaseTest = NewBikeUsecase(
	&bikeRepository,
	&pkg.RenterRepository,
	&pkg.CategoryRepository,
	&pkg.UserRepository,
	&pkg.ReviewRepository,
)

func TestBikeUsecase_CreateNewBike(t *testing.T) {
	renterId := "aefde097-3145-4961-9eed-9e916b9def36"

	renter := &model.Renter{
		ID:          "aefde097-3145-4961-9eed-9e916b9def36",
		UserId:      "b2a4d5da-198f-4742-adb1-6700957f9510",
		RentName:    "Twins' Brother Bike Rental",
		RentAddress: "Jl Morioh",
		Description: "Full with description texts",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	pkg.RenterRepository.Mock.On("FindById", renterId).Return(renter, nil)

	categoryId := "63390a70-bc40-4f95-9f72-4f054c437949"

	category := &model.Category{
		ID:        "169c38a9-7047-4216-bc2b-869db969a239",
		Name:      "BMX",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	pkg.CategoryRepository.Mock.On("FindById", categoryId).Return(category, nil)

	bikeRepository.Mock.On("Create", mock.Anything).Return(nil)

	bikeDTO := dto.BikeDTO{
		RenterId:     "aefde097-3145-4961-9eed-9e916b9def36",
		CategoryId:   "63390a70-bc40-4f95-9f72-4f054c437949",
		Name:         "Sample BMX Bike",
		PricePerHour: 15000,
		Condition:    "Good",
		Description:  "This is a description section",
		IsAvailable:  "1",
	}

	err := bikeUsecaseTest.CreateNewBike(bikeDTO)

	assert.Nil(t, err)
}

func TestBikeUsecase_CreateNewBikeReview(t *testing.T) {
	bikeId := "aefde097-3145-4961-9eed-9e916b9def36"

	bike := &model.Bike{
		ID:           "aefde097-3145-4961-9eed-9e916b9def36",
		RenterId:     "ffad8203-b32d-46dd-b488-a700ad61dac7",
		CategoryId:   "463c249c-300e-4138-8cfb-8a7fa5042da0",
		Name:         "Sample BMX Bike",
		PricePerHour: 15000,
		Condition:    "Good",
		Description:  "This is a description section",
		IsAvailable:  "1",
	}

	bikeRepository.Mock.On("FindById", bikeId).Return(bike, nil)

	userId := "97ed3f84-768e-481f-b364-58d53f90d5e3"

	user := &model.User{
		ID:        userId,
		Fullname:  "Arvin",
		Phone:     "085",
		Address:   "Jl Rinjani",
		Role:      "customer",
		Email:     "arvin@mail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	pkg.UserRepository.Mock.On("FindById", userId).Return(user, nil)

	pkg.ReviewRepository.Mock.On("Create", mock.Anything).Return(nil)

	reviewDTO := dto.ReviewDTO{
		UserId:      "97ed3f84-768e-481f-b364-58d53f90d5e3",
		Rating:      5,
		Description: "These bikes are over poweerr, sheeshhh....",
	}

	err := bikeUsecaseTest.CreateNewBikeReview(bikeId, reviewDTO)

	assert.Nil(t, err)
}

func TestBikeUsecase_FindAllBikes(t *testing.T) {
	bikes := &[]model.Bike{
		{
			ID:           "aefde097-3145-4961-9eed-9e916b9def36",
			RenterId:     "ffad8203-b32d-46dd-b488-a700ad61dac7",
			CategoryId:   "463c249c-300e-4138-8cfb-8a7fa5042da0",
			Name:         "Sample BMX Bike",
			PricePerHour: 15000,
			Condition:    "Good",
			Description:  "This is a description section",
			IsAvailable:  "1",
		},
	}

	bikeRepository.Mock.On("FindAll", "").Return(bikes, nil)

	results, err := bikeUsecaseTest.FindAllBikes("")

	assert.Nil(t, err)
	assert.NotNil(t, results)

	assert.Equal(t, (*bikes)[0].ID, (*results)[0].ID)
	assert.Equal(t, (*bikes)[0].RenterId, (*results)[0].RenterId)
	assert.Equal(t, (*bikes)[0].CategoryId, (*results)[0].CategoryId)
	assert.Equal(t, (*bikes)[0].Name, (*results)[0].Name)
	assert.Equal(t, (*bikes)[0].PricePerHour, (*results)[0].PricePerHour)
	assert.Equal(t, (*bikes)[0].Condition, (*results)[0].Condition)
	assert.Equal(t, (*bikes)[0].Description, (*results)[0].Description)
	assert.Equal(t, (*bikes)[0].IsAvailable, (*results)[0].IsAvailable)
}

func TestBikeUsecase_FindByIdBike(t *testing.T) {
	bikeId := "07f332fc-4a49-40a1-a7a8-72efeb2d9b8b"

	bike := &model.Bike{
		ID:           "07f332fc-4a49-40a1-a7a8-72efeb2d9b8b",
		RenterId:     "abd85a80-200b-4c76-9376-1f968e3e7393",
		CategoryId:   "79770d28-69d0-4c6c-95f7-505e86c880ba",
		Name:         "Sample BMX Bike",
		PricePerHour: 15000,
		Condition:    "Good",
		Description:  "This is a description section",
		IsAvailable:  "1",
	}

	bikeRepository.Mock.On("FindById", bikeId).Return(bike, nil)

	result, err := bikeUsecaseTest.FindByIdBike(bikeId)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, bike.ID, result.ID)
	assert.Equal(t, bike.RenterId, result.RenterId)
	assert.Equal(t, bike.CategoryId, result.CategoryId)
	assert.Equal(t, bike.Name, result.Name)
	assert.Equal(t, bike.PricePerHour, result.PricePerHour)
	assert.Equal(t, bike.Condition, result.Condition)
	assert.Equal(t, bike.Description, result.Description)
	assert.Equal(t, bike.IsAvailable, result.IsAvailable)
}

func TestBikeUsecase_FindBikesByRenter(t *testing.T) {
	renterId := "127fe83c-21b2-4d2e-ab98-369b88d4eec5"

	renter := &model.Renter{
		ID:          "127fe83c-21b2-4d2e-ab98-369b88d4eec5",
		UserId:      "d4e0632c-1b53-4983-99b0-f2de8f89d3aa",
		RentName:    "Rental Sepeda Sejahtera",
		RentAddress: "Jl Kalisapu",
		Description: "Description of rental bike",
	}

	pkg.RenterRepository.Mock.On("FindById", renterId).Return(renter, nil)

	bikes := &[]model.Bike{
		{
			ID:           "aefde097-3145-4961-9eed-9e916b9def36",
			RenterId:     "ffad8203-b32d-46dd-b488-a700ad61dac7",
			CategoryId:   "463c249c-300e-4138-8cfb-8a7fa5042da0",
			Name:         "Sample BMX Bike",
			PricePerHour: 15000,
			Condition:    "Good",
			Description:  "This is a description section",
			IsAvailable:  "1",
		},
	}

	bikeRepository.Mock.On("FindByIdRenter", renterId).Return(bikes, nil)

	results, err := bikeUsecaseTest.FindBikesByRenter(renterId)

	assert.Nil(t, err)
	assert.NotNil(t, results)

	assert.Equal(t, (*bikes)[0].ID, (*results)[0].ID)
	assert.Equal(t, (*bikes)[0].RenterId, (*results)[0].RenterId)
	assert.Equal(t, (*bikes)[0].CategoryId, (*results)[0].CategoryId)
	assert.Equal(t, (*bikes)[0].Name, (*results)[0].Name)
	assert.Equal(t, (*bikes)[0].PricePerHour, (*results)[0].PricePerHour)
	assert.Equal(t, (*bikes)[0].Condition, (*results)[0].Condition)
	assert.Equal(t, (*bikes)[0].Description, (*results)[0].Description)
	assert.Equal(t, (*bikes)[0].IsAvailable, (*results)[0].IsAvailable)
}

func TestBikeUsecase_FindBikesByCategory(t *testing.T) {
	categoryId := "127fe83c-21b2-4d2e-ab98-369b88d4eec5"

	category := &model.Category{
		ID:        "127fe83c-21b2-4d2e-ab98-369b88d4eec5",
		Name:      "BMX",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	pkg.CategoryRepository.Mock.On("FindById", categoryId).Return(category, nil)

	bikes := &[]model.Bike{
		{
			ID:           "aefde097-3145-4961-9eed-9e916b9def36",
			RenterId:     "ffad8203-b32d-46dd-b488-a700ad61dac7",
			CategoryId:   "463c249c-300e-4138-8cfb-8a7fa5042da0",
			Name:         "Sample BMX Bike",
			PricePerHour: 15000,
			Condition:    "Good",
			Description:  "This is a description section",
			IsAvailable:  "1",
		},
	}

	bikeRepository.Mock.On("FindByIdCategory", categoryId).Return(bikes, nil)

	results, err := bikeUsecaseTest.FindBikesByCategory(categoryId)

	assert.Nil(t, err)
	assert.NotNil(t, results)

	assert.Equal(t, (*bikes)[0].ID, (*results)[0].ID)
	assert.Equal(t, (*bikes)[0].RenterId, (*results)[0].RenterId)
	assert.Equal(t, (*bikes)[0].CategoryId, (*results)[0].CategoryId)
	assert.Equal(t, (*bikes)[0].Name, (*results)[0].Name)
	assert.Equal(t, (*bikes)[0].PricePerHour, (*results)[0].PricePerHour)
	assert.Equal(t, (*bikes)[0].Condition, (*results)[0].Condition)
	assert.Equal(t, (*bikes)[0].Description, (*results)[0].Description)
	assert.Equal(t, (*bikes)[0].IsAvailable, (*results)[0].IsAvailable)
}

func TestBikeUsecase_UpdateBike(t *testing.T) {
	bikeId := "07f332fc-4a49-40a1-a7a8-72efeb2d9b8b"

	bike := &model.Bike{
		ID:           "07f332fc-4a49-40a1-a7a8-72efeb2d9b8b",
		RenterId:     "abd85a80-200b-4c76-9376-1f968e3e7393",
		CategoryId:   "79770d28-69d0-4c6c-95f7-505e86c880ba",
		Name:         "Sample BMX Bike",
		PricePerHour: 15000,
		Condition:    "Good",
		Description:  "This is a description section",
		IsAvailable:  "1",
	}

	bikeRepository.Mock.On("FindById", bikeId).Return(bike, nil)

	renterId := "abd85a80-200b-4c76-9376-1f968e3e7393"

	renter := &model.Renter{
		ID:          "abd85a80-200b-4c76-9376-1f968e3e7393",
		UserId:      "d4e0632c-1b53-4983-99b0-f2de8f89d3aa",
		RentName:    "Rental Sepeda Sejahtera",
		RentAddress: "Jl Kalisapu",
		Description: "Description of rental bike",
	}

	pkg.RenterRepository.Mock.On("FindById", renterId).Return(renter, nil)

	categoryId := "79770d28-69d0-4c6c-95f7-505e86c880ba"

	category := &model.Category{
		ID:        "79770d28-69d0-4c6c-95f7-505e86c880ba",
		Name:      "BMX",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	pkg.CategoryRepository.Mock.On("FindById", categoryId).Return(category, nil)

	bikeRepository.Mock.On("Update", bikeId, mock.Anything).Return(nil)

	bikeDTO := dto.BikeDTO{
		RenterId:     "abd85a80-200b-4c76-9376-1f968e3e7393",
		CategoryId:   "79770d28-69d0-4c6c-95f7-505e86c880ba",
		Name:         "Sample BMX Bike",
		PricePerHour: 15000,
		Condition:    "Good",
		Description:  "This is a description section",
		IsAvailable:  "1",
	}

	err := bikeUsecaseTest.UpdateBike(bikeId, bikeDTO)

	assert.Nil(t, err)
}

func TestBikeUsecase_DeleteBike(t *testing.T) {
	bikeId := "07f332fc-4a49-40a1-a7a8-72efeb2d9b8b"

	bike := &model.Bike{
		ID:           "07f332fc-4a49-40a1-a7a8-72efeb2d9b8b",
		RenterId:     "abd85a80-200b-4c76-9376-1f968e3e7393",
		CategoryId:   "79770d28-69d0-4c6c-95f7-505e86c880ba",
		Name:         "Sample BMX Bike",
		PricePerHour: 15000,
		Condition:    "Good",
		Description:  "This is a description section",
		IsAvailable:  "1",
	}

	bikeRepository.Mock.On("FindById", bikeId).Return(bike, nil)

	bikeRepository.Mock.On("Delete", bikeId).Return(nil)

	err := bikeUsecaseTest.DeleteBike(bikeId)

	assert.Nil(t, err)
}
