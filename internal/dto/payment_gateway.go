package dto

import "github.com/midtrans/midtrans-go"

type PaymentGateway struct {
	Email    string
	Phone    string
	OrderId  string
	GrossAmt int64
	Items    []midtrans.ItemDetails
}
