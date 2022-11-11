package pg_midtrans

import (
	"context"
	"fmt"

	"github.com/arvinpaundra/go-rent-bike/configs"
	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

var snapClient snap.Client

type PaymentGatewayInterface interface {
	InitializeClientMidtrans()
	CreateTransaction(snap dto.PaymentGateway) string
	CreateUrlTransactionWithGateway(snap dto.PaymentGateway) string
}

type PaymentGateway struct{}

func (r PaymentGateway) InitializeClientMidtrans() {
	snapClient.New(configs.Cfg.MidtransServerKeyDev, midtrans.Sandbox)
}

func (r PaymentGateway) CreateTransaction(req dto.PaymentGateway) string {
	snapUrl, err := snapClient.CreateTransactionToken(generateSnapReq(req))

	if err != nil {
		fmt.Printf("Midtrans error : %v", err.GetMessage())
	}

	return snapUrl
}

func (r PaymentGateway) CreateUrlTransactionWithGateway(req dto.PaymentGateway) string {
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
		EnabledPayments: []snap.SnapPaymentType{
			snap.PaymentTypeBNIVA,
			snap.PaymentTypePermataVA,
			snap.PaymentTypeBCAVA,
			snap.PaymentTypeBRIVA,
			snap.PaymentTypeBankTransfer,
			snap.PaymentTypeGopay,
			snap.PaymentTypeShopeepay,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			Email: req.Email,
			Phone: req.Phone,
		},
		Items: &req.Items,
	}

	return reqSnap
}
