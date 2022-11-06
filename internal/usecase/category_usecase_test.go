package usecase

import (
	"testing"
	"time"

	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	mocking "github.com/arvinpaundra/go-rent-bike/internal/repository/gormdb/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepositoryMock = mocking.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryUsecaseTest = NewCategoryUsecase(&categoryRepositoryMock)

func TestCategoryUsecase_CreateCategory(t *testing.T) {
	categoryDTO := dto.CategoryDTO{
		Name: "Fixie",
	}

	category := model.Category{
		Name: "Fixie",
	}

	categoryRepositoryMock.Mock.On("Create", category).Return(nil)

	err := categoryUsecaseTest.CreateCategory(categoryDTO)

	assert.Nil(t, err)
}

func TestCategoryUsecase_FindAllCategory(t *testing.T) {
	expectedCategories := &[]model.Category{
		{
			ID:        uuid.NewString(),
			Name:      "Fixie",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        uuid.NewString(),
			Name:      "BMX",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	categoryRepositoryMock.Mock.On("FindAll").Return(expectedCategories, nil)

	categories, err := categoryUsecaseTest.FindAllCategories()

	assert.Nil(t, err)
	assert.NotNil(t, categories)
	assert.Equal(t, expectedCategories, categories)
}

func TestCategoryUsecase_FindByIdCategory(t *testing.T) {
	id := uuid.NewString()

	expectedCategory := &model.Category{
		ID:        id,
		Name:      "Fixie",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	categoryRepositoryMock.Mock.On("FindById", id).Return(expectedCategory, nil)

	category, err := categoryUsecaseTest.FindByIdCategory(id)

	assert.Nil(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, expectedCategory, category)
}

func TestCategoryUsecase_UpdateCategory(t *testing.T) {
	id := uuid.NewString()

	categoryDTO := dto.CategoryDTO{
		Name: "BMX",
	}

	category := model.Category{
		Name: "BMX",
	}

	categoryRepositoryMock.Mock.On("Update", id, category).Return(nil)

	err := categoryUsecaseTest.UpdateCategory(id, categoryDTO)

	assert.Nil(t, err)
}

func TestCategoryUsecase_DeleteCategory(t *testing.T) {
	id := uuid.NewString()

	categoryRepositoryMock.Mock.On("Delete", id).Return(uint(1), nil)

	rowAffected, err := categoryUsecaseTest.DeleteCategory(id)

	assert.Nil(t, err)
	assert.NotNil(t, rowAffected)
	assert.Equal(t, uint(1), rowAffected)
}
