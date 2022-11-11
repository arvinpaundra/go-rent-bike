package usecase

import (
	"testing"
	"time"

	"github.com/arvinpaundra/go-rent-bike/configs"
	"github.com/arvinpaundra/go-rent-bike/helper"

	"github.com/arvinpaundra/go-rent-bike/internal/dto"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userUsecaseTest = NewUserUsecase(
	&pkg.UserRepository,
	&pkg.HistoryRepository,
	&pkg.OrderRepository,
)

// TODO register user test
func TestUserUsecase_Register(t *testing.T) {
	userDTO := dto.UserDTO{
		Fullname: "Kujo Jotaro",
		Phone:    "08587654",
		Address:  "Jl Morioh",
		Role:     "customer",
		Email:    "jotarokujo@mail.com",
		Password: "123",
	}

	pkg.UserRepository.Mock.On("FindByEmail", userDTO.Email).Return(nil, nil)

	pkg.UserRepository.Mock.On("Create", mock.Anything).Return(nil)

	err := userUsecaseTest.RegisterUser(userDTO)

	assert.Nil(t, err)
}

// TODO login user test
func TestUserUsecase_LoginUser(t *testing.T) {
	configs.InitConfig()

	email := "arvin@mail.com"
	password := "123"

	hashedPassword, _ := helper.HashPassword(password)

	user := &model.User{
		ID:        "e694b986-cf9b-4b33-9147-3838e9014662",
		Fullname:  "Arvin Paundra Ardana",
		Phone:     "087654321",
		Address:   "Jl Morioh",
		Role:      "customer",
		Email:     "arvin@mail.com",
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	pkg.UserRepository.Mock.On("FindByEmail", mock.Anything).Return(user, nil)

	result, err := userUsecaseTest.LoginUser(email, password)

	assert.Nil(t, err)
	assert.NotEmpty(t, result)
}

func TestUserUsecase_FindAllUsers(t *testing.T) {
	users := &[]model.User{
		{
			Fullname:  "Arvin",
			Phone:     "085",
			Role:      "customer",
			Email:     "arvin@mail.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Fullname:  "ardana",
			Phone:     "098",
			Role:      "renter",
			Email:     "ardana@mail.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	pkg.UserRepository.Mock.On("FindAll").Return(users, nil)

	results, err := userUsecaseTest.FindAllUsers()

	assert.Nil(t, err)
	assert.NotNil(t, results)
}

func TestUserUsecase_FindByIdUser(t *testing.T) {
	userId := "eda51b42-36dc-4d81-8bec-90aee25790d1"

	user := &model.User{
		ID:        userId,
		Fullname:  "Arvin",
		Phone:     "085",
		Address:   "Jl Rinjani",
		Role:      "customer",
		Email:     "arvin@mail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	pkg.UserRepository.Mock.On("FindById", userId).Return(user, nil)

	result, err := userUsecaseTest.FindByIdUser(userId)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

// TODO find all user histories test
func TestUserUsecase_FindAllUserHistories(t *testing.T) {
	userId := "e694b986-cf9b-4b33-9147-3838e9014662"

	user := &model.User{
		ID:        userId,
		Fullname:  "Arvin",
		Phone:     "085",
		Address:   "Jl Rinjani",
		Role:      "customer",
		Email:     "arvin@mail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	pkg.UserRepository.Mock.On("FindById", userId).Return(user, nil)

	histories := &[]model.History{
		{
			ID:         "36dad17b-d850-44cc-9893-44667f7ce289",
			OrderId:    "81f3c938-c903-481b-aaf1-f9f254d1d126",
			RentStatus: "rented",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	pkg.HistoryRepository.Mock.On("FindAll", userId).Return(histories, nil)

	results, err := userUsecaseTest.FindAllUserHistories(userId)

	assert.Nil(t, err)
	assert.NotNil(t, results)

	assert.Equal(t, (*histories)[0].ID, (*results)[0].ID)
	assert.Equal(t, (*histories)[0].OrderId, (*results)[0].OrderId)
	assert.Equal(t, (*histories)[0].RentStatus, (*results)[0].RentStatus)
}

// TODO find all orders user test
func TestUserUsecase_FindAllOrdersUser(t *testing.T) {
	userId := "e694b986-cf9b-4b33-9147-3838e9014662"

	user := &model.User{
		ID:        userId,
		Fullname:  "Arvin",
		Phone:     "085",
		Address:   "Jl Rinjani",
		Role:      "customer",
		Email:     "arvin@mail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	pkg.UserRepository.Mock.On("FindById", userId).Return(user, nil)

	orders := &[]model.Order{
		{
			ID:           "193cca29-cb75-448b-b1f0-50001e86c17a",
			UserId:       "e694b986-cf9b-4b33-9147-3838e9014662",
			PaymentId:    "97f11c30-0fb6-42e0-9005-50315be0b905",
			TotalPayment: 200000,
			TotalQty:     3,
			TotalHour:    5,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	pkg.OrderRepository.Mock.On("FindAll", userId).Return(orders, nil)

	results, err := userUsecaseTest.FindAllOrdersUser(userId)

	assert.Nil(t, err)
	assert.NotNil(t, results)

	assert.Equal(t, (*orders)[0].ID, (*results)[0].ID)
	assert.Equal(t, (*orders)[0].UserId, (*results)[0].UserId)
	assert.Equal(t, (*orders)[0].PaymentId, (*results)[0].PaymentId)
	assert.Equal(t, (*orders)[0].TotalQty, (*results)[0].TotalQty)
	assert.Equal(t, (*orders)[0].TotalHour, (*results)[0].TotalHour)
	assert.Equal(t, (*orders)[0].TotalPayment, (*results)[0].TotalPayment)
}

// TODO find order user by id order test
func TestUserUsecase_FindByIdOrderUser(t *testing.T) {
	orderId := "193cca29-cb75-448b-b1f0-50001e86c17a"

	order := &model.Order{
		ID:           "193cca29-cb75-448b-b1f0-50001e86c17a",
		UserId:       "e694b986-cf9b-4b33-9147-3838e9014662",
		PaymentId:    "97f11c30-0fb6-42e0-9005-50315be0b905",
		TotalPayment: 200000,
		TotalQty:     3,
		TotalHour:    5,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	pkg.OrderRepository.Mock.On("FindById", orderId).Return(order, nil)

	result, err := userUsecaseTest.FindByIdOrderUser(orderId)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, order.ID, result.ID)
	assert.Equal(t, order.UserId, result.UserId)
	assert.Equal(t, order.PaymentId, result.PaymentId)
	assert.Equal(t, order.TotalHour, result.TotalHour)
	assert.Equal(t, order.TotalPayment, result.TotalPayment)
	assert.Equal(t, order.TotalQty, result.TotalQty)
}

// TODO update user test
func TestUserUsecase_UpdateUser(t *testing.T) {
	userId := "e694b986-cf9b-4b33-9147-3838e9014662"

	user := &model.User{
		Fullname:  "Arvin",
		Phone:     "085",
		Address:   "Jl Rinjani",
		UpdatedAt: time.Now(),
	}

	pkg.UserRepository.Mock.On("FindById", userId).Return(user, nil)

	pkg.UserRepository.Mock.On("Update", userId, mock.Anything).Return(nil)

	userDTO := dto.UserDTO{
		Fullname: "Kujo Jotaro",
		Phone:    "08654321",
		Address:  "Jl Morioh",
	}

	err := userUsecaseTest.UpdateUser(userId, userDTO)

	assert.Nil(t, err)
}

func TestUserUsecase_DeleteUser(t *testing.T) {
	userId := "eda51b42-36dc-4d81-8bec-90aee25790d1"

	pkg.UserRepository.Mock.On("Delete", userId).Return(nil)

	err := userUsecaseTest.DeleteUser(userId)

	assert.Nil(t, err)
}
