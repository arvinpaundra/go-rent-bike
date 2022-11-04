package model

import "time"

type Renter struct {
	ID          string    `json:"id" gorm:"primaryKey;size:255"`
	UserId      string    `json:"user_id" gorm:"size:255"`
	RentName    string    `json:"rent_name" gorm:"size:255"`
	RentAddress string    `json:"rent_address"`
	Description string    `json:"description"`
	User        User      `json:"owner"`
	Bikes       []Bike    `json:"bikes"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
