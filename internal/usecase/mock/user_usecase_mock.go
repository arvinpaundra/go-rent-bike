package mock

import (
	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/stretchr/testify/mock"
)

type UserUsecaseMock struct {
	Mock mock.Mock
}

func (u UserUsecaseMock) RegisterUser(userDTO dto.UserDTO) error {
	ret := u.Mock.Called(userDTO)

	return ret.Error(0)
}

func (u UserUsecaseMock) LoginUser(email string, password string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserUsecaseMock) FindAllUsers() (*[]model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserUsecaseMock) FindByIdUser(userId string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserUsecaseMock) UpdateUser(userId string, userDTO dto.UserDTO) error {
	//TODO implement me
	panic("implement me")
}

func (u UserUsecaseMock) DeleteUser(userId string) (uint, error) {
	//TODO implement me
	panic("implement me")
}
