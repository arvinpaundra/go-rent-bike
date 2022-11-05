package usecase

import (
	"time"

	"github.com/arvinpaundra/go-rent-bike/helper"
	"github.com/google/uuid"

	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"github.com/arvinpaundra/go-rent-bike/internal/repository/gormdb"
)

type UserUsecase interface {
	RegisterUser(userDTO dto.UserDTO) error
	LoginUser(email string, password string) (*model.User, error)
	FindAllUsers() (*[]model.User, error)
	FindByIdUser(userId string) (*model.User, error)
	FindAllUserHistories(userId string) (*[]model.History, error)
	UpdateUser(userId string, userDTO dto.UserDTO) error
	DeleteUser(userId string) (uint, error)
}

type userUsecase struct {
	userRepository    repository.UserRepository
	historyRepository gormdb.HistoryRepository
}

func (u userUsecase) RegisterUser(userDTO dto.UserDTO) error {
	hashedPassword, _ := helper.HashPassword(userDTO.Password)

	user := model.User{
		ID:        uuid.NewString(),
		Fullname:  userDTO.Fullname,
		Phone:     userDTO.Phone,
		Address:   userDTO.Address,
		Role:      userDTO.Role,
		Email:     userDTO.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := u.userRepository.Create(user)

	if err != nil {
		return err
	}

	return nil
}

func (u userUsecase) LoginUser(email string, password string) (*model.User, error) {
	user, err := u.userRepository.FindByEmailAndPassword(email, password)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userUsecase) FindAllUsers() (*[]model.User, error) {
	users, err := u.userRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u userUsecase) FindByIdUser(userId string) (*model.User, error) {
	user, err := u.userRepository.FindById(userId)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userUsecase) FindAllUserHistories(userId string) (*[]model.History, error) {
	if _, err := u.userRepository.FindById(userId); err != nil {
		return nil, err
	}

	histories, err := u.historyRepository.FindAll(userId)

	if err != nil {
		return nil, err
	}

	return histories, nil
}

func (u userUsecase) UpdateUser(userId string, userDTO dto.UserDTO) error {
	user := model.User{
		Address:   userDTO.Address,
		Fullname:  userDTO.Fullname,
		UpdatedAt: time.Time{},
	}

	err := u.userRepository.Update(userId, user)

	if err != nil {
		return err
	}

	return nil
}

func (u userUsecase) DeleteUser(userId string) (uint, error) {
	rowAffected, err := u.userRepository.Delete(userId)

	if err != nil {
		return rowAffected, err
	}

	return rowAffected, nil
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return userUsecase{userRepository: userRepo}
}
