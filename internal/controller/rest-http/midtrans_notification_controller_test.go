package rest_http

import (
	"bytes"
	"encoding/json"
	usecasemock "github.com/arvinpaundra/go-rent-bike/internal/usecase/mock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type suiteMidtrans struct {
	suite.Suite
	handler *MidtransNotificationController
	mocking *usecasemock.PaymentGatewayMock
}

func (s *suiteMidtrans) SetupSuite() {
	mock := &usecasemock.PaymentGatewayMock{}
	s.mocking = mock

	s.handler = &MidtransNotificationController{
		paymentGatewayUsecase: s.mocking,
	}
}

func (s *suiteMidtrans) TestHandlerNotification() {
	orderId := "478b3f5e-284e-440c-8c0f-af4f94c70d87"

	s.mocking.Mock.On("MidtransNotification", orderId).Return(nil)

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
			Name:               "success create category",
			ExpectedStatusCode: http.StatusOK,
			Method:             "POST",
			Header: map[string]string{
				"Content-Type": "application/json",
			},
			Body: map[string]interface{}{
				"order_id":           "478b3f5e-284e-440c-8c0f-af4f94c70d87",
				"transaction_status": "success",
				"fraud_status":       "accept",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success update transaction status",
				"data":    nil,
			},
		},
		{
			Name:               "failed wrong content-type",
			ExpectedStatusCode: http.StatusBadRequest,
			Method:             "POST",
			Header: map[string]string{
				"Content-Type": "text/plain",
			},
			Body: map[string]interface{}{
				"order_id":           "478b3f5e-284e-440c-8c0f-af4f94c70d87",
				"transaction_status": "success",
				"fraud_status":       "accept",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"status":  "error",
				"message": "fill all required fields",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/webhook", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/midtrans")
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])

			err := s.handler.HandlerNotification(ctx)
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

func (s *suiteMidtrans) TearDownSuite() {
	s.mocking = nil
}

func TestSuiteMidtrans(t *testing.T) {
	suite.Run(t, new(suiteMidtrans))
}
