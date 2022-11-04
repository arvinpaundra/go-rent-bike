package dto

type OrderDTO struct {
	CustomerId  string   `json:"customer_id" form:"customer_id"`
	BikeIds     []string `json:"bike_ids" form:"bike_ids"`
	TotalHour   int      `json:"total_hour" form:"total_hour"`
	PaymentType string   `json:"payment_type" form:"payment_type"`
}
