package model

type OrderDetail struct {
	ID      string `json:"id" gorm:"primaryKey;size:255"`
	OrderId string `json:"order_id" gorm:"size:255"`
	BikeId  string `json:"bike_id" gorm:"size:255"`
	Bike    *Bike  `json:"bike,omitempty"`
}
