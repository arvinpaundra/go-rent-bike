package dto

type RenterDTO struct {
	UserId      string `json:"user_id" form:"user_id"`
	RentName    string `json:"rent_name" form:"rent_name"`
	RentAddress string `json:"rent_address" form:"rent_address"`
	Description string `json:"description" form:"description"`
}
