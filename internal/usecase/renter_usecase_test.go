package usecase

import (
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	repomock "github.com/arvinpaundra/go-rent-bike/internal/repository/gormdb/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var renterRepository = repomock.RenterRepositoryMock{Mock: mock.Mock{}}
var renterUsecaseTest = NewRenterUsecase(&renterRepository)

// TODO create new renter test

// TODO create new renter report

func TestRenterUsecase_FindAllRenters(t *testing.T) {
	renters := &[]model.Renter{
		{
			RentName:    "Abadi Sejahtera",
			RentAddress: "Jl Ketapang",
			Description: "This is a description section",
		},
	}

	renterRepository.Mock.On("FindAll", "").Return(renters, nil)

	results, err := renterUsecaseTest.FindAllRenters("")

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, renters, results)
}

func TestRenterUsecase_FindByIdRenter(t *testing.T) {
	renterId := "aefde097-3145-4961-9eed-9e916b9def36"

	renter := &model.Renter{
		ID:          "aefde097-3145-4961-9eed-9e916b9def36",
		RentName:    "Abadi Sejahtera",
		RentAddress: "Jl Ketapang",
		Description: "This is a description section",
	}

	renterRepository.Mock.On("FindById", renterId).Return(renter, nil)

	result, err := renterUsecaseTest.FindByIdRenter(renterId)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, renter, result)
}

// TODO find all renter reports

// TODO update renter

func TestRenterUsecase_DeleteRenter(t *testing.T) {
	renterId := "aefde097-3145-4961-9eed-9e916b9def36"

	renterRepository.Mock.On("Delete", renterId).Return(uint(1), nil)

	result, err := renterUsecaseTest.DeleteRenter(renterId)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, uint(1), result)
}
