package usecase

import (
	"time"

	"github.com/arvinpaundra/go-rent-bike/configs"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository/gormdb"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

var c coreapi.Client

type PaymentGatewayUsecase interface {
	MidtransNotification(orderId string) error
}

type paymentGatewayUsecase struct {
	orderRepository   gormdb.OrderRepository
	paymentRepository gormdb.PaymentRepository
	historyRepository gormdb.HistoryRepository
}

func (u paymentGatewayUsecase) InitializeCoreapiClient() {
	c.New(configs.Cfg.MidtransServerKeyDev, midtrans.Sandbox)
}

func (u paymentGatewayUsecase) MidtransNotification(orderId string) error {
	u.InitializeCoreapiClient()

	transactionStatusRes, midtransError := c.CheckTransaction(orderId)

	if midtransError != nil {
		return midtransError
	}

	order, err := u.orderRepository.FindById(orderId)

	if err != nil {
		return err
	}

	var payment *model.Payment
	payment, err = u.paymentRepository.FindById(order.PaymentId)

	if err != nil {
		return err
	}

	var history *model.History
	history, err = u.historyRepository.FindByIdOrder(orderId)

	if err != nil {
		return err
	}

	if transactionStatusRes.TransactionStatus == "settlement" && transactionStatusRes.FraudStatus == "accept" {
		payment.PaymentStatus = "settlement"
		payment.PaymentType = transactionStatusRes.PaymentType
		payment.UpdatedAt = time.Now()

		if err := u.paymentRepository.Update(payment.ID, *payment); err != nil {
			return err
		}

		history.RentStatus = "rented"
		history.UpdatedAt = time.Now()

		if err := u.historyRepository.Update(orderId, *history); err != nil {
			return err
		}
	} else if transactionStatusRes.TransactionStatus == "deny" {
		payment.PaymentStatus = "deny"
		payment.PaymentType = transactionStatusRes.PaymentType
		payment.UpdatedAt = time.Now()

		if err := u.paymentRepository.Update(payment.ID, *payment); err != nil {
			return err
		}

		history.RentStatus = "denied"
		history.UpdatedAt = time.Now()

		if err := u.historyRepository.Update(orderId, *history); err != nil {
			return err
		}
	} else if transactionStatusRes.TransactionStatus == "cancel" || transactionStatusRes.TransactionStatus == "expired" {
		payment.PaymentStatus = transactionStatusRes.TransactionStatus
		payment.PaymentType = transactionStatusRes.PaymentType
		payment.UpdatedAt = time.Now()

		if err := u.paymentRepository.Update(payment.ID, *payment); err != nil {
			return err
		}

		history.RentStatus = "canceled"
		history.UpdatedAt = time.Now()

		if err := u.historyRepository.Update(orderId, *history); err != nil {
			return err
		}
	} else if transactionStatusRes.TransactionStatus == "pending" {
		payment.PaymentType = transactionStatusRes.PaymentType
		payment.UpdatedAt = time.Now()

		if err := u.paymentRepository.Update(payment.ID, *payment); err != nil {
			return err
		}
	}

	return nil
}

func NewPaymentGatewayUsecase() PaymentGatewayUsecase {
	return paymentGatewayUsecase{}
}
