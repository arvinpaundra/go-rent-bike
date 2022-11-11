package gormdb

import (
	"errors"
	"github.com/arvinpaundra/go-rent-bike/pkg"

	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"gorm.io/gorm"
)

type HistoryRepository struct {
	DB *gorm.DB
}

func (r HistoryRepository) Create(historyUC model.History) error {
	err := r.DB.Model(&model.History{}).Create(&historyUC).Error

	if err != nil {
		return err
	}

	return nil
}

func (r HistoryRepository) FindAll(userId string) (*[]model.History, error) {
	histories := &[]model.History{}

	err := r.DB.Model(&model.History{}).Preload("Order", "user_id = ?", userId).Find(&histories).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrRecordNotFound
		}

		return nil, err
	}

	return histories, nil
}

func (r HistoryRepository) FindByIdOrder(orderId string) (*model.History, error) {
	history := &model.History{}

	err := r.DB.Model(&model.History{}).Where("order_id = ?", orderId).Take(&history).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrRecordNotFound
		}

		return nil, err
	}

	return history, nil
}

func (r HistoryRepository) Update(orderId string, historyUC model.History) error {
	err := r.DB.Model(&model.History{}).Where("order_id = ?", orderId).Updates(&historyUC).Error

	if err != nil {
		return err
	}

	return nil
}

func NewHistoryRepository(db *gorm.DB) repository.HistoryRepository {
	return HistoryRepository{db}
}
