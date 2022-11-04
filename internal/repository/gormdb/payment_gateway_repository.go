package gormdb

import (
	"context"
	"fmt"

	"github.com/arvinpaundra/go-rent-bike/configs"
	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

var snapClient snap.Client

type PaymentGatewayRepository struct{}

func (r PaymentGatewayRepository) InitializeClientMidtrans() {
	snapClient.New(configs.Cfg.MidtransServerKeyDev, midtrans.Sandbox)
}

func (r PaymentGatewayRepository) CreateTransaction(req dto.PaymentGateway) string {
	snapUrl, err := snapClient.CreateTransactionToken(generateSnapReq(req))

	if err != nil {
		fmt.Printf("Midtrans error : %v", err.GetMessage())
	}

	return snapUrl
}

func (r PaymentGatewayRepository) CreateUrlTransactionWithGateway(req dto.PaymentGateway) string {
	snapClient.Options.SetContext(context.Background())

	snapUrl, err := snapClient.CreateTransactionUrl(generateSnapReq(req))

	if err != nil {
		fmt.Printf("Midtrans error : %v", err.GetMessage())
	}

	return snapUrl
}

func generateSnapReq(req dto.PaymentGateway) *snap.Request {
	reqSnap := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  req.OrderId,
			GrossAmt: req.GrossAmt,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		CustomerDetail: &midtrans.CustomerDetails{
			Email: req.Email,
			Phone: req.Phone,
		},
		Items: &req.Items,
	}

	return reqSnap
}

func NewPaymentGatewayRepository() repository.PaymentGatewayRepository {
	return PaymentGatewayRepository{}
}
