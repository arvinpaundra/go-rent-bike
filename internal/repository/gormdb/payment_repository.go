package gormdb

import (
	"errors"
	"github.com/arvinpaundra/go-rent-bike/pkg"

	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	DB *gorm.DB
}

func (r PaymentRepository) Create(paymentUC model.Payment) error {
	err := r.DB.Model(&model.Payment{}).Create(&paymentUC).Error

	if err != nil {
		return err
	}

	return nil
}

func (r PaymentRepository) FindById(paymentId string) (*model.Payment, error) {
	payment := &model.Payment{}

	err := r.DB.Model(&model.Payment{}).Where("id = ?", paymentId).Take(&payment).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkg.ErrRecordNotFound
		}

		return nil, err
	}

	return payment, nil
}

func (r PaymentRepository) Update(paymentId string, paymentUC model.Payment) error {
	err := r.DB.Model(&model.Payment{}).Where("id = ?", paymentId).Updates(&paymentUC).Error

	if err != nil {
		return err
	}

	return nil
}

func NewPaymentRepository(db *gorm.DB) repository.PaymentRepository {
	return PaymentRepository{db}
}
