package usecase

import (
	"time"

	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"github.com/arvinpaundra/go-rent-bike/internal/repository/gormdb"
	"github.com/google/uuid"
)

type RenterUsecase interface {
	CreateRenter(renterDTO dto.RenterDTO) error
	FindAllRenters(rentName string) (*[]model.Renter, error)
	FindByIdRenter(renterId string) (*model.Renter, error)
	UpdateRenter(renterId string, renterDTO dto.RenterDTO) error
	DeleteRenter(renterId string) (uint, error)
}

type renterUsecase struct {
	renterRepository repository.RenterRepository
	userRepository   gormdb.UserRepository
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

func (r renterUsecase) UpdateRenter(renterId string, renterDTO dto.RenterDTO) error {
	renter := model.Renter{
		RentName:    renterDTO.RentName,
		RentAddress: renterDTO.RentAddress,
		Description: renterDTO.Description,
	}

	err := r.renterRepository.Update(renterId, renter)

	if err != nil {
		return err
	}

	return nil
}

func (r renterUsecase) DeleteRenter(renterId string) (uint, error) {
	rowAffected, err := r.renterRepository.Delete(renterId)

	if err != nil {
		return rowAffected, err
	}

	return rowAffected, nil
}

func NewRenterUsecase(renterRepo repository.RenterRepository) RenterUsecase {
	return renterUsecase{renterRepository: renterRepo}
}
