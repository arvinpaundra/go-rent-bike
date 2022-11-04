package model

import "time"

type Bike struct {
	ID           string    `json:"id" gorm:"primaryKey;size:255"`
	RenterId     string    `json:"renter_id" gorm:"size:255"`
	CategoryId   string    `json:"category_id" gorm:"size:255"`
	Name         string    `json:"name" gorm:"size:255"`
	PricePerHour float32   `json:"price_per_hour"`
	Condition    string    `json:"condition" gorm:"size:100"`
	Description  string    `json:"description"`
	IsAvailable  string    `json:"is_available" gorm:"size:1"`
	Category     Category  `json:"category"`
	Reviews      []Review  `json:"reviews"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
