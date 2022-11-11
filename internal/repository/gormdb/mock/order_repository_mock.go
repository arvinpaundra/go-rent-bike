package repomock

import (
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/stretchr/testify/mock"
)

type OrderRepositoryMock struct {
	Mock mock.Mock
}

func (o *OrderRepositoryMock) Create(orderUC model.Order) error {
	ret := o.Mock.Called(orderUC)

	return ret.Error(0)
}

func (o *OrderRepositoryMock) FindAll(userId string) (*[]model.Order, error) {
	ret := o.Mock.Called(userId)

	return ret.Get(0).(*[]model.Order), ret.Error(1)
}

func (o *OrderRepositoryMock) FindById(orderId string) (*model.Order, error) {
	ret := o.Mock.Called(orderId)

	return ret.Get(0).(*model.Order), ret.Error(1)
}

func (o *OrderRepositoryMock) Update(orderId string, orderUC model.Order) error {
	ret := o.Mock.Called(orderId, orderUC)

	return ret.Error(0)
}
