package repomock

import (
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/stretchr/testify/mock"
)

type PaymentRepositoryMock struct {
	Mock mock.Mock
}

func (p *PaymentRepositoryMock) Create(paymentUC model.Payment) error {
	ret := p.Mock.Called(paymentUC)

	return ret.Error(0)
}

func (p *PaymentRepositoryMock) FindById(paymentId string) (*model.Payment, error) {
	ret := p.Mock.Called(paymentId)

	return ret.Get(0).(*model.Payment), ret.Error(1)
}

func (p *PaymentRepositoryMock) Update(paymentId string, paymentUC model.Payment) error {
	ret := p.Mock.Called(paymentId, paymentUC)

	return ret.Error(0)
}
