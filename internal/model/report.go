package model

import "time"

type Report struct {
	ID         string    `json:"id"`
	RenterId   string    `json:"renter_id"`
	UserId     string    `json:"user_id"`
	TitleIssue string    `json:"title_issue"`
	BodyIssue  string    `json:"body_issue"`
	User       *User     `json:"user,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
