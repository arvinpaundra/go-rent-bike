package mock

import (
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (u *UserRepositoryMock) Create(user model.User) error {
	ret := u.Mock.Called(user)

	return ret.Error(0)
}

func (u *UserRepositoryMock) FindByEmailAndPassword(email string, password string) (*model.User, error) {
	ret := u.Mock.Called(email, password)

	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}

	return ret.Get(0).(*model.User), ret.Error(1)
}

func (u *UserRepositoryMock) FindAll() (*[]model.User, error) {
	ret := u.Mock.Called()

	return ret.Get(0).(*[]model.User), ret.Error(1)
}

func (u *UserRepositoryMock) FindById(userId string) (*model.User, error) {
	ret := u.Mock.Called(userId)

	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}

	return ret.Get(0).(*model.User), ret.Error(1)
}

func (u *UserRepositoryMock) Update(userId string, userUC model.User) error {
	ret := u.Mock.Called(userId, userUC)

	return ret.Error(0)
}

func (u *UserRepositoryMock) Delete(userId string) (uint, error) {
	ret := u.Mock.Called(userId)

	return ret.Get(0).(uint), ret.Error(1)
}
