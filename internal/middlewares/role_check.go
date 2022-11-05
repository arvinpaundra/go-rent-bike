package mddlwrs

import (
	"net/http"

	"github.com/arvinpaundra/go-rent-bike/helper"
	"github.com/labstack/echo/v4"
)

func CheckIsRenter(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenExtracted, err := helper.ExtractToken(c)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"status":  "error",
				"message": err.Error(),
				"data":    nil,
			})
		}

		payloads := tokenExtracted.(map[string]string)
		role := payloads["role"]

		if role != "renter" {
			return c.JSON(http.StatusForbidden, map[string]interface{}{
				"status":  "error",
				"message": "user role must be renter",
				"data":    nil,
			})
		}

		return next(c)
	}
}
