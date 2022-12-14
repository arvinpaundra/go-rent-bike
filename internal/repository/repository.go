package repository

import (
	"github.com/arvinpaundra/go-rent-bike/internal/model"
)

type UserRepository interface {
	Create(userUC model.User) error
	FindByEmail(email string) (*model.User, error)
	FindAll() (*[]model.User, error)
	FindById(userId string) (*model.User, error)
	Update(userId string, userUC model.User) error
	Delete(userId string) error
}

type CategoryRepository interface {
	Create(categoryUC model.Category) error
	FindAll() (*[]model.Category, error)
	FindById(categoryId string) (*model.Category, error)
	Update(categoryId string, categoryUC model.Category) error
	Delete(categoryId string) error
}

type RenterRepository interface {
	Create(renterUC model.Renter) error
	FindAll(rentName string) (*[]model.Renter, error)
	FindById(renterId string) (*model.Renter, error)
	FindByIdUser(userId string) (*model.Renter, error)
	Update(renterId string, renterUC model.Renter) error
	Delete(renterId string) error
}

type BikeRepository interface {
	Create(bikeUC model.Bike) error
	FindAll(bikeName string) (*[]model.Bike, error)
	FindById(bikeId string) (*model.Bike, error)
	FindByIdRenter(renterId string) (*[]model.Bike, error)
	FindByIdCategory(categoryId string) (*[]model.Bike, error)
	Update(bikeId string, bikeUC model.Bike) error
	Delete(bikeId string) error
}

type PaymentRepository interface {
	Create(paymentUC model.Payment) error
	FindById(paymentId string) (*model.Payment, error)
	Update(paymentId string, paymentUC model.Payment) error
}

type OrderRepository interface {
	Create(orderUC model.Order) error
	FindAll(userId string) (*[]model.Order, error)
	FindById(orderId string) (*model.Order, error)
}

type OrderDetailRepository interface {
	Create(orderDetailUC []model.OrderDetail) error
	FindByIdOrder(orderId string) (*[]model.OrderDetail, error)
}

type ReviewRepository interface {
	Create(reviewUC model.Review) error
}

type HistoryRepository interface {
	Create(historyUC model.History) error
	FindAll(userId string) (*[]model.History, error)
	FindByIdOrder(orderId string) (*model.History, error)
	Update(orderId string, historyUC model.History) error
}

type ReportRepository interface {
	Create(reportUC model.Report) error
	FindAll(renterId string) (*[]model.Report, error)
}
