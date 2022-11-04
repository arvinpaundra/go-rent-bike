package dto

type BikeDTO struct {
	RenterId     string  `json:"renter_id" form:"renter_id"`
	CategoryId   string  `json:"category_id" form:"category_id"`
	Name         string  `json:"name" form:"name"`
	PricePerHour float32 `json:"price_per_hour" form:"price_per_hour"`
	Condition    string  `json:"condition" form:"condition"`
	Description  string  `json:"description" form:"description"`
	IsAvailable  string  `json:"is_available" form:"is_available"`
}
