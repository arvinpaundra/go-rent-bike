package repo_mock

import (
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (r *UserRepositoryMock) Create(userUC model.User) error {
	ret := r.Mock.Called(userUC)

	return ret.Error(0)
}

func (r *UserRepositoryMock) FindByEmailAndPassword(email string, password string) (*model.User, error) {
	ret := r.Mock.Called(email, password)

	return ret.Get(0).(*model.User), ret.Error(1)
}

func (r *UserRepositoryMock) FindAll() (*[]model.User, error) {
	ret := r.Mock.Called()

	return ret.Get(0).(*[]model.User), ret.Error(1)
}

func (r *UserRepositoryMock) FindById(userId string) (*model.User, error) {
	ret := r.Mock.Called(userId)

	return ret.Get(0).(*model.User), ret.Error(1)
}

func (r *UserRepositoryMock) Update(userId string, userUC model.User) error {
	ret := r.Mock.Called(userId, userUC)

	return ret.Error(0)
}

func (r *UserRepositoryMock) Delete(userId string) (uint, error) {
	ret := r.Mock.Called(userId)

	return ret.Get(0).(uint), ret.Error(1)
}
