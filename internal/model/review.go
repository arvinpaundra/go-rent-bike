package model

import "time"

type Review struct {
	ID          string    `json:"id" gorm:"primaryKey;size:255"`
	BikeId      string    `json:"bike_id" gorm:"size:255"`
	UserId      string    `json:"user_id" gorm:"size:255"`
	Rating      int       `json:"rating"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
