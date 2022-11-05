package rest_http

import (
	"errors"
	"net/http"

	"github.com/arvinpaundra/go-rent-bike/internal"
	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/usecase"
	"github.com/labstack/echo/v4"
)

type OrderController struct {
	orderUsecase usecase.OrderUsecase
}

func NewOrderController(orderUsecase usecase.OrderUsecase) *OrderController {
	return &OrderController{orderUsecase}
}

func (h *OrderController) HandlerCreateNewOrder(c echo.Context) error {
	orderDTO := dto.OrderDTO{}

	if err := c.Bind(&orderDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "fill all required fields",
			"data":    nil,
		})
	}

	data, err := h.orderUsecase.CreateOrder(orderDTO)

	if err != nil {
		if errors.Is(err, internal.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"status":  "error",
				"message": "record not found",
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
		"message": "success create an order",
		"data":    data,
	})
}

func (h *OrderController) HandlerReturnBike(c echo.Context) error {
	orderId := c.Param("id")

	err := h.orderUsecase.UpdateRentStatus(orderId)

	if err != nil {
		if errors.Is(err, internal.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"status":  "error",
				"message": "order not found",
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
		"message": "success return bike",
		"data":    nil,
	})
}
