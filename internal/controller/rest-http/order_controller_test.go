package rest_http

import (
	"bytes"
	"encoding/json"
	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	mocking "github.com/arvinpaundra/go-rent-bike/internal/usecase/mock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type suiteOrders struct {
	suite.Suite
	handler *OrderController
	mocking *mocking.OrderUsecaseMock
}

func (s *suiteOrders) SetupSuite() {
	mock := &mocking.OrderUsecaseMock{}
	s.mocking = mock

	s.handler = &OrderController{
		orderUsecase: s.mocking,
	}
}

func (s *suiteOrders) TestHandlerCreateNewOrder() {
	orderDTO := dto.OrderDTO{
		CustomerId:  "81cef832-e4de-4588-a026-6a106cf10a19",
		BikeIds:     []string{"92d88bd9-d3d2-4bd5-adba-a8161cc26cc1"},
		TotalHour:   4,
		PaymentType: "bank_transfer",
	}

	payment := model.Payment{
		ID:            "47fed3fe-5718-4b20-a525-a914ab80ba5a",
		PaymentStatus: "pending",
		PaymentType:   "bank_transfer",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	order := map[string]interface{}{
		"order_id":       "53d60e0e-8b92-416b-ab2d-0b645f54483e",
		"total_payments": float32(200000),
		"payments": map[string]interface{}{
			"id":             payment.ID,
			"payment_status": payment.PaymentStatus,
			"payment_type":   payment.PaymentType,
			"created_at":     payment.CreatedAt,
			"updated_at":     payment.UpdatedAt,
		},
		"payment_link": "https://app.sandbox.midtrans.com/v3/redirect/....",
	}

	s.mocking.Mock.On("CreateOrder", orderDTO).Return(order, nil)

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Header             map[string]string
		Body               map[string]interface{}
		HasReturnBody      bool
		ExpectedResult     map[string]interface{}
	}{
		{
			Name:               "success create order",
			ExpectedStatusCode: http.StatusCreated,
			Method:             "POST",
			Header: map[string]string{
				"Content-Type": "application/json",
			},
			Body: map[string]interface{}{
				"customer_id":  "81cef832-e4de-4588-a026-6a106cf10a19",
				"bike_ids":     []string{"92d88bd9-d3d2-4bd5-adba-a8161cc26cc1"},
				"total_hour":   int(4),
				"payment_type": "bank_transfer",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success create an order",
				"data":    order,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/orders", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])

			err := s.handler.HandlerCreateNewOrder(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				expectedOrder := v.ExpectedResult["data"].(map[string]interface{})
				orderRes := resp["data"].(map[string]interface{})

				s.Equal(v.ExpectedResult["status"], resp["status"])
				s.Equal(v.ExpectedResult["message"], resp["message"])
				s.Equal(expectedOrder["order_id"], orderRes["order_id"])
				s.Equal(expectedOrder["payment_link"], orderRes["payment_link"])
				s.Equal(expectedOrder["total_payments"], float32(orderRes["total_payments"].(float64)))
			}
		})
	}
}

func (s *suiteOrders) TestHandlerReturnBike() {
	orderId := "47fed3fe-5718-4b20-a525-a914ab80ba5a"

	s.mocking.Mock.On("UpdateRentStatus", orderId).Return(nil)

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		ExpectedResult     map[string]interface{}
	}{
		{
			Name:               "success return bike",
			ExpectedStatusCode: http.StatusOK,
			Method:             "GET",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success return bike",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/orders", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id/return")
			ctx.SetParamNames("id")
			ctx.SetParamValues(orderId)

			err := s.handler.HandlerReturnBike(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				s.Equal(v.ExpectedResult["status"], resp["status"])
				s.Equal(v.ExpectedResult["message"], resp["message"])
			}
		})
	}
}

func (s *suiteOrders) TearDownSuite() {
	s.mocking = nil
}

func TestSuiteOrders(t *testing.T) {
	suite.Run(t, new(suiteOrders))
}
