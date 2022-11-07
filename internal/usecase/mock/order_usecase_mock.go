package mock

import (
	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/stretchr/testify/mock"
)

type OrderUsecaseMock struct {
	Mock mock.Mock
}

func (u *OrderUsecaseMock) CreateOrder(orderDTO dto.OrderDTO) (map[string]interface{}, error) {
	ret := u.Mock.Called(orderDTO)

	return ret.Get(0).(map[string]interface{}), ret.Error(1)
}

func (u *OrderUsecaseMock) UpdateRentStatus(orderId string) error {
	ret := u.Mock.Called(orderId)

	return ret.Error(0)
}
