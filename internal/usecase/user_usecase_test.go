package usecase

import (
	"testing"

	"github.com/arvinpaundra/go-rent-bike/internal/model"
	repomock "github.com/arvinpaundra/go-rent-bike/internal/repository/gormdb/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = repomock.UserRepositoryMock{Mock: mock.Mock{}}
var userUsecaseTest = NewUserUsecase(&userRepository)

// TODO register user test
//func TestUserUsecase_Register(t *testing.T) {
//	// expected return
//	expected := model.User{
//		Fullname: "Arvin Paundra",
//		Phone:    "085",
//		Address:  "Jl Rinjani",
//		Role:     "customer",
//		Email:    "arvin@mail.com",
//		Password: "123",
//	}
//
//	userDTO := dto.UserDTO{
//		Fullname: "Arvin Paundra",
//		Phone:    "085",
//		Address:  "Jl Rinjani",
//		Role:     "customer",
//		Email:    "arvin@mail.com",
//		Password: "123",
//	}
//
//	userRepository.Mock.On("Create", expected).Return(nil)
//
//	err := userUsecaseTest.RegisterUser(userDTO)
//
//	assert.Nil(t, err)
//}

// TODO login user test
// func TestUserUsecase_LoginUser(t *testing.T) {
// 	testCases := []struct {
// 		Name         string
// 		Body         map[string]string
// 		Error        bool
// 		ErrorMessage error
// 		Expected     *model.User
// 		Result       *model.User
// 	}{
// 		{
// 			Name: "success login",
// 			Body: map[string]string{
// 				"email":    "arvin@mail.com",
// 				"password": "123",
// 			},
// 			Error:        false,
// 			ErrorMessage: nil,
// 			Expected: &model.User{
// 				ID:       "eda51b42-36dc-4d81-8bec-90aee25790d1",
// 				Fullname: "Arvin Paundra",
// 				Phone:    "0854321",
// 				Address:  "Jl Rinjani",
// 				Role:     "customer",
// 				Email:    "arvin@mail.com",
// 			},
// 			Result: &model.User{
// 				ID:       "eda51b42-36dc-4d81-8bec-90aee25790d1",
// 				Fullname: "Arvin Paundra",
// 				Phone:    "0854321",
// 				Address:  "Jl Rinjani",
// 				Role:     "customer",
// 				Email:    "arvin@mail.com",
// 			},
// 		},
// 		{
// 			Name: "failed login",
// 			Body: map[string]string{
// 				"email":    "jotaro@mail.com",
// 				"password": "123",
// 			},
// 			Error:        true,
// 			ErrorMessage: errors.New("user not found"),
// 			Expected: &model.User{
// 				ID:       "63390a70-bc40-4f95-9f72-4f054c437949",
// 				Fullname: "Kujo Jotaro",
// 				Phone:    "0987654321",
// 				Address:  "Jl Ketapang",
// 				Role:     "renter",
// 				Email:    "jotaro@mail.com",
// 			},
// 			Result: nil,
// 		},
// 	}

// 	for _, v := range testCases {
// 		t.Run(v.Name, func(t *testing.T) {
// 			userRepository.Mock.On("FindByEmailAndPassword", v.Body["email"], v.Body["password"], mock.Anything).Return(v.Expected, v.ErrorMessage)

// 			result, err := userUsecaseTest.LoginUser(v.Body["email"], v.Body["password"])

// 			if v.Error {
// 				assert.NotEqual(t, v.Expected, v.Result)
// 				assert.Nil(t, result)
// 				assert.NotNil(t, err)
// 			} else {
// 				assert.Equal(t, v.Expected, v.Result)
// 				assert.Equal(t, v.Body["email"], v.Result.Email)
// 				assert.Nil(t, err)
// 				assert.NotNil(t, result)
// 			}
// 		})
// 	}
// }

func TestUserUsecase_FindAllUsers(t *testing.T) {
	users := &[]model.User{
		{
			Fullname: "Arvin",
			Phone:    "085",
			Role:     "customer",
			Email:    "arvin@mail.com",
		},
		{
			Fullname: "ardana",
			Phone:    "098",
			Role:     "renter",
			Email:    "ardana@mail.com",
		},
	}

	userRepository.Mock.On("FindAll").Return(users, nil)

	results, err := userUsecaseTest.FindAllUsers()

	assert.Nil(t, err)
	assert.NotNil(t, results)
}

func TestUserUsecase_FindByIdUser(t *testing.T) {
	userId := "eda51b42-36dc-4d81-8bec-90aee25790d1"

	user := &model.User{
		ID:       userId,
		Fullname: "Arvin",
		Phone:    "085",
		Address:  "Jl Rinjani",
		Role:     "customer",
		Email:    "arvin@mail.com",
	}

	userRepository.Mock.On("FindById", userId).Return(user, nil)

	result, err := userUsecaseTest.FindByIdUser(userId)

	assert.Nil(t, err)
	assert.NotNil(t, result)
}

// TODO find all user histories test

// TODO find all orders user test

// TODO find order user by id order test

// TODO update user test

func TestUserUsecase_DeleteUser(t *testing.T) {
	userId := "eda51b42-36dc-4d81-8bec-90aee25790d1"

	userRepository.Mock.On("Delete", userId).Return(uint(1), nil)

	rowAffected, err := userUsecaseTest.DeleteUser(userId)

	assert.Nil(t, err)
	assert.NotNil(t, rowAffected)
	assert.Equal(t, uint(1), rowAffected)
}
