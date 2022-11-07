package rest_http

import (
	"bytes"
	"encoding/json"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	mocking "github.com/arvinpaundra/go-rent-bike/internal/usecase/mock"
	"github.com/stretchr/testify/suite"
)

type suiteCategory struct {
	suite.Suite
	handler *CategoryController
	mocking *mocking.CategoryUsecaseMock
}

func (s *suiteCategory) SetupSuite() {
	mock := &mocking.CategoryUsecaseMock{}
	s.mocking = mock

	s.handler = &CategoryController{
		categoryUsecase: s.mocking,
	}
}

func (s *suiteCategory) TestHandlerCreateCategory() {
	categoryDTO := dto.CategoryDTO{
		Name: "Mountain Bike",
	}

	s.mocking.Mock.On("CreateCategory", categoryDTO).Return(nil)

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
			Name:               "success create category",
			ExpectedStatusCode: http.StatusCreated,
			Method:             "POST",
			Header: map[string]string{
				"Content-Type": "application/json",
			},
			Body: map[string]string{
				"name": "Mountain Bike",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success create category",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/categories", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])

			err := s.handler.HandlerCreateCategory(ctx)
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

func (s *suiteCategory) TestHandlerFindAllCategories() {
	categories := &[]model.Category{
		{
			ID:        "8ad58074-228c-430d-918e-01105cc084fa",
			Name:      "Mountain Bike",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "8ad58074-228c-430d-918e-01105cc084fa",
			Name:      "Fixie",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	s.mocking.Mock.On("FindAllCategories").Return(categories, nil)

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
			Name:               "success get categories",
			ExpectedStatusCode: http.StatusOK,
			Method:             "GET",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success get all categories",
				"data": map[string]interface{}{
					"categories": categories,
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/categories", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			err := s.handler.HandlerFindAllCategories(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				data := resp["data"].(map[string]interface{})
				categories := data["categories"].([]interface{})

				expectRes := v.ExpectedResult["data"].(map[string]interface{})
				expectCategories := expectRes["categories"].(*[]model.Category)

				s.NotEmpty(categories)
				s.NotEmpty(expectCategories)
			}
		})
	}
}

func (s *suiteCategory) TestHandlerFindCategoryById() {
	categoryId := "8ad58074-228c-430d-918e-01105cc084fa"

	category := &model.Category{
		ID:        "8ad58074-228c-430d-918e-01105cc084fa",
		Name:      "Mountain Bike",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.mocking.Mock.On("FindByIdCategory", categoryId).Return(category, nil)

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
			Name:               "success get categories",
			ExpectedStatusCode: http.StatusOK,
			Method:             "GET",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success get category by id",
				"data": map[string]interface{}{
					"category": category,
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/categories", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(categoryId)

			err := s.handler.HandlerFindCategoryById(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}
				err := json.NewDecoder(w.Result().Body).Decode(&resp)
				s.NoError(err)

				data := resp["data"].(map[string]interface{})
				category := data["category"].(map[string]interface{})

				expectRes := v.ExpectedResult["data"].(map[string]interface{})
				expectCategory := expectRes["category"].(*model.Category)

				s.Equal(expectCategory.ID, category["id"])
				s.Equal(expectCategory.Name, category["name"])
			}
		})
	}
}

func (s *suiteCategory) TestHandlerUpdateCategory() {
	categoryId := "8ad58074-228c-430d-918e-01105cc084fa"

	categoryDTO := dto.CategoryDTO{
		Name: "Fixie",
	}

	s.mocking.Mock.On("UpdateCategory", categoryId, categoryDTO).Return(nil)

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
			Name:               "success update category",
			ExpectedStatusCode: http.StatusOK,
			Method:             "PUT",
			Header: map[string]string{
				"Content-Type": "application/json",
			},
			Body: map[string]string{
				"name": "Fixie",
			},
			HasReturnBody: true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success update category by id",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/categories", bytes.NewReader(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(categoryId)
			ctx.Request().Header.Set("Content-Type", v.Header["Content-Type"])

			err := s.handler.HandlerUpdateCategory(ctx)
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

func (s *suiteCategory) TestHandlerDeleteCategory() {
	categoryId := "8ad58074-228c-430d-918e-01105cc084fa"

	s.mocking.Mock.On("DeleteCategory", categoryId).Return(uint(1), nil)

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		HasReturnBody      bool
		ExpectedResult     map[string]interface{}
	}{
		{
			Name:               "success delete category",
			ExpectedStatusCode: http.StatusOK,
			Method:             "DELETE",
			HasReturnBody:      true,
			ExpectedResult: map[string]interface{}{
				"status":  "success",
				"message": "success delete category by id",
				"data":    nil,
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/categories", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(categoryId)

			err := s.handler.HandlerDeleteCategory(ctx)
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

func (s *suiteCategory) TearDownSuite() {
	s.mocking = nil
}

func TestSuiteCategory(t *testing.T) {
	suite.Run(t, new(suiteCategory))
}
