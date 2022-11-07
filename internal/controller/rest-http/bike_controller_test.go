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

type suiteBikes struct {
	suite.Suite
	handler *BikeController
	mocking *mocking.BikeUsecaseMock
}

func (s *suiteBikes) SetupSuite() {
	mock := &mocking.BikeUsecaseMock{}
	s.mocking = mock

	s.handler = &BikeController{
		bikeUsecase: s.mocking,
	}
}

func (s *suiteBikes) TestHandlerAddNewBike() {
	bikeDTO := dto.BikeDTO{
		RenterId:     "478b3f5e-284e-440c-8c0f-af4f94c70d87",
		CategoryId:   "8cfb93e7-a1f4-47e2-bb5b-ffea24761322",
		Name:         "Sample Mountain Bike",
		PricePerHour: 12000,
		Condition:    "Good",
		Description:  "Description section.",
		IsAvailable:  "1",
	}

	s.mocking.Mock.On("CreateNewBike", bikeDTO).Return(nil)

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
			Name:               "success add bike",
			ExpectedStatusCode: http.StatusCreated,
			Method:             "POST",
			Header: map[string]string{
				"Content-Type": "application/json",
			},
			Body: map[string]interface{}{
				"renter_id":      "478b3f5e-284e-440c-8c0f-af4f94c70d87",
				"category_id":    "8cfb93e7-a1f4-47e2-bb5b-ffea24761322",
				"name":           "Sample Mountain Bike",
				"price_per_hour": float32(12000),
				"condition":      "Good",
				"description":    "Description section.",
				"is_available":   "1",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success add new bike",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/bikes", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])

			err := s.handler.HandlerAddNewBike(ctx)
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

func (s *suiteBikes) TestHandlerCreateNewBikeReview() {
	bikeId := "e4a37040-5fde-4921-b60d-c1628452b4b2"

	reviewDTO := dto.ReviewDTO{
		UserId:      "8ad58074-228c-430d-918e-01105cc084fa",
		Rating:      5,
		Description: "What a good bike.",
	}

	s.mocking.Mock.On("CreateNewBikeReview", bikeId, reviewDTO).Return(nil)

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
			Name:               "success add bike review",
			ExpectedStatusCode: http.StatusCreated,
			Method:             "POST",
			Header: map[string]string{
				"Content-Type": "application/json",
			},
			Body: map[string]interface{}{
				"user_id":     "8ad58074-228c-430d-918e-01105cc084fa",
				"rating":      5,
				"description": "What a good bike.",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success create new bike review",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/bikes", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])
			ctx.SetPath("/:id/reviews")
			ctx.SetParamNames("id")
			ctx.SetParamValues(bikeId)

			err := s.handler.HandlerCreateNewBikeReview(ctx)
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

func (s *suiteBikes) TestHandlerFindAllBikes() {
	bikes := &[]model.Bike{
		{
			ID:           "281b273f-49bd-482e-b939-71636e2be32d",
			RenterId:     "8ad58074-228c-430d-918e-01105cc084fa",
			CategoryId:   "e4a37040-5fde-4921-b60d-c1628452b4b2",
			Name:         "Sample Mountain Bike",
			PricePerHour: float32(12000),
			Condition:    "Good",
			Description:  "Description section",
			IsAvailable:  "1",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	s.mocking.Mock.On("FindAllBikes", "").Return(bikes, nil)

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
			Name:               "success get bikes",
			ExpectedStatusCode: http.StatusOK,
			Method:             "GET",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success get all bikes",
				"data": map[string]interface{}{
					"bikes": bikes,
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/bikes", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			err := s.handler.HandlerFindAllBikes(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				data := resp["data"].(map[string]interface{})
				bikes := data["bikes"].([]interface{})

				expectRes := v.ExpectedResult["data"].(map[string]interface{})
				expectCategories := expectRes["bikes"].(*[]model.Bike)

				s.NotEmpty(bikes)
				s.NotEmpty(expectCategories)
			}
		})
	}
}

func (s *suiteBikes) TestHandlerFindByIdBike() {
	bikeId := "281b273f-49bd-482e-b939-71636e2be32d"

	bike := &model.Bike{
		ID:           "281b273f-49bd-482e-b939-71636e2be32d",
		RenterId:     "8ad58074-228c-430d-918e-01105cc084fa",
		CategoryId:   "e4a37040-5fde-4921-b60d-c1628452b4b2",
		Name:         "Sample Mountain Bike",
		PricePerHour: float32(12000),
		Condition:    "Good",
		Description:  "Description section",
		IsAvailable:  "1",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	s.mocking.Mock.On("FindByIdBike", bikeId).Return(bike, nil)

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
			Name:               "success get bike by id",
			ExpectedStatusCode: http.StatusOK,
			Method:             "GET",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success get bike by id",
				"data": map[string]interface{}{
					"bike": bike,
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/bikes", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(bikeId)

			err := s.handler.HandlerFindByIdBike(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				data := resp["data"].(map[string]interface{})
				bike := data["bike"].(map[string]interface{})

				expectRes := v.ExpectedResult["data"].(map[string]interface{})
				expectBike := expectRes["bike"].(*model.Bike)

				s.Equal(expectBike.ID, bike["id"])
				s.Equal(expectBike.CategoryId, bike["category_id"])
				s.Equal(expectBike.RenterId, bike["renter_id"])
				s.Equal(expectBike.Name, bike["name"])
				s.Equal(expectBike.Condition, bike["condition"])
				s.Equal(expectBike.PricePerHour, float32(bike["price_per_hour"].(float64)))
				s.Equal(expectBike.Description, bike["description"])
				s.Equal(expectBike.IsAvailable, bike["is_available"])
			}
		})
	}
}

func (s *suiteBikes) TestHandlerFindBikesByRenter() {
	renterId := "8ad58074-228c-430d-918e-01105cc084fa"

	bikes := &[]model.Bike{
		{
			ID:           "281b273f-49bd-482e-b939-71636e2be32d",
			RenterId:     "8ad58074-228c-430d-918e-01105cc084fa",
			CategoryId:   "e4a37040-5fde-4921-b60d-c1628452b4b2",
			Name:         "Sample Mountain Bike",
			PricePerHour: float32(12000),
			Condition:    "Good",
			Description:  "Description section",
			IsAvailable:  "1",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	s.mocking.Mock.On("FindBikesByRenter", renterId).Return(bikes, nil)

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
			Name:               "success get bikes by renter",
			ExpectedStatusCode: http.StatusOK,
			Method:             "GET",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success get bikes by renter",
				"data": map[string]interface{}{
					"bikes": bikes,
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/bikes", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/renters/:renterId")
			ctx.SetParamNames("renterId")
			ctx.SetParamValues(renterId)

			err := s.handler.HandlerFindBikesByRenter(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				data := resp["data"].(map[string]interface{})
				bikes := data["bikes"].([]interface{})

				expectRes := v.ExpectedResult["data"].(map[string]interface{})
				expectCategories := expectRes["bikes"].(*[]model.Bike)

				s.NotEmpty(bikes)
				s.NotEmpty(expectCategories)
			}
		})
	}
}

func (s *suiteBikes) TestHandlerFindBikesByCategory() {
	categoryId := "8ad58074-228c-430d-918e-01105cc084fa"

	bikes := &[]model.Bike{
		{
			ID:           "281b273f-49bd-482e-b939-71636e2be32d",
			RenterId:     "8ad58074-228c-430d-918e-01105cc084fa",
			CategoryId:   "e4a37040-5fde-4921-b60d-c1628452b4b2",
			Name:         "Sample Mountain Bike",
			PricePerHour: float32(12000),
			Condition:    "Good",
			Description:  "Description section",
			IsAvailable:  "1",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	s.mocking.Mock.On("FindBikesByCategory", categoryId).Return(bikes, nil)

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
			Name:               "success get bikes by category",
			ExpectedStatusCode: http.StatusOK,
			Method:             "GET",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success get bikes by category",
				"data": map[string]interface{}{
					"bikes": bikes,
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/bikes", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/categories/:categoryId")
			ctx.SetParamNames("categoryId")
			ctx.SetParamValues(categoryId)

			err := s.handler.HandlerFindBikesByCategory(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				data := resp["data"].(map[string]interface{})
				bikes := data["bikes"].([]interface{})

				expectRes := v.ExpectedResult["data"].(map[string]interface{})
				expectCategories := expectRes["bikes"].(*[]model.Bike)

				s.NotEmpty(bikes)
				s.NotEmpty(expectCategories)
			}
		})
	}
}

func (s *suiteBikes) TestHandlerUpdateBike() {
	bikeId := "8ad58074-228c-430d-918e-01105cc084fa"

	bikeDTO := dto.BikeDTO{
		RenterId:     "478b3f5e-284e-440c-8c0f-af4f94c70d87",
		CategoryId:   "8cfb93e7-a1f4-47e2-bb5b-ffea24761322",
		Name:         "Sample Mountain Bike",
		PricePerHour: 12000,
		Condition:    "Good",
		Description:  "Description section.",
		IsAvailable:  "1",
	}

	s.mocking.Mock.On("UpdateBike", bikeId, bikeDTO).Return(nil)

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
			Name:               "success update bike",
			ExpectedStatusCode: http.StatusOK,
			Method:             "PUT",
			Header: map[string]string{
				"Content-Type": "application/json",
			},
			Body: map[string]interface{}{
				"renter_id":      "478b3f5e-284e-440c-8c0f-af4f94c70d87",
				"category_id":    "8cfb93e7-a1f4-47e2-bb5b-ffea24761322",
				"name":           "Sample Mountain Bike",
				"price_per_hour": float32(12000),
				"condition":      "Good",
				"description":    "Description section.",
				"is_available":   "1",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success update bike by id",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/bikes", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(bikeId)
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])

			err := s.handler.HandlerUpdateBike(ctx)
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

func (s *suiteBikes) TestHandlerDeleteBike() {
	bikeId := "8ad58074-228c-430d-918e-01105cc084fa"

	s.mocking.Mock.On("DeleteBike", bikeId).Return(nil)

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		ExpectedResult     map[string]interface{}
	}{
		{
			Name:               "success delete bike",
			ExpectedStatusCode: http.StatusOK,
			Method:             "DELETE",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success delete bike by id",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/bikes", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(bikeId)

			err := s.handler.HandlerDeleteBike(ctx)
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

func (s *suiteBikes) TearDownSuite() {
	s.mocking = nil
}

func TestSuiteBike(t *testing.T) {
	suite.Run(t, new(suiteBikes))
}
