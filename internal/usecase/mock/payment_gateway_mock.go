package mock

import "github.com/stretchr/testify/mock"

type PaymentGatewayMock struct {
	Mock mock.Mock
}

func (u *PaymentGatewayMock) MidtransNotification(orderId string) error {
	ret := u.Mock.Called(orderId)

	return ret.Error(0)
}
