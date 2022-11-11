package usecasemock

import (
	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/stretchr/testify/mock"
)

type RenterUsecaseMock struct {
	Mock mock.Mock
}

func (r *RenterUsecaseMock) CreateRenter(renterDTO dto.RenterDTO) error {
	ret := r.Mock.Called(renterDTO)

	return ret.Error(0)
}

func (r *RenterUsecaseMock) CreateReportRenter(renterId string, reportDTO dto.ReportDTO) error {
	ret := r.Mock.Called(renterId, reportDTO)

	return ret.Error(0)
}

func (r *RenterUsecaseMock) FindAllRenters(rentName string) (*[]model.Renter, error) {
	ret := r.Mock.Called(rentName)

	return ret.Get(0).(*[]model.Renter), ret.Error(1)
}

func (r *RenterUsecaseMock) FindByIdRenter(renterId string) (*model.Renter, error) {
	ret := r.Mock.Called(renterId)

	return ret.Get(0).(*model.Renter), ret.Error(1)
}

func (r *RenterUsecaseMock) FindAllRenterReports(renterId string) (*[]model.Report, error) {
	ret := r.Mock.Called(renterId)

	return ret.Get(0).(*[]model.Report), ret.Error(1)
}

func (r *RenterUsecaseMock) UpdateRenter(renterId string, renterDTO dto.RenterDTO) error {
	ret := r.Mock.Called(renterId, renterDTO)

	return ret.Error(0)
}

func (r *RenterUsecaseMock) DeleteRenter(renterId string) error {
	ret := r.Mock.Called(renterId)

	return ret.Error(0)
}
