package model

import "time"

type Order struct {
	ID           string        `json:"id" gorm:"primaryKey;size:255"`
	UserId       string        `json:"user_id" gorm:"size:255"`
	PaymentId    string        `json:"payment_id" gorm:"size:255"`
	TotalQty     int           `json:"total_qty"`
	TotalHour    int           `json:"total_hour"`
	OrderDetails []OrderDetail `json:"order_details,omitempty"`
	Payment      Payment       `json:"payment_details"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}
