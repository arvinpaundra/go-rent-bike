package mock

import (
	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/stretchr/testify/mock"
)

type CategoryUsecaseMock struct {
	Mock mock.Mock
}

func (u *CategoryUsecaseMock) CreateCategory(categoryDTO dto.CategoryDTO) error {
	ret := u.Mock.Called(categoryDTO)

	return ret.Error(0)
}

func (u *CategoryUsecaseMock) FindAllCategories() (*[]model.Category, error) {
	ret := u.Mock.Called()

	return ret.Get(0).(*[]model.Category), ret.Error(1)
}

func (u *CategoryUsecaseMock) FindByIdCategory(categoryId string) (*model.Category, error) {
	ret := u.Mock.Called(categoryId)

	return ret.Get(0).(*model.Category), ret.Error(1)
}

func (u *CategoryUsecaseMock) UpdateCategory(categoryId string, categoryDTO dto.CategoryDTO) error {
	ret := u.Mock.Called(categoryId, categoryDTO)

	return ret.Error(0)
}

func (u *CategoryUsecaseMock) DeleteCategory(categoryId string) (uint, error) {
	ret := u.Mock.Called(categoryId)

	return ret.Get(0).(uint), ret.Error(1)
}
