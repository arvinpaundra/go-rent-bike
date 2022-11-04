package model

import "time"

type User struct {
	ID        string    `json:"id" gorm:"primaryKey;size:255"`
	Fullname  string    `json:"fullname" gorm:"size:255"`
	Phone     string    `json:"phone" gorm:"size:13"`
	Address   string    `json:"address"`
	Role      string    `json:"role" gorm:"size:50"`
	Email     string    `json:"email" gorm:"size:255"`
	Password  string    `json:"password" gorm:"size:255"`
	Orders    []Order   `json:"orders"`
	Reviews   []Review  `json:"reviews"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
