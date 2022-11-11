package repomock

import (
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/stretchr/testify/mock"
)

type ReportRepositoryMock struct {
	Mock mock.Mock
}

func (r *ReportRepositoryMock) Create(reportUC model.Report) error {
	ret := r.Mock.Called(reportUC)

	return ret.Error(0)
}

func (r *ReportRepositoryMock) FindAll(renterId string) (*[]model.Report, error) {
	ret := r.Mock.Called(renterId)

	return ret.Get(0).(*[]model.Report), ret.Error(1)
}
