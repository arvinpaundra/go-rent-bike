package usecase

import (
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var orderUsecaseTest = NewOrderUsecase(
	&pkg.OrderRepository,
	&pkg.OrderDetailRepository,
	&pkg.UserRepository,
	&pkg.BikeRepository,
	&pkg.PaymentRepository,
	&pkg.HistoryRepository,
)

// TODO belum berhasil buat test midtrans
//func TestOrderUsecase_CreateOrder(t *testing.T) {
//	pgMidtrans.PaymentGateway{}.InitializeClientMidtrans()
//
//	customerId := "28dc0243-7553-4ebc-9937-a0f5505df7e3"
//
//	customer := &model.User{
//		ID:        "28dc0243-7553-4ebc-9937-a0f5505df7e3",
//		Fullname:  "Arvin Paundra",
//		Phone:     "0876534321",
//		Address:   "Jl Rinjani",
//		Role:      "customer",
//		Email:     "arvin@mail.com",
//		CreatedAt: time.Now(),
//		UpdatedAt: time.Now(),
//	}
//
//	pkg.UserRepository.Mock.On("FindById", customerId).Return(customer, nil)
//
//	orderDTO := dto.OrderDTO{
//		CustomerId:  "28dc0243-7553-4ebc-9937-a0f5505df7e3",
//		BikeIds:     []string{"c4b10642-95a5-4aca-a612-fdc1b0837f37"},
//		TotalHour:   5,
//		PaymentType: "bank_transfer",
//	}
//
//	bikes := &[]model.Bike{
//		{
//			ID:           "c4b10642-95a5-4aca-a612-fdc1b0837f37",
//			RenterId:     "ffad8203-b32d-46dd-b488-a700ad61dac7",
//			CategoryId:   "463c249c-300e-4138-8cfb-8a7fa5042da0",
//			Name:         "Sample BMX Bike",
//			PricePerHour: 15000,
//			Condition:    "Good",
//			Description:  "This is a description section",
//			IsAvailable:  "1",
//		},
//	}
//
//	for i := range orderDTO.BikeIds {
//		pkg.BikeRepository.Mock.On("FindById").Return((*bikes)[i], nil)
//
//		(*bikes)[i].IsAvailable = "0"
//		(*bikes)[i].UpdatedAt = time.Now()
//
//		pkg.BikeRepository.Mock.On("Update", (*bikes)[i].ID, (*bikes)[i]).Return(nil)
//	}
//
//	paymentUC := model.Payment{
//		ID:            "9b5d19ff-0406-46b6-b9f5-d64c2fd9a5bd",
//		PaymentStatus: "pending",
//		PaymentType:   orderDTO.PaymentType,
//		PaymentLink:   "https://app.sandbox.midtrans.com/snap/redirect/v3/...",
//		CreatedAt:     time.Now(),
//		UpdatedAt:     time.Now(),
//	}
//
//	pkg.PaymentRepository.Mock.On("Create", paymentUC).Return(nil)
//
//	var totalPayments float32
//
//	for i := range *bikes {
//		totalPayments += ((*bikes)[i].PricePerHour * float32(orderDTO.TotalHour))
//	}
//
//	orderUC := model.Order{
//		ID:           "2a47fd52-893f-4717-a2c3-4bd2c68912af",
//		UserId:       "28dc0243-7553-4ebc-9937-a0f5505df7e3",
//		PaymentId:    "9b5d19ff-0406-46b6-b9f5-d64c2fd9a5bd",
//		TotalPayment: totalPayments,
//		TotalQty:     len(*bikes),
//		TotalHour:    orderDTO.TotalHour,
//		CreatedAt:    time.Now(),
//		UpdatedAt:    time.Now(),
//	}
//
//	pkg.OrderRepository.Mock.On("Create", orderUC).Return(nil)
//
//	orderDetailsUC := &[]model.OrderDetail{
//		{
//			ID:      "3a85833f-e946-4729-9c89-48d087b89e86\n\n",
//			OrderId: "2a47fd52-893f-4717-a2c3-4bd2c68912af",
//			BikeId:  "c4b10642-95a5-4aca-a612-fdc1b0837f37",
//			Bike: &model.Bike{
//				ID:           "c4b10642-95a5-4aca-a612-fdc1b0837f37",
//				RenterId:     "ffad8203-b32d-46dd-b488-a700ad61dac7",
//				CategoryId:   "463c249c-300e-4138-8cfb-8a7fa5042da0",
//				Name:         "Sample BMX Bike",
//				PricePerHour: 15000,
//				Condition:    "Good",
//				Description:  "This is a description section",
//				IsAvailable:  "1",
//			},
//		},
//	}
//
//	pkg.OrderDetailRepository.Mock.On("Create", orderDetailsUC).Return(nil)
//
//	historyUC := model.History{
//		ID:         "6872428e-2fb4-4100-a400-a8b434998ba8",
//		OrderId:    "2a47fd52-893f-4717-a2c3-4bd2c68912af",
//		RentStatus: "pending payment",
//		CreatedAt:  time.Now(),
//		UpdatedAt:  time.Now(),
//	}
//
//	pkg.HistoryRepository.Mock.On("Create", historyUC).Return(nil)
//
//	//items := []midtrans.ItemDetails{
//	//	{
//	//		ID:       (*bikes)[1].ID,
//	//		Name:     (*bikes)[1].Name,
//	//		Price:    int64((*bikes)[1].PricePerHour),
//	//		Qty:      int32(orderUC.TotalHour),
//	//		Category: (*bikes)[1].Category.Name,
//	//	},
//	//}
//
//	//snapReq := dto.PaymentGateway{
//	//	Email:    customer.Email,
//	//	Phone:    customer.Phone,
//	//	OrderId:  orderUC.ID,
//	//	GrossAmt: int64(totalPayments),
//	//	Items:    items,
//	//}
//
//	snapUrl := fmt.Sprintf("https://app.sandbox.midtrans.com/snap/v3/redirection/%v", orderUC.ID)
//
//	paymentUC.PaymentLink = snapUrl
//	pkg.PaymentRepository.Mock.On("Update", paymentUC.ID, paymentUC).Return(nil)
//
//	result, err := orderUsecaseTest.CreateOrder(orderDTO)
//
//	assert.Nil(t, err)
//	assert.NotNil(t, result)
//}

func TestOrderUsecase_UpdateRentStatus(t *testing.T) {
	orderId := "a1dcbf01-144c-4507-939c-449c18d5fbac"

	order := &model.Order{
		ID:           "a1dcbf01-144c-4507-939c-449c18d5fbac",
		UserId:       "02629953-7ac7-4c77-83c0-136a0f252427",
		PaymentId:    "8807fa35-2305-47e1-86ed-d95deed566e4",
		TotalPayment: 200000,
		TotalQty:     1,
		TotalHour:    5,
		OrderDetails: []model.OrderDetail{
			{
				ID:      "6faa175f-ee36-4489-a7b6-424b82a1b855",
				OrderId: "a1dcbf01-144c-4507-939c-449c18d5fbac",
				BikeId:  "aefde097-3145-4961-9eed-9e916b9def36",
				Bike: &model.Bike{
					ID:           "aefde097-3145-4961-9eed-9e916b9def36",
					RenterId:     "ffad8203-b32d-46dd-b488-a700ad61dac7",
					CategoryId:   "463c249c-300e-4138-8cfb-8a7fa5042da0",
					Name:         "Sample BMX Bike",
					PricePerHour: 15000,
					Condition:    "Good",
					Description:  "This is a description section",
					IsAvailable:  "0",
				},
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	pkg.OrderRepository.Mock.On("FindById", orderId).Return(order, nil)

	for i := range order.OrderDetails {
		bike := order.OrderDetails[i].Bike

		// update the bike availability
		bike.IsAvailable = "1"

		pkg.BikeRepository.Mock.On("Update", bike.ID, *bike).Return(nil)
	}

	history := &model.History{
		ID:         "25512ed6-7969-4b84-a099-c4de82968ed7",
		OrderId:    "a1dcbf01-144c-4507-939c-449c18d5fbac",
		RentStatus: "rented",
	}

	pkg.HistoryRepository.Mock.On("FindByIdOrder", orderId).Return(history, nil)

	history.RentStatus = "done"

	pkg.HistoryRepository.Mock.On("Update", orderId, *history).Return(nil)

	err := orderUsecaseTest.UpdateRentStatus(orderId)

	assert.Nil(t, err)
}
