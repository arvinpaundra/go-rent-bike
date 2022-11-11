package gormdb

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"github.com/arvinpaundra/go-rent-bike/pkg"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"
)

type suiteOrder struct {
	suite.Suite
	mock            sqlmock.Sqlmock
	orderRepository repository.OrderRepository
}

func (s *suiteOrder) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()

	s.NoError(err)

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}))

	s.orderRepository = NewOrderRepository(dbGorm)
}

func (s *suiteOrder) TestCreate() {
	orderUC := model.Order{
		ID:           "OID-1",
		UserId:       "UID-1",
		PaymentId:    "PID-1",
		TotalPayment: 200000,
		TotalQty:     3,
		TotalHour:    5,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `orders` (`id`,`user_id`,`payment_id`,`total_payment`,`total_qty`,`total_hour`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?)")).
		WithArgs("OID-1", "UID-1", "PID-1", float32(200000), 3, 5, pkg.Anytime{}, pkg.Anytime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.orderRepository.Create(orderUC)

	s.Nil(err)
}

func (s *suiteOrder) TestFindAll() {
	order := model.Order{
		ID:           "OID-1",
		UserId:       "UID-1",
		PaymentId:    "PID-1",
		TotalPayment: 200000,
		TotalQty:     3,
		TotalHour:    5,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	row := sqlmock.NewRows([]string{"id", "user_id", "payment_id", "total_payment", "total_qty", "total_hour", "created_at", "updated_at"}).
		AddRow(order.ID, order.UserId, order.PaymentId, order.TotalPayment, order.TotalQty, order.TotalHour, order.CreatedAt, order.UpdatedAt)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `orders` WHERE user_id = ?")).
		WithArgs("UID-1").
		WillReturnRows(row)

	results, err := s.orderRepository.FindAll("UID-1")

	s.Nil(err)
	s.NotNil(results)

	s.Equal(order.ID, (*results)[0].ID)
	s.Equal(order.UserId, (*results)[0].UserId)
	s.Equal(order.PaymentId, (*results)[0].PaymentId)
	s.Equal(order.TotalPayment, (*results)[0].TotalPayment)
	s.Equal(order.TotalHour, (*results)[0].TotalHour)
	s.Equal(order.TotalQty, (*results)[0].TotalQty)
}

func (s *suiteOrder) TestFindById() {
	order := model.Order{
		ID:           "OID-1",
		UserId:       "UID-1",
		PaymentId:    "PID-1",
		TotalPayment: 200000,
		TotalQty:     3,
		TotalHour:    5,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	orderRow := sqlmock.NewRows([]string{"id", "user_id", "payment_id", "total_payment", "total_qty", "total_hour", "created_at", "updated_at"}).
		AddRow(order.ID, order.UserId, order.PaymentId, order.TotalPayment, order.TotalQty, order.TotalHour, order.CreatedAt, order.UpdatedAt)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `orders` WHERE id = ? LIMIT 1")).
		WithArgs("OID-1").
		WillReturnRows(orderRow)

	orderDetail := model.OrderDetail{
		ID:      "ODID-1",
		OrderId: "OID-1",
		BikeId:  "BID-1",
	}

	orderDetailRow := sqlmock.NewRows([]string{"id", "order_id", "bike_id"}).
		AddRow(orderDetail.ID, orderDetail.OrderId, orderDetail.BikeId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `order_details` WHERE `order_details`.`order_id` = ?")).
		WithArgs("OID-1").
		WillReturnRows(orderDetailRow)

	bike := model.Bike{
		ID:           "BID-1",
		RenterId:     "RID-1",
		CategoryId:   "CID-1",
		Name:         "Sample Mountain Bike",
		PricePerHour: 15000,
		Condition:    "Perfect",
		Description:  "Bike descriptions.",
		IsAvailable:  "1",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	bikeRow := sqlmock.NewRows([]string{"id", "renter_id", "category_id", "name", "price_per_hour", "condition", "description", "is_available"}).
		AddRow(bike.ID, bike.RenterId, bike.CategoryId, bike.Name, bike.PricePerHour, bike.Condition, bike.Description, bike.IsAvailable)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `bikes` WHERE `bikes`.`id` = ?")).
		WithArgs("BID-1").
		WillReturnRows(bikeRow)

	category := model.Category{
		ID:        "CID-1",
		Name:      "BMX",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	categoryRow := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(category.ID, category.Name, category.CreatedAt, category.UpdatedAt)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `categories` WHERE `categories`.`id` = ?")).
		WithArgs("CID-1").
		WillReturnRows(categoryRow)

	payment := model.Payment{
		ID:            "PID-1",
		PaymentStatus: "pending",
		PaymentType:   "bank_transfer",
		PaymentLink:   "https://app.sandbox.midtrans.com/snap/redirect/v3/...",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	paymentRow := sqlmock.NewRows([]string{"id", "payment_status", "payment_type", "payment_link", "created_at", "updated_at"}).
		AddRow(payment.ID, payment.PaymentStatus, payment.PaymentType, payment.PaymentLink, payment.CreatedAt, payment.UpdatedAt)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `payments` WHERE `payments`.`id` = ?")).
		WithArgs("PID-1").
		WillReturnRows(paymentRow)

	result, err := s.orderRepository.FindById("OID-1")

	s.Nil(err)
	s.NotNil(result)

	s.Equal(order.ID, result.ID)
	s.Equal(order.UserId, result.UserId)
	s.Equal(order.PaymentId, result.PaymentId)
	s.Equal(order.TotalPayment, result.TotalPayment)
	s.Equal(order.TotalHour, result.TotalHour)
	s.Equal(order.TotalQty, result.TotalQty)
}

func TestOrderRepository(t *testing.T) {
	suite.Run(t, new(suiteOrder))
}
