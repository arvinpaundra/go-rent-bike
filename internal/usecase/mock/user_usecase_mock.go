package usecasemock

import (
	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/stretchr/testify/mock"
)

type UserUsecaseMock struct {
	Mock mock.Mock
}

func (u *UserUsecaseMock) RegisterUser(userDTO dto.UserDTO) error {
	ret := u.Mock.Called(userDTO)

	return ret.Error(0)
}

func (u *UserUsecaseMock) LoginUser(email string, password string) (string, error) {
	ret := u.Mock.Called(email, password)

	return ret.Get(0).(string), ret.Error(1)
}

func (u *UserUsecaseMock) FindAllUsers() (*[]model.User, error) {
	ret := u.Mock.Called()

	return ret.Get(0).(*[]model.User), ret.Error(1)
}

func (u *UserUsecaseMock) FindByIdUser(userId string) (*model.User, error) {
	ret := u.Mock.Called(userId)

	return ret.Get(0).(*model.User), ret.Error(1)
}

func (u *UserUsecaseMock) FindAllUserHistories(userId string) (*[]model.History, error) {
	ret := u.Mock.Called(userId)

	return ret.Get(0).(*[]model.History), ret.Error(1)
}

func (u *UserUsecaseMock) FindAllOrdersUser(userId string) (*[]model.Order, error) {
	ret := u.Mock.Called(userId)

	return ret.Get(0).(*[]model.Order), ret.Error(1)
}

func (u *UserUsecaseMock) FindByIdOrderUser(orderId string) (*model.Order, error) {
	ret := u.Mock.Called(orderId)

	return ret.Get(0).(*model.Order), ret.Error(1)
}

func (u *UserUsecaseMock) UpdateUser(userId string, userDTO dto.UserDTO) error {
	ret := u.Mock.Called(userId, userDTO)

	return ret.Error(0)
}

func (u *UserUsecaseMock) DeleteUser(userId string) error {
	ret := u.Mock.Called(userId)

	return ret.Error(0)
}
