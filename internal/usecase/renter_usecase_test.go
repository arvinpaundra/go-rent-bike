package usecase

import (
	"testing"
	"time"

	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/pkg"
	"github.com/stretchr/testify/mock"

	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/stretchr/testify/assert"
)

var renterUsecaseTest = NewRenterUsecase(
	&pkg.RenterRepository,
	&pkg.UserRepository,
	&pkg.ReportRepository,
)

func TestRenterUsecase_CreateRenter(t *testing.T) {
	userId := "e694b986-cf9b-4b33-9147-3838e9014662"

	user := &model.User{
		ID:        "e694b986-cf9b-4b33-9147-3838e9014662",
		Fullname:  "Arvin",
		Phone:     "085",
		Address:   "Jl Rinjani",
		Role:      "customer",
		Email:     "arvin@mail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	pkg.UserRepository.Mock.On("FindById", userId).Return(user, nil)

	pkg.RenterRepository.Mock.On("Create", mock.Anything).Return(nil)

	renterDTO := dto.RenterDTO{
		UserId:      "e694b986-cf9b-4b33-9147-3838e9014662",
		RentName:    "Twins' Brother Bike Rental",
		RentAddress: "Jl Morioh",
		Description: "Full with description texts",
	}

	err := renterUsecaseTest.CreateRenter(renterDTO)

	assert.Nil(t, err)
}

func TestRenterUsecase_CreateReportRenter(t *testing.T) {
	renterId := "aefde097-3145-4961-9eed-9e916b9def36"

	renter := &model.Renter{
		ID:          "aefde097-3145-4961-9eed-9e916b9def36",
		UserId:      "b2a4d5da-198f-4742-adb1-6700957f9510",
		RentName:    "Twins' Brother Bike Rental",
		RentAddress: "Jl Morioh",
		Description: "Full with description texts",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	pkg.RenterRepository.Mock.On("FindById", renterId).Return(renter, nil)

	pkg.ReportRepository.Mock.On("Create", mock.Anything).Return(nil)

	reportDTO := dto.ReportDTO{
		UserId:     "b2a4d5da-198f-4742-adb1-6700957f9510",
		TitleIssue: "Dasar Penipu",
		BodyIssue:  "Dasar penipu, semoga masuk nereka kamu!!!!",
	}

	err := renterUsecaseTest.CreateReportRenter(renterId, reportDTO)

	assert.Nil(t, err)
}

func TestRenterUsecase_FindAllRenters(t *testing.T) {
	renters := &[]model.Renter{
		{
			RentName:    "Abadi Sejahtera",
			RentAddress: "Jl Ketapang",
			Description: "This is a description section",
		},
	}

	pkg.RenterRepository.Mock.On("FindAll", "").Return(renters, nil)

	results, err := renterUsecaseTest.FindAllRenters("")

	assert.Nil(t, err)
	assert.NotNil(t, results)

	assert.Equal(t, (*renters)[0].ID, (*results)[0].ID)
	assert.Equal(t, (*renters)[0].UserId, (*results)[0].UserId)
	assert.Equal(t, (*renters)[0].RentName, (*results)[0].RentName)
	assert.Equal(t, (*renters)[0].RentAddress, (*results)[0].RentAddress)
	assert.Equal(t, (*renters)[0].Description, (*results)[0].Description)
}

func TestRenterUsecase_FindByIdRenter(t *testing.T) {
	renterId := "aefde097-3145-4961-9eed-9e916b9def36"

	renter := &model.Renter{
		ID:          "aefde097-3145-4961-9eed-9e916b9def36",
		UserId:      "b2a4d5da-198f-4742-adb1-6700957f9510",
		RentName:    "Twins' Brother Bike Rental",
		RentAddress: "Jl Morioh",
		Description: "Full with description texts",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	pkg.RenterRepository.Mock.On("FindById", renterId).Return(renter, nil)

	result, err := renterUsecaseTest.FindByIdRenter(renterId)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, renter.ID, result.ID)
	assert.Equal(t, renter.UserId, result.UserId)
	assert.Equal(t, renter.RentName, result.RentName)
	assert.Equal(t, renter.RentAddress, result.RentAddress)
	assert.Equal(t, renter.Description, result.Description)
}

func TestRenterUsecase_FindAllRenterReports(t *testing.T) {
	renterId := "aefde097-3145-4961-9eed-9e916b9def36"

	renter := &model.Renter{
		ID:          "aefde097-3145-4961-9eed-9e916b9def36",
		UserId:      "b2a4d5da-198f-4742-adb1-6700957f9510",
		RentName:    "Twins' Brother Bike Rental",
		RentAddress: "Jl Morioh",
		Description: "Full with description texts",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	pkg.RenterRepository.Mock.On("FindById", renterId).Return(renter, nil)

	reports := &[]model.Report{
		{
			ID:         "0c535949-0b0d-447b-b03b-d394789acf9e",
			RenterId:   "aefde097-3145-4961-9eed-9e916b9def36",
			UserId:     "b2a4d5da-198f-4742-adb1-6700957f9510",
			TitleIssue: "Dasar Penipu",
			BodyIssue:  "Dasar penipu, semoga masuk nereka kamu!!!!",
			User:       nil,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	pkg.ReportRepository.Mock.On("FindAll", renterId).Return(reports, nil)

	results, err := renterUsecaseTest.FindAllRenterReports(renterId)

	assert.Nil(t, err)
	assert.NotNil(t, results)

	assert.Equal(t, (*reports)[0].ID, (*results)[0].ID)
	assert.Equal(t, (*reports)[0].RenterId, (*results)[0].RenterId)
	assert.Equal(t, (*reports)[0].UserId, (*results)[0].UserId)
	assert.Equal(t, (*reports)[0].BodyIssue, (*results)[0].BodyIssue)
	assert.Equal(t, (*reports)[0].BodyIssue, (*results)[0].BodyIssue)
}

func TestRenterUsecase_UpdateRenter(t *testing.T) {
	renterId := "aefde097-3145-4961-9eed-9e916b9def36"

	renter := &model.Renter{
		ID:          "aefde097-3145-4961-9eed-9e916b9def36",
		UserId:      "b2a4d5da-198f-4742-adb1-6700957f9510",
		RentName:    "Twins' Brother Bike Rental",
		RentAddress: "Jl Morioh",
		Description: "Full with description texts",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	pkg.RenterRepository.Mock.On("FindById", renterId).Return(renter, nil)

	pkg.RenterRepository.Mock.On("Update", renterId, mock.Anything).Return(nil)

	renterDTO := dto.RenterDTO{
		UserId:      "e694b986-cf9b-4b33-9147-3838e9014662",
		RentName:    "Twins' Brother Bike Rental",
		RentAddress: "Jl Morioh",
		Description: "Full with description texts",
	}

	err := renterUsecaseTest.UpdateRenter(renterId, renterDTO)

	assert.Nil(t, err)
}

func TestRenterUsecase_DeleteRenter(t *testing.T) {
	renterId := "aefde097-3145-4961-9eed-9e916b9def36"

	renter := &model.Renter{
		ID:          "aefde097-3145-4961-9eed-9e916b9def36",
		RentName:    "Abadi Sejahtera",
		RentAddress: "Jl Ketapang",
		Description: "This is a description section",
	}

	pkg.RenterRepository.Mock.On("FindById", renterId).Return(renter, nil)

	pkg.RenterRepository.Mock.On("Delete", renterId).Return(nil)

	err := renterUsecaseTest.DeleteRenter(renterId)

	assert.Nil(t, err)
}
