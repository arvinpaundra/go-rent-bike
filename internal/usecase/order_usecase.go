package usecase

import (
	"errors"
	"time"

	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	pgMidtrans "github.com/arvinpaundra/go-rent-bike/internal/midtrans"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"github.com/arvinpaundra/go-rent-bike/internal/repository/gormdb"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
)

type OrderUsecase interface {
	CreateOrder(orderDTO dto.OrderDTO) (map[string]interface{}, error)
	UpdateRentStatus(orderId string) error
}

type orderUsecase struct {
	orderRepository          repository.OrderRepository
	paymentGatewayRepository pgMidtrans.PaymentGatewayRepository
	orderDetailRepository    gormdb.OrderDetailRepository
	userRepository           gormdb.UserRepository
	bikeRepository           gormdb.BikeRepository
	paymentRepository        gormdb.PaymentRepository
	historyRepository        gormdb.HistoryRepository
}

func (u orderUsecase) CreateOrder(orderDTO dto.OrderDTO) (map[string]interface{}, error) {
	u.paymentGatewayRepository.InitializeClientMidtrans()

	var err error

	// check if customer is exist
	var customer *model.User
	customer, err = u.userRepository.FindById(orderDTO.CustomerId)

	if err != nil {
		return nil, err
	}

	// check the bikes that customers choose
	// if the each bike are exist, append to slice bikes
	bikes := []model.Bike{}
	for i := range orderDTO.BikeIds {
		bike, err := u.bikeRepository.FindById(orderDTO.BikeIds[i])

		if err != nil {
			return nil, err
		} else if bike.IsAvailable == "0" {
			return nil, errors.New("bike not available")
		}

		bike.IsAvailable = "0"
		bike.UpdatedAt = time.Now()

		// update availability of each bike to "0", which is unavailable
		if err := u.bikeRepository.Update(bike.ID, *bike); err != nil {
			return nil, err
		}

		bikes = append(bikes, *bike)
	}

	// calculate total payments
	var totalPayments float32

	for i := range bikes {
		totalPayments += (bikes[i].PricePerHour * float32(orderDTO.TotalHour))
	}

	// initiate the payment, then create payment
	paymentId := uuid.NewString()
	payment := model.Payment{
		ID:            paymentId,
		PaymentStatus: "pending",
		PaymentType:   orderDTO.PaymentType,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := u.paymentRepository.Create(payment); err != nil {
		return nil, err
	}

	// initiate the order, then create order
	orderId := uuid.NewString()
	order := model.Order{
		ID:           orderId,
		UserId:       orderDTO.CustomerId,
		PaymentId:    paymentId,
		TotalPayment: totalPayments,
		TotalQty:     len(bikes),
		TotalHour:    orderDTO.TotalHour,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := u.orderRepository.Create(order); err != nil {
		return nil, err
	}

	// initiate the order detail with loop over the bikes above and append to slice bikesRented
	// then create the order detail
	bikesRented := []model.OrderDetail{}
	for i := range bikes {
		bike := model.OrderDetail{
			ID:      uuid.NewString(),
			OrderId: orderId,
			BikeId:  bikes[i].ID,
		}

		bikesRented = append(bikesRented, bike)
	}

	if err := u.orderDetailRepository.Create(bikesRented); err != nil {
		return nil, err
	}

	// initiate history, then create new user history
	history := model.History{
		ID:         uuid.NewString(),
		OrderId:    orderId,
		RentStatus: "pending payment",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := u.historyRepository.Create(history); err != nil {
		return nil, err
	}

	// set the item details to send to payment gateway
	items := []midtrans.ItemDetails{}
	for i := range bikes {
		item := midtrans.ItemDetails{
			ID:       bikes[i].ID,
			Name:     bikes[i].Name,
			Price:    int64(bikes[i].PricePerHour),
			Qty:      int32(order.TotalHour),
			Category: bikes[i].Category.Name,
		}

		items = append(items, item)
	}

	// init the request body to send to payment gateway
	snapReq := dto.PaymentGateway{
		Email:    customer.Email,
		Phone:    customer.Phone,
		OrderId:  orderId,
		GrossAmt: int64(totalPayments),
		Items:    items,
	}

	// send request to payment gateway
	snapUrl := u.paymentGatewayRepository.CreateUrlTransactionWithGateway(snapReq)

	// setup response from my app
	data := map[string]interface{}{
		"order_id":       order.ID,
		"total_payments": totalPayments,
		"payments": map[string]interface{}{
			"id":             payment.ID,
			"payment_status": payment.PaymentStatus,
			"payment_type":   payment.PaymentType,
			"created_at":     payment.CreatedAt,
			"updated_at":     payment.UpdatedAt,
		},
		"payment_link": snapUrl,
	}

	return data, nil
}

func (u orderUsecase) UpdateRentStatus(orderId string) error {
	var err error

	var order *model.Order
	order, err = u.orderRepository.FindById(orderId)
	if err != nil {
		return err
	}

	// get the bike from Order.OrderDetails
	for i := range order.OrderDetails {
		bike := order.OrderDetails[i].Bike

		// update the bike availability
		bike.IsAvailable = "1"

		err := u.bikeRepository.Update(bike.ID, *bike)

		if err != nil {
			return err
		}
	}

	var history *model.History
	history, err = u.historyRepository.FindByIdOrder(orderId)

	if err != nil {
		return err
	}

	history.RentStatus = "done"

	err = u.historyRepository.Update(orderId, *history)

	if err != nil {
		return err
	}

	return nil
}

func NewOrderUsecase(orderRepo repository.OrderRepository) OrderUsecase {
	return orderUsecase{
		orderRepository: orderRepo,
	}
}
