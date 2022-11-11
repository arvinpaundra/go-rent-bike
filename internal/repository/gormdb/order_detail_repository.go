package gormdb

import (
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"gorm.io/gorm"
)

type OrderDetailRepository struct {
	DB *gorm.DB
}

func (r OrderDetailRepository) Create(orderDetailUC []model.OrderDetail) error {
	err := r.DB.Model(&model.OrderDetail{}).Create(&orderDetailUC).Error

	if err != nil {
		return err
	}

	return nil
}

func (r OrderDetailRepository) FindByIdOrder(orderId string) (*[]model.OrderDetail, error) {
	details := &[]model.OrderDetail{}

	err := r.DB.Model(&model.OrderDetail{}).Where("order_id = ?", orderId).Preload("Bike.Category").Find(&details).Error

	if err != nil {
		return nil, err
	}

	return details, nil
}

func NewOrderDetailRepository(db *gorm.DB) repository.OrderDetailRepository {
	return OrderDetailRepository{db}
}
