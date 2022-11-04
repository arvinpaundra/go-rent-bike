package gormdb

import (
	"github.com/arvinpaundra/go-rent-bike/database"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"gorm.io/gorm"
)

type OrderDetailRepository struct {
	DB *gorm.DB
}

func (r OrderDetailRepository) Create(orderDetailUC []model.OrderDetail) error {
	err := database.DB.Model(&model.OrderDetail{}).Create(&orderDetailUC).Error

	if err != nil {
		return err
	}

	return nil
}

func NewOrderDetailRepository(db *gorm.DB) repository.OrderDetailRepository {
	return OrderDetailRepository{db}
}
