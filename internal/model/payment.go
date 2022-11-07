package model

import "time"

type Payment struct {
	ID            string    `json:"id" gorm:"size:255"`
	PaymentStatus string    `json:"payment_status" gorm:"size:20"`
	PaymentType   string    `json:"payment_type" gorm:"size:50"`
	PaymentLink   string    `json:"payment_link" gorm:"size:255"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
