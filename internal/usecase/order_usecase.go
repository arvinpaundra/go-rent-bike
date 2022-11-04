package usecase

import (
	"time"

	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"github.com/arvinpaundra/go-rent-bike/internal/repository/gormdb"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
)

type OrderUsecase interface {
	CreateOrder(orderDTO dto.OrderDTO) (map[string]interface{}, error)
	FindAllOrdersUser(userId string) (*[]model.Order, error)
	FindByIdOrderUser(orderId string) (*model.Order, error)
	UpdateOrder(orderId string) (map[string]interface{}, error)
}

type orderUsecase struct {
	orderRepository          repository.OrderRepository
	paymentGatewayRepository gormdb.PaymentGatewayRepository
	orderDetailRepository    gormdb.OrderDetailRepository
	userRepository           gormdb.UserRepository
	bikeRepository           gormdb.BikeRepository
	paymentRepository        gormdb.PaymentRepository
}

func (u orderUsecase) CreateOrder(orderDTO dto.OrderDTO) (map[string]interface{}, error) {
	u.paymentGatewayRepository.InitializeClientMidtrans()

	var err error

	// check if customer is exist
	customer := &model.User{}
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
		TotalPayment:  totalPayments,
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
		ID:        orderId,
		UserId:    orderDTO.CustomerId,
		PaymentId: paymentId,
		TotalQty:  len(bikes),
		TotalHour: orderDTO.TotalHour,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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

	snapReq := dto.PaymentGateway{
		Email:    customer.Email,
		Phone:    customer.Phone,
		OrderId:  orderId,
		GrossAmt: int64(totalPayments),
		Items:    items,
	}

	snapUrl := u.paymentGatewayRepository.CreateUrlTransactionWithGateway(snapReq)

	data := map[string]interface{}{
		"order_id": order.ID,
		"payments": map[string]interface{}{
			"id":             payment.ID,
			"total_payment":  payment.TotalPayment,
			"payment_status": payment.PaymentStatus,
			"payment_type":   payment.PaymentType,
			"created_at":     payment.CreatedAt,
			"updated_at":     payment.UpdatedAt,
		},
		"payment_link": snapUrl,
	}

	return data, nil
}

func (u orderUsecase) FindAllOrdersUser(userId string) (*[]model.Order, error) {
	if _, err := u.userRepository.FindById(userId); err != nil {
		return nil, err
	}

	orders, err := u.orderRepository.FindAll(userId)

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (u orderUsecase) FindByIdOrderUser(orderId string) (*model.Order, error) {
	order, err := u.orderRepository.FindById(orderId)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (u orderUsecase) UpdateOrder(orderId string) (map[string]interface{}, error) {
	panic("implement me")
}

func NeworderUsecase(orderRepo repository.OrderRepository) OrderUsecase {
	return orderUsecase{
		orderRepository: orderRepo,
	}
}
