package gormdb

import (
	"github.com/arvinpaundra/go-rent-bike/database"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"gorm.io/gorm"
)

type ReviewRepository struct {
	DB *gorm.DB
}

func (r ReviewRepository) Create(reviewUC model.Review) error {
	err := database.DB.Model(&model.Review{}).Create(reviewUC).Error

	if err != nil {
		return err
	}

	return nil
}

func NewReviewRepositoryGorm(db *gorm.DB) repository.ReviewRepository {
	return ReviewRepository{db}
}
