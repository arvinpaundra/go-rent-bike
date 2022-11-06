package usecase

import (
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	repomock "github.com/arvinpaundra/go-rent-bike/internal/repository/gormdb/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var bikeRepository = repomock.BikeRepositoryMock{Mock: mock.Mock{}}
var bikeUsecaseTest = NewBikeUsecase(&bikeRepository)

// TODO create new bike test
//func TestBikeUsecase_CreateNewBike(t *testing.T) {
//	bikeDTO := dto.BikeDTO{
//		RenterId:     "aefde097-3145-4961-9eed-9e916b9def36",
//		CategoryId:   "63390a70-bc40-4f95-9f72-4f054c437949",
//		Name:         "Sample BMX Bike",
//		PricePerHour: 15000,
//		Condition:    "Good",
//		Description:  "This is a description section",
//		IsAvailable:  "1",
//	}
//
//	bike := model.Bike{
//		RenterId:     "aefde097-3145-4961-9eed-9e916b9def36",
//		CategoryId:   "63390a70-bc40-4f95-9f72-4f054c437949",
//		Name:         "Sample BMX Bike",
//		PricePerHour: 15000,
//		Condition:    "Good",
//		Description:  "This is a description section",
//		IsAvailable:  "1",
//	}
//
//	bikeRepository.Mock.On("Create", bike).Return(nil)
//
//	err := bikeUsecaseTest.CreateNewBike(bikeDTO)
//
//	assert.Nil(t, err)
//}

// TODO create new bike review test

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
		{
			ID:           "79770d28-69d0-4c6c-95f7-505e86c880ba",
			RenterId:     "abd85a80-200b-4c76-9376-1f968e3e7393",
			CategoryId:   "07f332fc-4a49-40a1-a7a8-72efeb2d9b8b",
			Name:         "Sample Mountain Bike",
			PricePerHour: 20000,
			Condition:    "Great",
			Description:  "Description of Mountain Bike",
			IsAvailable:  "0",
		},
	}

	bikeRepository.Mock.On("FindAll", "").Return(bikes, nil)

	results, err := bikeUsecaseTest.FindAllBikes("")

	assert.Nil(t, err)
	assert.NotNil(t, results)
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
	assert.Equal(t, bikeId, result.ID)
}

// TODO find bikes by renter test

// TODO find bikes by category test

//func TestBikeUsecase_FindBikesByRenter(t *testing.T) {
//	var err error
//	renterId := "127fe83c-21b2-4d2e-ab98-369b88d4eec5"
//
//renter := &model.Renter{
//	ID:          "127fe83c-21b2-4d2e-ab98-369b88d4eec5",
//	UserId:      "d4e0632c-1b53-4983-99b0-f2de8f89d3aa",
//	RentName:    "Rental Sepeda Sejahtera",
//	RentAddress: "Jl Kalisapu",
//	Description: "Description of rental bike",
//}
//
//bikeRepository.RenterRepository.Mock.On("FindById", renterId).Return(renter, nil)

//	bikes := &[]model.Bike{
//		{
//			ID:           "aefde097-3145-4961-9eed-9e916b9def36",
//			RenterId:     "ffad8203-b32d-46dd-b488-a700ad61dac7",
//			CategoryId:   "463c249c-300e-4138-8cfb-8a7fa5042da0",
//			Name:         "Sample BMX Bike",
//			PricePerHour: 15000,
//			Condition:    "Good",
//			Description:  "This is a description section",
//			IsAvailable:  "1",
//		},
//		{
//			ID:           "79770d28-69d0-4c6c-95f7-505e86c880ba",
//			RenterId:     "abd85a80-200b-4c76-9376-1f968e3e7393",
//			CategoryId:   "07f332fc-4a49-40a1-a7a8-72efeb2d9b8b",
//			Name:         "Sample Mountain Bike",
//			PricePerHour: 20000,
//			Condition:    "Great",
//			Description:  "Description of Mountain Bike",
//			IsAvailable:  "0",
//		},
//	}
//
//	bikeRepository.Mock.On("FindByIdRenter", renterId).Return(bikes, nil)
//
//	results := make(map[string]interface{})
//	results, err = bikeUsecaseTest.FindBikesByRenter(renterId)
//
//	assert.Nil(t, err)
//	assert.NotNil(t, results)
//}

//func TestBikeUsecase_FindBikesByCategory(t *testing.T) {
//	categoryId := "127fe83c-21b2-4d2e-ab98-369b88d4eec5"
//
//	bikes := &[]model.Bike{
//		{
//			ID:           "aefde097-3145-4961-9eed-9e916b9def36",
//			RenterId:     "ffad8203-b32d-46dd-b488-a700ad61dac7",
//			CategoryId:   "463c249c-300e-4138-8cfb-8a7fa5042da0",
//			Name:         "Sample BMX Bike",
//			PricePerHour: 15000,
//			Condition:    "Good",
//			Description:  "This is a description section",
//			IsAvailable:  "1",
//		},
//		{
//			ID:           "79770d28-69d0-4c6c-95f7-505e86c880ba",
//			RenterId:     "abd85a80-200b-4c76-9376-1f968e3e7393",
//			CategoryId:   "07f332fc-4a49-40a1-a7a8-72efeb2d9b8b",
//			Name:         "Sample Mountain Bike",
//			PricePerHour: 20000,
//			Condition:    "Great",
//			Description:  "Description of Mountain Bike",
//			IsAvailable:  "0",
//		},
//	}
//
//	bikeRepository.Mock.On("FindByIdCategory", categoryId).Return(bikes, nil)
//
//	results, err := bikeUsecaseTest.FindBikesByCategory(categoryId)
//
//	assert.Nil(t, err)
//	assert.NotNil(t, results)
//}

func TestBikeUsecase_DeleteBike(t *testing.T) {
	bikeId := "07f332fc-4a49-40a1-a7a8-72efeb2d9b8b"

	bikeRepository.Mock.On("Delete", bikeId).Return(nil)

	err := bikeUsecaseTest.DeleteBike(bikeId)

	assert.Nil(t, err)
}
