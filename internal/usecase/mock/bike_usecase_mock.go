package mock

import (
	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/stretchr/testify/mock"
)

type BikeUsecaseMock struct {
	Mock mock.Mock
}

func (u *BikeUsecaseMock) CreateNewBike(bikeDTO dto.BikeDTO) error {
	ret := u.Mock.Called(bikeDTO)

	return ret.Error(0)
}

func (u *BikeUsecaseMock) CreateNewBikeReview(bikeId string, reviewDTO dto.ReviewDTO) error {
	ret := u.Mock.Called(bikeId, reviewDTO)

	return ret.Error(0)
}

func (u *BikeUsecaseMock) FindAllBikes(bikeName string) (*[]model.Bike, error) {
	ret := u.Mock.Called(bikeName)

	return ret.Get(0).(*[]model.Bike), ret.Error(1)
}

func (u *BikeUsecaseMock) FindByIdBike(bikeId string) (*model.Bike, error) {
	ret := u.Mock.Called(bikeId)

	return ret.Get(0).(*model.Bike), ret.Error(1)
}

func (u *BikeUsecaseMock) FindBikesByRenter(renterId string) (*[]model.Bike, error) {
	ret := u.Mock.Called(renterId)

	return ret.Get(0).(*[]model.Bike), ret.Error(1)
}

func (u *BikeUsecaseMock) FindBikesByCategory(categoryId string) (*[]model.Bike, error) {
	ret := u.Mock.Called(categoryId)

	return ret.Get(0).(*[]model.Bike), ret.Error(1)
}

func (u *BikeUsecaseMock) UpdateBike(bikeId string, bikeDTO dto.BikeDTO) error {
	ret := u.Mock.Called(bikeId, bikeDTO)

	return ret.Error(0)
}

func (u *BikeUsecaseMock) DeleteBike(bikeId string) error {
	ret := u.Mock.Called(bikeId)

	return ret.Error(0)
}
