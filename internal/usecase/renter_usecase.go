package usecase

import (
	"time"

	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"github.com/google/uuid"
)

type RenterUsecase interface {
	CreateRenter(renterDTO dto.RenterDTO) error
	CreateReportRenter(renterId string, reportDTO dto.ReportDTO) error
	FindAllRenters(rentName string) (*[]model.Renter, error)
	FindByIdRenter(renterId string) (*model.Renter, error)
	FindAllRenterReports(renterId string) (*[]model.Report, error)
	UpdateRenter(renterId string, renterDTO dto.RenterDTO) error
	DeleteRenter(renterId string) error
}

type renterUsecase struct {
	renterRepository repository.RenterRepository
	userRepository   repository.UserRepository
	reportRepository repository.ReportRepository
}

func (r renterUsecase) CreateRenter(renterDTO dto.RenterDTO) error {
	userId := renterDTO.UserId

	if _, err := r.userRepository.FindById(userId); err != nil {
		return err
	}

	renter := model.Renter{
		ID:          uuid.NewString(),
		UserId:      userId,
		RentName:    renterDTO.RentName,
		RentAddress: renterDTO.RentAddress,
		Description: renterDTO.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := r.renterRepository.Create(renter); err != nil {
		return err
	}

	return nil
}

func (r renterUsecase) CreateReportRenter(renterId string, reportDTO dto.ReportDTO) error {
	if _, err := r.renterRepository.FindById(renterId); err != nil {
		return err
	}

	report := model.Report{
		ID:         uuid.NewString(),
		RenterId:   renterId,
		UserId:     reportDTO.UserId,
		TitleIssue: reportDTO.TitleIssue,
		BodyIssue:  reportDTO.BodyIssue,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := r.reportRepository.Create(report)

	if err != nil {
		return err
	}

	return nil
}

func (r renterUsecase) FindAllRenters(rentName string) (*[]model.Renter, error) {
	renters, err := r.renterRepository.FindAll(rentName)

	if err != nil {
		return nil, err
	}

	return renters, nil
}

func (r renterUsecase) FindByIdRenter(renterId string) (*model.Renter, error) {
	renter, err := r.renterRepository.FindById(renterId)

	if err != nil {
		return nil, err
	}

	return renter, nil
}

func (r renterUsecase) FindAllRenterReports(renterId string) (*[]model.Report, error) {
	if _, err := r.renterRepository.FindById(renterId); err != nil {
		return nil, err
	}

	reports, err := r.reportRepository.FindAll(renterId)

	if err != nil {
		return nil, err
	}

	return reports, nil
}

func (r renterUsecase) UpdateRenter(renterId string, renterDTO dto.RenterDTO) error {
	var err error

	var renter *model.Renter
	renter, err = r.renterRepository.FindById(renterId)

	if err != nil {
		return err
	}

	renterUC := model.Renter{
		ID:          renter.ID,
		UserId:      renter.UserId,
		RentName:    renterDTO.RentName,
		RentAddress: renterDTO.RentAddress,
		Description: renterDTO.Description,
		CreatedAt:   renter.CreatedAt,
		UpdatedAt:   time.Now(),
	}

	err = r.renterRepository.Update(renterId, renterUC)

	if err != nil {
		return err
	}

	return nil
}

func (r renterUsecase) DeleteRenter(renterId string) error {
	if _, err := r.renterRepository.FindById(renterId); err != nil {
		return err
	}

	err := r.renterRepository.Delete(renterId)

	if err != nil {
		return err
	}

	return nil
}

func NewRenterUsecase(
	renterRepo repository.RenterRepository,
	userRepo repository.UserRepository,
	reportRepo repository.ReportRepository,
) RenterUsecase {
	return renterUsecase{
		renterRepository: renterRepo,
		userRepository:   userRepo,
		reportRepository: reportRepo,
	}
}
