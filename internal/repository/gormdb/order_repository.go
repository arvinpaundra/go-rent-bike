package gormdb

import (
	"errors"

	"github.com/arvinpaundra/go-rent-bike/database"
	"github.com/arvinpaundra/go-rent-bike/internal"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderRepository struct {
	DB *gorm.DB
}

func (r OrderRepository) Create(orderUC model.Order) error {
	err := database.DB.Model(&model.Order{}).Create(&orderUC).Error

	if err != nil {
		return err
	}

	return nil
}

func (r OrderRepository) FindAll(userId string) (*[]model.Order, error) {
	orders := &[]model.Order{}

	err := database.DB.Model(&model.Order{}).Where("user_id = ?", userId).Find(&orders).Error

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r OrderRepository) FindById(orderId string) (*model.Order, error) {
	order := &model.Order{}

	err := database.DB.Model(&model.Order{}).Where("id = ?", orderId).Preload("OrderDetails.Bike.Category").Preload(clause.Associations).Take(&order).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, internal.ErrRecordNotFound
		}

		return nil, err
	}

	return order, nil
}

func (r OrderRepository) Update(orderId string, orderUC model.Order) error {
	err := database.DB.Model(&model.Order{}).Where("id = ?", orderId).Save(&orderUC).Error

	if err != nil {
		return err
	}

	return nil
}

func NewOrderRepository(db *gorm.DB) repository.OrderRepository {
	return OrderRepository{db}
}
