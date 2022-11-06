package repo_mock

import (
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (c *CategoryRepositoryMock) Create(categoryUC model.Category) error {
	ret := c.Mock.Called(categoryUC)

	return ret.Error(0)
}

func (c *CategoryRepositoryMock) FindAll() (*[]model.Category, error) {
	ret := c.Mock.Called()

	return ret.Get(0).(*[]model.Category), ret.Error(1)
}

func (c *CategoryRepositoryMock) FindById(categoryId string) (*model.Category, error) {
	ret := c.Mock.Called(categoryId)

	return ret.Get(0).(*model.Category), ret.Error(1)
}

func (c *CategoryRepositoryMock) Update(categoryId string, categoryUC model.Category) error {
	ret := c.Mock.Called(categoryId, categoryUC)

	return ret.Error(0)
}

func (c *CategoryRepositoryMock) Delete(categoryId string) (uint, error) {
	ret := c.Mock.Called(categoryId)

	return ret.Get(0).(uint), ret.Error(1)
}
