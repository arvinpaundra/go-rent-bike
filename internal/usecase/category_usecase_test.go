package usecase

import (
	"testing"
	"time"

	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryUsecaseTest = NewCategoryUsecase(&pkg.CategoryRepository)

func TestCategoryUsecase_CreateCategory(t *testing.T) {
	pkg.CategoryRepository.Mock.On("Create", mock.Anything).Return(nil)

	categoryDTO := dto.CategoryDTO{
		Name: "Fixie",
	}

	err := categoryUsecaseTest.CreateCategory(categoryDTO)

	assert.Nil(t, err)
}

func TestCategoryUsecase_FindAllCategory(t *testing.T) {
	expectedCategories := &[]model.Category{
		{
			ID:        "169c38a9-7047-4216-bc2b-869db969a239",
			Name:      "Fixie",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "41efbedc-4319-4d36-970e-c45dfcdab4f0",
			Name:      "BMX",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	pkg.CategoryRepository.Mock.On("FindAll").Return(expectedCategories, nil)

	results, err := categoryUsecaseTest.FindAllCategories()

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, (*expectedCategories)[0].ID, (*results)[0].ID)
	assert.Equal(t, (*expectedCategories)[0].Name, (*results)[0].Name)
}

func TestCategoryUsecase_FindByIdCategory(t *testing.T) {
	categoryId := "169c38a9-7047-4216-bc2b-869db969a239"

	category := &model.Category{
		ID:        "169c38a9-7047-4216-bc2b-869db969a239",
		Name:      "Fixie",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	pkg.CategoryRepository.Mock.On("FindById", categoryId).Return(category, nil)

	result, err := categoryUsecaseTest.FindByIdCategory(categoryId)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.ID, result.ID)
	assert.Equal(t, category.Name, result.Name)
}

func TestCategoryUsecase_UpdateCategory(t *testing.T) {
	categoryId := "169c38a9-7047-4216-bc2b-869db969a239"

	category := &model.Category{
		ID:        "169c38a9-7047-4216-bc2b-869db969a239",
		Name:      "BMX",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	pkg.CategoryRepository.Mock.On("FindById", categoryId).Return(category, nil)

	pkg.CategoryRepository.Mock.On("Update", categoryId, mock.Anything).Return(nil)

	categoryDTO := dto.CategoryDTO{
		Name: "BMX",
	}

	err := categoryUsecaseTest.UpdateCategory(categoryId, categoryDTO)

	assert.Nil(t, err)
}

func TestCategoryUsecase_DeleteCategory(t *testing.T) {
	categoryId := "169c38a9-7047-4216-bc2b-869db969a239"

	category := &model.Category{
		ID:        "169c38a9-7047-4216-bc2b-869db969a239",
		Name:      "BMX",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	pkg.CategoryRepository.Mock.On("FindById", categoryId).Return(category, nil)

	pkg.CategoryRepository.Mock.On("Delete", categoryId).Return(nil)

	err := categoryUsecaseTest.DeleteCategory(categoryId)

	assert.Nil(t, err)
}
