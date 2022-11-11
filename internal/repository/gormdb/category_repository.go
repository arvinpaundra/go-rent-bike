package gormdb

import (
	"errors"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"github.com/arvinpaundra/go-rent-bike/pkg"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func (r CategoryRepository) Create(categoryUC model.Category) error {
	err := r.DB.Model(&model.Category{}).Create(&categoryUC).Error

	if err != nil {
		return err
	}

	return nil
}

func (r CategoryRepository) FindAll() (*[]model.Category, error) {
	categories := &[]model.Category{}

	err := r.DB.Model(&model.Category{}).Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r CategoryRepository) FindById(categoryId string) (*model.Category, error) {
	category := &model.Category{}

	err := r.DB.Model(&model.Category{}).Where("id = ?", categoryId).Take(&category).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrRecordNotFound
		}

		return nil, err
	}

	return category, nil
}

func (r CategoryRepository) Update(categoryId string, categoryUC model.Category) error {
	err := r.DB.Model(&model.Category{}).Where("id = ?", categoryId).Updates(&categoryUC).Error

	if err != nil {
		return err
	}

	return nil
}

func (r CategoryRepository) Delete(categoryId string) error {
	err := r.DB.Model(&model.Category{}).Where("id = ?", categoryId).Delete(&model.Category{}).Error

	if err != nil {
		return err
	}

	return nil
}

func NewCategoryRepositoryGorm(db *gorm.DB) repository.CategoryRepository {
	return CategoryRepository{db}
}
