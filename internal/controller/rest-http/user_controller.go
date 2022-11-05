package rest_http

import (
	"errors"
	"net/http"

	"github.com/arvinpaundra/go-rent-bike/helper"
	"github.com/arvinpaundra/go-rent-bike/internal"
	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/usecase"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}

func (u *UserController) HandlerRegister(c echo.Context) error {
	userDTO := dto.UserDTO{}

	if err := c.Bind(&userDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "fill all required field",
			"data":    nil,
		})
	}

	if userDTO.Role != "customer" && userDTO.Role != "renter" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "only allowed roles renter and customer",
			"data":    nil,
		})
	}

	err := u.userUsecase.RegisterUser(userDTO)

	if err != nil {
		if errors.Is(err, internal.ErrDataAlreadyExist) {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"status":  "error",
				"message": "email already exist",
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

func (u *UserController) HandlerLogin(c echo.Context) error {
	loginDTO := struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}{}

	if err := c.Bind(&loginDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "email and password is required",
			"data":    nil,
		})
	}

	user, err := u.userUsecase.LoginUser(loginDTO.Email, loginDTO.Password)

	if err != nil {
		if errors.Is(err, internal.ErrRecordNotFound) {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"status":  "error",
				"message": "wrong email or password",
				"data":    nil,
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	token, _ := helper.CreateToken(user.ID, user.Role)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "login success",
		"data": map[string]string{
			"token": token,
		},
	})
}

func (u *UserController) HandlerFindAllUsers(c echo.Context) error {
	users, err := u.userUsecase.FindAllUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success get all users",
		"data": map[string]*[]model.User{
			"users": users,
		},
	})
}

func (u *UserController) HandlerFindUserById(c echo.Context) error {
	userId := c.Param("id")

	user, err := u.userUsecase.FindByIdUser(userId)

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
		"message": "success get user by id",
		"data": map[string]*model.User{
			"user": user,
		},
	})
}

func (u *UserController) HandlerFindAllUserHistories(c echo.Context) error {
	userId := c.Param("id")

	histories, err := u.userUsecase.FindAllUserHistories(userId)

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

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success get user histories",
		"data": map[string]*[]model.History{
			"histories": histories,
		},
	})
}

func (h *UserController) HandlerFindAllOrdersUser(c echo.Context) error {
	userId := c.Param("id")

	orders, err := h.userUsecase.FindAllOrdersUser(userId)

	if err != nil {
		if errors.Is(err, internal.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"status":  "error",
				"message": "customer not found",
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
		"message": "success get all user orders",
		"data":    orders,
	})
}

func (h *UserController) HandlerFindByIdOrderUser(c echo.Context) error {
	_ = c.Param("id")
	orderId := c.Param("orderId")

	order, err := h.userUsecase.FindByIdOrderUser(orderId)

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
		"message": "success get order by id order and user",
		"data":    order,
	})
}

func (u *UserController) HandlerUpdateUser(c echo.Context) error {
	userId := c.Param("id")
	userDTO := dto.UserDTO{}

	if err := c.Bind(&userDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "fill all required field",
			"data":    nil,
		})
	}

	err := u.userUsecase.UpdateUser(userId, userDTO)

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

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success update user",
		"data":    nil,
	})
}

func (u *UserController) HandlerDeleteUser(c echo.Context) error {
	userId := c.Param("id")

	rowAffected, err := u.userUsecase.DeleteUser(userId)

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

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success delete user",
		"data": map[string]uint{
			"row_affected": rowAffected,
		},
	})
}
