package repomock

import (
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/stretchr/testify/mock"
)

type HistoryRepositoryMock struct {
	Mock mock.Mock
}

func (h *HistoryRepositoryMock) Create(historyUC model.History) error {
	ret := h.Mock.Called(historyUC)

	return ret.Error(0)
}

func (h *HistoryRepositoryMock) FindAll(userId string) (*[]model.History, error) {
	ret := h.Mock.Called(userId)

	return ret.Get(0).(*[]model.History), ret.Error(1)
}

func (h *HistoryRepositoryMock) FindByIdOrder(orderId string) (*model.History, error) {
	ret := h.Mock.Called(orderId)

	return ret.Get(0).(*model.History), ret.Error(1)
}

func (h *HistoryRepositoryMock) Update(orderId string, historyUC model.History) error {
	ret := h.Mock.Called(orderId, historyUC)

	return ret.Error(0)
}
