package repomock

import (
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/stretchr/testify/mock"
)

type BikeRepositoryMock struct {
	Mock mock.Mock
}

func (r *BikeRepositoryMock) Create(bikeUC model.Bike) error {
	ret := r.Mock.Called(bikeUC)

	return ret.Error(0)
}

func (r *BikeRepositoryMock) FindAll(bikeName string) (*[]model.Bike, error) {
	ret := r.Mock.Called(bikeName)

	return ret.Get(0).(*[]model.Bike), ret.Error(1)
}

func (r *BikeRepositoryMock) FindById(bikeId string) (*model.Bike, error) {
	ret := r.Mock.Called(bikeId)

	return ret.Get(0).(*model.Bike), ret.Error(1)
}

func (r *BikeRepositoryMock) FindByIdRenter(renterId string) (*[]model.Bike, error) {
	ret := r.Mock.Called(renterId)

	return ret.Get(0).(*[]model.Bike), ret.Error(1)
}

func (r *BikeRepositoryMock) FindByIdCategory(categoryId string) (*[]model.Bike, error) {
	ret := r.Mock.Called(categoryId)

	return ret.Get(0).(*[]model.Bike), ret.Error(1)
}

func (r *BikeRepositoryMock) Update(bikeId string, bikeUC model.Bike) error {
	ret := r.Mock.Called(bikeId, bikeUC)

	return ret.Error(0)
}

func (r *BikeRepositoryMock) Delete(bikeId string) error {
	ret := r.Mock.Called(bikeId)

	return ret.Error(0)
}
