package repomock

import (
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/stretchr/testify/mock"
)

type OrderDetailRepositoryMock struct {
	Mock mock.Mock
}

func (o *OrderDetailRepositoryMock) Create(orderDetailUC []model.OrderDetail) error {
	ret := o.Mock.Called(orderDetailUC)

	return ret.Error(0)
}

func (o *OrderDetailRepositoryMock) FindByIdOrder(orderId string) (*[]model.OrderDetail, error) {
	ret := o.Mock.Called(orderId)

	return ret.Get(0).(*[]model.OrderDetail), ret.Error(1)
}
