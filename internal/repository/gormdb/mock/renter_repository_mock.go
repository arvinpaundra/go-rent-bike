package repomock

import (
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/stretchr/testify/mock"
)

type RenterRepositoryMock struct {
	Mock mock.Mock
}

func (r *RenterRepositoryMock) Create(renterUC model.Renter) error {
	ret := r.Mock.Called(renterUC)

	return ret.Error(0)
}

func (r *RenterRepositoryMock) FindAll(rentName string) (*[]model.Renter, error) {
	ret := r.Mock.Called(rentName)

	return ret.Get(0).(*[]model.Renter), ret.Error(1)
}

func (r *RenterRepositoryMock) FindById(renterId string) (*model.Renter, error) {
	ret := r.Mock.Called(renterId)

	return ret.Get(0).(*model.Renter), ret.Error(1)
}

func (r *RenterRepositoryMock) FindByIdUser(userId string) (*model.Renter, error) {
	ret := r.Mock.Called(userId)

	return ret.Get(0).(*model.Renter), ret.Error(1)
}

func (r *RenterRepositoryMock) Update(renterId string, renterUC model.Renter) error {
	ret := r.Mock.Called(renterId, renterUC)

	return ret.Error(0)
}

func (r *RenterRepositoryMock) Delete(renterId string) error {
	ret := r.Mock.Called(renterId)

	return ret.Error(0)
}
