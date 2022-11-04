package gormdb

import (
	"errors"
	"time"

	"github.com/arvinpaundra/go-rent-bike/database"
	"github.com/arvinpaundra/go-rent-bike/internal"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func (c CategoryRepository) Create(categoryUC model.Category) error {
	category := model.Category{}

	_ = database.DB.Model(&model.Category{}).Where("name = ?", categoryUC.Name).Take(&category)

	if category.ID != "" {
		return internal.ErrDataAlreadyExist
	}

	newCategory := model.Category{
		ID:        uuid.NewString(),
		Name:      categoryUC.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := database.DB.Model(&model.Category{}).Create(&newCategory).Error

	if err != nil {
		return err
	}

	return nil
}

func (c CategoryRepository) FindAll() (*[]model.Category, error) {
	categories := &[]model.Category{}

	err := database.DB.Model(&model.Category{}).Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (c CategoryRepository) FindById(categoryId string) (*model.Category, error) {
	category := &model.Category{}

	err := database.DB.Model(&model.Category{}).Where("id = ?", categoryId).Take(&category).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, internal.ErrRecordNotFound
		}

		return nil, err
	}

	return category, nil
}

func (c CategoryRepository) Update(categoryId string, categoryUC model.Category) error {
	category := model.Category{}

	msg := database.DB.Model(&model.Category{}).Where("id = ?", categoryId).Take(&category).Error

	if msg != nil {
		return internal.ErrRecordNotFound
	}

	updatedCategory := model.Category{
		ID:        categoryId,
		Name:      categoryUC.Name,
		CreatedAt: category.UpdatedAt,
		UpdatedAt: time.Now(),
	}

	err := database.DB.Model(&model.Category{}).Where("id = ?", categoryId).Save(&updatedCategory).Error

	if err != nil {
		return err
	}

	return nil
}

func (c CategoryRepository) Delete(categoryId string) (uint, error) {
	category := model.Category{}

	msg := database.DB.Model(&model.Category{}).Where("id = ?", categoryId).Take(&category).Error

	if msg != nil {
		if errors.Is(msg, gorm.ErrRecordNotFound) {
			return 0, internal.ErrRecordNotFound
		}

		return 0, msg
	}

	err := database.DB.Model(&model.Category{}).Where("id = ?", categoryId).Delete(&model.Category{}).Error

	if err != nil {
		return 0, err
	}

	return 1, nil
}

func NewCategoryRepositoryGorm(db *gorm.DB) repository.CategoryRepository {
	return CategoryRepository{db}
}
