package gormdb

import (
	"github.com/arvinpaundra/go-rent-bike/database"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"gorm.io/gorm"
)

type ReportRepository struct {
	DB *gorm.DB
}

func (r ReportRepository) Create(reportUC model.Report) error {
	err := database.DB.Model(&model.Report{}).Create(&reportUC).Error

	if err != nil {
		return err
	}

	return nil
}

func (r ReportRepository) FindAll(renterId string) (*[]model.Report, error) {
	reports := &[]model.Report{}

	err := database.DB.Model(&model.Report{}).Where("renter_id = ?", renterId).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Omit("password")
	}).Find(&reports).Error

	if err != nil {
		return nil, err
	}

	return reports, nil
}

func NewReportRepository(db *gorm.DB) repository.ReportRepository {
	return ReportRepository{db}
}
