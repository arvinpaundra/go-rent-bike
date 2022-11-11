package rest_http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	usecasemock "github.com/arvinpaundra/go-rent-bike/internal/usecase/mock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type suiteRenter struct {
	suite.Suite
	handler *RenterController
	mocking *usecasemock.RenterUsecaseMock
}

func (s *suiteRenter) SetupSuite() {
	mock := &usecasemock.RenterUsecaseMock{}
	s.mocking = mock

	s.handler = &RenterController{
		renterUsecase: s.mocking,
	}
}

func (s *suiteRenter) TestHandlerCreateRenter() {

}

func (s *suiteRenter) TestHandlerCreateReportRenter() {
	renterId := "e1c74c4a-2d34-4ba3-8742-73b0130afae5"
	reportDTO := dto.ReportDTO{
		UserId:     "cc3c5658-2e2a-4f83-bf0a-5adeab34f036",
		TitleIssue: "Title Issue",
		BodyIssue:  "Body issue.",
	}

	s.mocking.Mock.On("CreateReportRenter", renterId, reportDTO).Return(nil)

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
			Name:               "success create new renter",
			ExpectedStatusCode: http.StatusCreated,
			Method:             "POST",
			Header: map[string]string{
				"Content-Type": "application/json",
			},
			Body: map[string]interface{}{
				"user_id":     "cc3c5658-2e2a-4f83-bf0a-5adeab34f036",
				"title_issue": "Title Issue",
				"body_issue":  "Body issue.",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success report renter",
				"data":    nil,
			},
		},
		{
			Name:               "failed wrong content type",
			ExpectedStatusCode: http.StatusBadRequest,
			Method:             "POST",
			Header: map[string]string{
				"Content-Type": "text/plain",
			},
			Body: map[string]interface{}{
				"user_id":     "cc3c5658-2e2a-4f83-bf0a-5adeab34f036",
				"title_issue": "Title Issue",
				"body_issue":  "Body issue.",
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
			r := httptest.NewRequest(v.Method, "/renters", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])
			ctx.SetPath("/:id/reports")
			ctx.SetParamNames("id")
			ctx.SetParamValues(renterId)

			err := s.handler.HandlerCreateReportRenter(ctx)
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

func (s *suiteRenter) TestHandlerFindAllRenters() {
	renters := &[]model.Renter{
		{
			ID:          "e1c74c4a-2d34-4ba3-8742-73b0130afae5",
			UserId:      "cc3c5658-2e2a-4f83-bf0a-5adeab34f036",
			RentName:    "Twins' Brother Bike Rental",
			RentAddress: "Jl Morioh",
			Description: "Full with description texts",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	s.mocking.Mock.On("FindAllRenters", "").Return(renters, nil)

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
			Name:               "success get renters",
			ExpectedStatusCode: http.StatusOK,
			Method:             "GET",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success get all renters",
				"data": map[string]interface{}{
					"renters": renters,
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/renters", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			err := s.handler.HandlerFindAllRenters(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				data := resp["data"].(map[string]interface{})
				renters := data["renters"].([]interface{})

				s.NotNil(renters)
			}
		})
	}
}

func (s *suiteRenter) TestHandlerFindRenterById() {
	renterId := "e1c74c4a-2d34-4ba3-8742-73b0130afae5"
	renter := &model.Renter{
		ID:          "e1c74c4a-2d34-4ba3-8742-73b0130afae5",
		UserId:      "cc3c5658-2e2a-4f83-bf0a-5adeab34f036",
		RentName:    "Twins' Brother Bike Rental",
		RentAddress: "Jl Morioh",
		Description: "Full with description texts",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.mocking.Mock.On("FindByIdRenter", "e1c74c4a-2d34-4ba3-8742-73b0130afae5").Return(renter, nil)

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
			Name:               "success get renter by id",
			ExpectedStatusCode: http.StatusOK,
			Method:             "GET",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success get all renters",
				"data": map[string]interface{}{
					"renter": renter,
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/renters", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(renterId)

			err := s.handler.HandlerFindRenterById(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				data := resp["data"].(map[string]interface{})
				renter := data["renter"]

				s.NotNil(renter)
			}
		})
	}
}

func (s *suiteRenter) TestHandlerFindAllRenterReports() {
	renterId := "e1c74c4a-2d34-4ba3-8742-73b0130afae5"
	reports := &[]model.Report{
		{
			ID:         "2ce20f26-32e6-4b22-ad0e-49ee2f11a0c0",
			RenterId:   "e1c74c4a-2d34-4ba3-8742-73b0130afae5",
			UserId:     "cc3c5658-2e2a-4f83-bf0a-5adeab34f036",
			TitleIssue: "Title Issue",
			BodyIssue:  "Body issue.",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	s.mocking.Mock.On("FindAllRenterReports", "e1c74c4a-2d34-4ba3-8742-73b0130afae5").Return(reports, nil)

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
			Name:               "success get renter reports",
			ExpectedStatusCode: http.StatusOK,
			Method:             "GET",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success get all reports",
				"data": map[string]interface{}{
					"reports": reports,
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/renters", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id/reports")
			ctx.SetParamNames("id")
			ctx.SetParamValues(renterId)

			err := s.handler.HandlerFindAllRenterReports(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				data := resp["data"].(map[string]interface{})
				reports := data["reports"].([]interface{})

				s.NotNil(reports)
			}
		})
	}
}

func (s *suiteRenter) TestHandlerUpdateRenter() {
	renterId := "e1c74c4a-2d34-4ba3-8742-73b0130afae5"
	renterDTO := dto.RenterDTO{
		RentName:    "Updated Rental Name",
		RentAddress: "Updated Rental Address",
		Description: "Updated descriptions.",
	}

	s.mocking.Mock.On("UpdateRenter", renterId, renterDTO).Return(nil)

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
			Name:               "success update renter",
			ExpectedStatusCode: http.StatusOK,
			Method:             "PUT",
			Header: map[string]string{
				"Content-Type": "application/json",
			},
			Body: map[string]interface{}{
				"rent_name":    "Updated Rental Name",
				"rent_address": "Updated Rental Address",
				"description":  "Updated descriptions.",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success update renter",
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
				"rent_name":    "Updated Rental Name",
				"rent_address": "Updated Rental Address",
				"description":  "Updated descriptions.",
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
			r := httptest.NewRequest(v.Method, "/renters", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(renterId)

			err := s.handler.HandlerUpdateRenter(ctx)
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

func (s *suiteRenter) TestHandlerDeleteRenter() {
	renterId := "e1c74c4a-2d34-4ba3-8742-73b0130afae5"

	s.mocking.Mock.On("DeleteRenter", renterId).Return(nil)

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		ExpectedResult     map[string]interface{}
	}{
		{
			Name:               "success delete renter",
			ExpectedStatusCode: http.StatusOK,
			Method:             "DELETE",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success delete renter",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/renters", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(renterId)

			err := s.handler.HandlerDeleteRenter(ctx)
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

func (s *suiteRenter) TearDownSuite() {
	s.mocking = nil
}

func TestSuiteRenter(t *testing.T) {
	suite.Run(t, new(suiteRenter))
}
