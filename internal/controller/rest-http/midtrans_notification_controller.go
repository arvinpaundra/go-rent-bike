package rest_http

import (
	"errors"
	"github.com/arvinpaundra/go-rent-bike/pkg"
	"net/http"

	"github.com/arvinpaundra/go-rent-bike/internal/usecase"
	"github.com/labstack/echo/v4"
)

type MidtransNotificationController struct {
	paymentGatewayUsecase usecase.PaymentGatewayUsecase
}

func (h *MidtransNotificationController) HandlerNotification(c echo.Context) error {
	var notificationPayloads map[string]interface{}

	if err := c.Bind(&notificationPayloads); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "fill all required fields",
			"data":    nil,
		})
	}

	orderId := notificationPayloads["order_id"].(string)

	err := h.paymentGatewayUsecase.MidtransNotification(orderId)

	if err != nil {
		if errors.Is(err, pkg.ErrRecordNotFound) {
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
		"message": "success update transaction status",
		"data":    nil,
	})
}

func NewMidtransNotificationController(paymentGatewayUsecase usecase.PaymentGatewayUsecase) *MidtransNotificationController {
	return &MidtransNotificationController{paymentGatewayUsecase}
}
