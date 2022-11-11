package rest_http

import (
	"bytes"
	"encoding/json"
	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	usecasemock "github.com/arvinpaundra/go-rent-bike/internal/usecase/mock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type suiteUsers struct {
	suite.Suite
	handler *UserController
	mocking *usecasemock.UserUsecaseMock
}

func (s *suiteUsers) SetupSuite() {
	mock := &usecasemock.UserUsecaseMock{}
	s.mocking = mock

	s.handler = &UserController{
		userUsecase: s.mocking,
	}
}

func (s *suiteUsers) TestHandlerRegister() {
	userDTO := dto.UserDTO{
		Fullname: "Arvin Paundra",
		Phone:    "087654321",
		Address:  "Jl Rinjani",
		Role:     "customer",
		Email:    "arvin@mail.com",
		Password: "123",
	}

	s.mocking.Mock.On("RegisterUser", userDTO).Return(nil)

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
			Name:               "success register",
			ExpectedStatusCode: http.StatusCreated,
			Method:             "POST",
			Header: map[string]string{
				"Content-Type": "application/json",
			},
			Body: map[string]interface{}{
				"address":  "Jl Rinjani",
				"phone":    "087654321",
				"fullname": "Arvin Paundra",
				"role":     "customer",
				"email":    "arvin@mail.com",
				"password": "123",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "register success",
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
				"address":  "Jl Rinjani",
				"phone":    "087654321",
				"fullname": "Arvin Paundra",
				"role":     "customer",
				"email":    "arvin@mail.com",
				"password": "123",
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
			r := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("register")
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])

			err := s.handler.HandlerRegister(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				s.Equal(v.ExpectedResult, resp)
			}
		})
	}
}

func (s *suiteUsers) TestHandlerLogin() {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Njc2NzExODMsImlhdCI6MTY2NzY2NzU4Mywicm9sZSI6ImN1c3RvbWVyIiwidXNlcl9pZCI6IjJkMjcyMjUyLTdiNWQtNGY1MC04NWVlLWU1NzhlMzgyNjUxMCJ9.O2cLNLUwnfrB9mlgzS3hc4BI2k5sYNO2kVGkDEkfrAg"

	s.mocking.Mock.On("LoginUser", "arvin@mail.com", "123").Return(token, nil)

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
			Name:               "success login",
			ExpectedStatusCode: http.StatusOK,
			Method:             "POST",
			Header: map[string]string{
				"Content-Type": "application/json",
			},
			Body: map[string]interface{}{
				"email":    "arvin@mail.com",
				"password": "123",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "login success",
				"data": map[string]interface{}{
					"token": token,
				},
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
				"email":    "arvin@mail.com",
				"password": "123",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"status":  "error",
				"message": "email and password is required",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("login")
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])

			err := s.handler.HandlerLogin(ctx)
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

func (s *suiteUsers) TestHandlerFindAllUser() {
	users := &[]model.User{
		{
			Fullname: "Arvin",
			Phone:    "085",
			Role:     "customer",
			Email:    "arvin@mail.com",
		},
		{
			Fullname: "ardana",
			Phone:    "098",
			Role:     "renter",
			Email:    "ardana@mail.com",
		},
	}

	s.mocking.Mock.On("FindAllUsers").Return(users, nil)

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		ExpectedResult     map[string]interface{}
	}{
		{
			Name:               "success get users",
			ExpectedStatusCode: http.StatusOK,
			Method:             "GET",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success get all users",
				"data": map[string]interface{}{
					"users": users,
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/users", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/")

			err := s.handler.HandlerFindAllUsers(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				data := resp["data"].(map[string]interface{})
				users := data["users"].([]interface{})

				s.NotEmpty(users)
			}
		})
	}
}

func (s *suiteUsers) TestHandlerFindUserById() {
	userId := "aefde097-3145-4961-9eed-9e916b9def36"

	user := &model.User{
		Fullname: "Arvin",
		Phone:    "085",
		Role:     "customer",
		Email:    "arvin@mail.com",
	}

	s.mocking.Mock.On("FindByIdUser", userId).Return(user, nil)

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Header             map[string]string
		Body               map[string]string
		HasReturnBody      bool
		ExpectedResult     map[string]interface{}
	}{
		{
			Name:               "success get user",
			ExpectedStatusCode: http.StatusOK,
			Method:             "GET",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success get user by id",
				"data": map[string]interface{}{
					"user": user,
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/users", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(userId)

			err := s.handler.HandlerFindUserById(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				data := resp["data"].(map[string]interface{})
				user := data["user"].(map[string]interface{})

				expectRes := v.ExpectedResult["data"].(map[string]interface{})
				expectUser := expectRes["user"].(*model.User)

				s.Equal(expectUser.Email, user["email"])
				s.Equal(expectUser.Fullname, user["fullname"])
				s.Equal(expectUser.Role, user["role"])
				s.Equal(expectUser.Phone, user["phone"])
			}
		})
	}
}

func (s *suiteUsers) TestHandlerFindAllUserHistories() {
	userId := "2d272252-7b5d-4f50-85ee-e578e3826510"

	histories := &[]model.History{
		{
			ID:         "07f332fc-4a49-40a1-a7a8-72efeb2d9b8b",
			OrderId:    "b1f4bae5-4adc-4da1-bd54-b1a8584b301f",
			RentStatus: "rented",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	s.mocking.Mock.On("FindAllUserHistories", userId).Return(histories, nil)

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Header             map[string]string
		Body               map[string]string
		HasReturnBody      bool
		ExpectedResult     map[string]interface{}
	}{
		{
			Name:               "success get user histories",
			ExpectedStatusCode: http.StatusOK,
			Method:             "GET",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success get user histories",
				"data": map[string]interface{}{
					"histories": histories,
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/users", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id/histories")
			ctx.SetParamNames("id")
			ctx.SetParamValues(userId)

			err := s.handler.HandlerFindAllUserHistories(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				data := resp["data"].(map[string]interface{})
				histories := data["histories"].([]interface{})

				expectRes := v.ExpectedResult["data"].(map[string]interface{})
				expectHistories := expectRes["histories"].(*[]model.History)

				s.NotNil(histories)
				s.NotNil(expectHistories)
			}
		})
	}
}

func (s *suiteUsers) TestHandlerFindAllOrdersUser() {
	userId := "2d272252-7b5d-4f50-85ee-e578e3826510"

	orders := &[]model.Order{
		{
			ID:           "e156b109-3d44-4116-b039-7ab6129b6e84",
			UserId:       "883a840d-0bc5-4182-a0cf-9d4f2fcd5364",
			PaymentId:    "2725898e-cf67-491b-9df3-87198721ad96",
			TotalPayment: 200000,
			TotalQty:     3,
			TotalHour:    5,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	s.mocking.Mock.On("FindAllOrdersUser", userId).Return(orders, nil)

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Header             map[string]string
		Body               map[string]string
		HasReturnBody      bool
		ExpectedResult     map[string]interface{}
	}{
		{
			Name:               "success get user orders",
			ExpectedStatusCode: http.StatusOK,
			Method:             "GET",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success get all user orders",
				"data": map[string]interface{}{
					"orders": orders,
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/users", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id/orders")
			ctx.SetParamNames("id")
			ctx.SetParamValues(userId)

			err := s.handler.HandlerFindAllOrdersUser(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				data := resp["data"].(map[string]interface{})
				orders := data["orders"].([]interface{})

				expectRes := v.ExpectedResult["data"].(map[string]interface{})
				expectOrders := expectRes["orders"].(*[]model.Order)

				s.NotNil(orders)
				s.NotNil(expectOrders)
			}
		})
	}
}

func (s *suiteUsers) TestHandlerFindByIdOrderUser() {
	userId := "2d272252-7b5d-4f50-85ee-e578e3826510"
	orderId := "e156b109-3d44-4116-b039-7ab6129b6e84"

	order := &model.Order{
		ID:           "e156b109-3d44-4116-b039-7ab6129b6e84",
		UserId:       "883a840d-0bc5-4182-a0cf-9d4f2fcd5364",
		PaymentId:    "2725898e-cf67-491b-9df3-87198721ad96",
		TotalPayment: 200000,
		TotalQty:     3,
		TotalHour:    5,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	s.mocking.Mock.On("FindByIdOrderUser", orderId).Return(order, nil)

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Header             map[string]string
		Body               map[string]string
		HasReturnBody      bool
		ExpectedResult     map[string]interface{}
	}{
		{
			Name:               "success get user order by id",
			ExpectedStatusCode: http.StatusOK,
			Method:             "GET",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success get all user order by id",
				"data": map[string]interface{}{
					"order": order,
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/users", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id/orders/:orderId")
			ctx.SetParamNames("id", "orderId")
			ctx.SetParamValues(userId, orderId)

			err := s.handler.HandlerFindByIdOrderUser(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				data := resp["data"].(map[string]interface{})
				order := data["order"].(map[string]interface{})

				expectRes := v.ExpectedResult["data"].(map[string]interface{})
				expectOrder := expectRes["order"].(*model.Order)

				s.Equal(expectOrder.ID, order["id"])
				s.Equal(float64(expectOrder.TotalPayment), order["total_payment"])
				s.Equal(float64(expectOrder.TotalQty), order["total_qty"])
				s.Equal(float64(expectOrder.TotalHour), order["total_hour"])
			}
		})
	}
}

func (s *suiteUsers) TestHandlerUpdateUser() {
	userId := "2d272252-7b5d-4f50-85ee-e578e3826510"

	userDTO := dto.UserDTO{
		Fullname: "Kujo Jotaro",
		Phone:    "08654433",
		Address:  "Jl Sengon",
		Role:     "customer",
	}

	s.mocking.Mock.On("UpdateUser", userId, userDTO).Return(nil)

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
			Name:               "success update user",
			ExpectedStatusCode: http.StatusOK,
			Method:             "POST",
			Header: map[string]string{
				"Content-Type": "application/json",
			},
			Body: map[string]interface{}{
				"fullname": "Kujo Jotaro",
				"phone":    "08654433",
				"address":  "Jl Sengon",
				"role":     "customer",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success update user",
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
				"fullname": "Kujo Jotaro",
				"phone":    "08654433",
				"address":  "Jl Sengon",
				"role":     "customer",
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
			r := httptest.NewRequest(v.Method, "/users", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(userId)
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])

			err := s.handler.HandlerUpdateUser(ctx)
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

func (s *suiteUsers) TestHandlerDeleteUser() {
	userId := "2d272252-7b5d-4f50-85ee-e578e3826510"

	s.mocking.Mock.On("DeleteUser", userId).Return(nil)

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		ExpectedResult     map[string]interface{}
	}{
		{
			Name:               "success delete user",
			ExpectedStatusCode: http.StatusOK,
			Method:             "DELETE",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success delete user",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/users", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(userId)

			err := s.handler.HandlerDeleteUser(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				s.Equal(v.ExpectedResult["message"], resp["message"])
			}
		})
	}
}

func (s *suiteUsers) TearDownSuite() {
	s.mocking = nil
}

func TestSuiteUserController(t *testing.T) {
	suite.Run(t, new(suiteUsers))
}
