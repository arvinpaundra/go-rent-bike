package dto

type ReviewDTO struct {
	UserId      string `json:"user_id" form:"user_id"`
	Rating      int    `json:"rating" form:"rating"`
	Description string `json:"description" form:"description"`
}
