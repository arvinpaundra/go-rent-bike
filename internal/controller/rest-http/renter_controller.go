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

type RenterController struct {
	renterUsecase usecase.RenterUsecase
}

func NewRenterController(renterUsecase usecase.RenterUsecase) *RenterController {
	return &RenterController{renterUsecase}
}

func (r RenterController) HandlerCreateRenter(c echo.Context) error {
	renterDTO := dto.RenterDTO{}

	if err := c.Bind(&renterDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "fill all required fields (renter field)",
			"data":    nil,
		})
	}

	err := r.renterUsecase.CreateRenter(renterDTO)

	if err != nil {
		if errors.Is(err, internal.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"status":  "error",
				"message": "user not found",
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
		"message": "register success",
		"data":    nil,
	})
}

func (r RenterController) HandlerCreateReportRenter(c echo.Context) error {
	renterId := c.Param("id")
	reportDTO := dto.ReportDTO{}

	if err := c.Bind(&reportDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "fill all required fields",
			"data":    nil,
		})
	}

	err := r.renterUsecase.CreateReportRenter(renterId, reportDTO)

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

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "success report renter",
		"data":    nil,
	})
}

func (r RenterController) HandlerFindAllRenters(c echo.Context) error {
	search := c.QueryParam("rental_name")

	renters, err := r.renterUsecase.FindAllRenters(search)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success get all renters",
		"data": map[string]*[]model.Renter{
			"renters": renters,
		},
	})
}

func (r RenterController) HandlerFindRenterById(c echo.Context) error {
	renterId := c.Param("id")

	renter, err := r.renterUsecase.FindByIdRenter(renterId)

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
		"message": "success get renter by id",
		"data": map[string]*model.Renter{
			"renter": renter,
		},
	})
}

func (r RenterController) HandlerFindAllRenterReports(c echo.Context) error {
	renterId := c.Param("id")

	reports, err := r.renterUsecase.FindAllRenterReports(renterId)

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

	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"status":  "success",
		"message": "success get all reports",
		"data": map[string]*[]model.Report{
			"reports": reports,
		},
	})
}

func (r RenterController) HandlerUpdateRenter(c echo.Context) error {
	renterId := c.Param("id")
	renterDTO := dto.RenterDTO{}

	if err := c.Bind(&renterDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	err := r.renterUsecase.UpdateRenter(renterId, renterDTO)

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

	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"status":  "success",
		"message": "success update renter",
		"data":    nil,
	})
}

func (r RenterController) HandlerDeleteRenter(c echo.Context) error {
	renterId := c.Param("id")

	rowAffected, err := r.renterUsecase.DeleteRenter(renterId)

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
		"message": "success delete renter",
		"data": map[string]uint{
			"row_affected": rowAffected,
		},
	})
}
