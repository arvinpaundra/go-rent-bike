package model

import "time"

type Category struct {
	ID        string    `json:"id" gorm:"primaryKey;size:255"`
	Name      string    `json:"name" gorm:"size:100"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
