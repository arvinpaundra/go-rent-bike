package usecase

import (
	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
)

type CategoryUsecase interface {
	CreateCategory(categoryDTO dto.CategoryDTO) error
	FindAllCategories() (*[]model.Category, error)
	FindByIdCategory(categoryId string) (*model.Category, error)
	UpdateCategory(categoryId string, categoryDTO dto.CategoryDTO) error
	DeleteCategory(categoryId string) (uint, error)
}

type categoryUsecase struct {
	categoryRepository repository.CategoryRepository
}

func (c categoryUsecase) CreateCategory(categoryDTO dto.CategoryDTO) error {
	category := model.Category{
		Name: categoryDTO.Name,
	}

	err := c.categoryRepository.Create(category)

	if err != nil {
		return err
	}

	return nil
}

func (c categoryUsecase) FindAllCategories() (*[]model.Category, error) {
	categories, err := c.categoryRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (c categoryUsecase) FindByIdCategory(categoryId string) (*model.Category, error) {
	category, err := c.categoryRepository.FindById(categoryId)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (c categoryUsecase) UpdateCategory(categoryId string, categoryDTO dto.CategoryDTO) error {
	category := model.Category{
		Name: categoryDTO.Name,
	}

	err := c.categoryRepository.Update(categoryId, category)

	if err != nil {
		return err
	}

	return nil
}

func (c categoryUsecase) DeleteCategory(categoryId string) (uint, error) {
	rowAffected, err := c.categoryRepository.Delete(categoryId)

	if err != nil {
		return rowAffected, err
	}

	return rowAffected, nil
}

func NewCategoryUsecase(categoryRepo repository.CategoryRepository) CategoryUsecase {
	return categoryUsecase{categoryRepo}
}
