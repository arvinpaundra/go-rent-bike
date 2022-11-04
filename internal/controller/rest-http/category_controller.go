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

type CategoryController struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewCategoryController(categoryUsecase usecase.CategoryUsecase) *CategoryController {
	return &CategoryController{categoryUsecase}
}

func (h *CategoryController) HandlerCreateCategory(c echo.Context) error {
	categoryDTO := dto.CategoryDTO{}

	if err := c.Bind(&categoryDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "name field is required",
			"data":    nil,
		})
	}

	err := h.categoryUsecase.CreateCategory(categoryDTO)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "sucess create category",
		"data":    nil,
	})
}

func (h *CategoryController) HandlerFindAllCategories(c echo.Context) error {
	categories, err := h.categoryUsecase.FindAllCategories()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success get all categories",
		"data": map[string]*[]model.Category{
			"categories": categories,
		},
	})
}

func (h *CategoryController) HandlerFindCategoryById(c echo.Context) error {
	categoryId := c.Param("id")

	category, err := h.categoryUsecase.FindByIdCategory(categoryId)

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
		"message": "success get category by id",
		"data": map[string]*model.Category{
			"category": category,
		},
	})
}

func (h *CategoryController) HandlerUpdateCategory(c echo.Context) error {
	categoryId := c.Param("id")
	categoryDTO := dto.CategoryDTO{}

	if err := c.Bind(&categoryDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "name field is required",
			"data":    nil,
		})
	}

	err := h.categoryUsecase.UpdateCategory(categoryId, categoryDTO)

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
		"message": "success update category by id",
		"data":    nil,
	})
}

func (h *CategoryController) HandlerDeleteCategory(c echo.Context) error {
	categoryId := c.Param("id")

	rowAffected, err := h.categoryUsecase.DeleteCategory(categoryId)

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
		"status":  "error",
		"message": "success delete category by id",
		"data": map[string]uint{
			"row_affected": rowAffected,
		},
	})
}
