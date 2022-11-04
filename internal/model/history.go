package model

import "time"

type History struct {
	ID         string    `json:"id" gorm:"primaryKey;size:255"`
	OrderId    string    `json:"order_id" gorm:"size:255"`
	RentStatus string    `json:"rent_status" gorm:"size:50"`
	Order      Order     `json:"order"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
