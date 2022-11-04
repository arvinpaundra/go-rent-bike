package rest_http

import (
	"errors"
	"net/http"

	"github.com/arvinpaundra/go-rent-bike/internal"
	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/usecase"
	"github.com/labstack/echo/v4"
)

type BikeController struct {
	bikeUsecase usecase.BikeUsecase
}

func NewBikeController(bikeUsecase usecase.BikeUsecase) *BikeController {
	return &BikeController{bikeUsecase}
}

func (h *BikeController) HandlerAddNewBike(c echo.Context) error {
	bikeDTO := dto.BikeDTO{}

	if err := c.Bind(&bikeDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "fill all required fields",
			"data":    nil,
		})
	}

	err := h.bikeUsecase.CreateNewBike(bikeDTO)

	if err != nil {
		if errors.Is(err, internal.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"status":  "error",
				"message": "renter or category not found",
				"data":    nil,
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "success add new bike",
		"data":    nil,
	})
}

func (h *BikeController) HandlerFindAllBikes(c echo.Context) error {
	search := c.QueryParam("bike_name")

	bikes, err := h.bikeUsecase.FindAllBikes(search)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success get all bikes",
		"data": map[string]*[]model.Bike{
			"bikes": bikes,
		},
	})
}

func (h *BikeController) HandlerFindByIdBike(c echo.Context) error {
	bikeId := c.Param("id")

	bike, err := h.bikeUsecase.FindByIdBike(bikeId)

	if err != nil {
		if errors.Is(err, internal.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"status":  "error",
				"message": "bike not found",
				"data":    nil,
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success get bike by id",
		"data": map[string]*model.Bike{
			"bike": bike,
		},
	})
}

func (h *BikeController) HandlerFindBikesByRenter(c echo.Context) error {
	renterId := c.Param("renterId")

	data, err := h.bikeUsecase.FindBikesByRenter(renterId)

	if err != nil {
		if errors.Is(err, internal.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"status":  "error",
				"message": "renter not found",
				"data":    nil,
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success get bikes by renter",
		"data": map[string]interface{}{
			"rental_name": data["rental_name"].(string),
			"bikes":       data["bikes"].(*[]model.Bike),
		},
	})
}

func (h *BikeController) HandlerFindBikesByCategory(c echo.Context) error {
	categoryId := c.Param("categoryId")

	data, err := h.bikeUsecase.FindBikesByCategory(categoryId)

	if err != nil {
		if errors.Is(err, internal.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"status":  "error",
				"message": "category not found",
				"data":    nil,
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success get all bikes by category",
		"data": map[string]interface{}{
			"category": data["category"].(string),
			"bikes":    data["bikes"].(*[]model.Bike),
		},
	})
}

func (h *BikeController) HandlerUpdateBike(c echo.Context) error {
	bikeId := c.Param("id")
	bikeDTO := dto.BikeDTO{}

	if err := c.Bind(&bikeDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "fill all required fields",
			"data":    nil,
		})
	}

	err := h.bikeUsecase.UpdateBike(bikeId, bikeDTO)

	if err != nil {
		if errors.Is(err, internal.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"status":  "error",
				"message": "renter or category or bike are not found",
				"data":    nil,
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success update bike by id",
		"data":    nil,
	})
}

func (h *BikeController) HandlerDeleteBike(c echo.Context) error {
	bikeId := c.Param("id")

	err := h.bikeUsecase.DeleteBike(bikeId)

	if err != nil {
		if errors.Is(err, internal.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"status":  "error",
				"message": "bike not found",
				"data":    nil,
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success delete bike by id",
		"data":    nil,
	})
}
