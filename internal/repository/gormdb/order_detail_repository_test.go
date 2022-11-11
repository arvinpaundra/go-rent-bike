package gormdb

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/arvinpaundra/go-rent-bike/internal/repository"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"
)

type suiteOrderDetail struct {
	suite.Suite
	mock                  sqlmock.Sqlmock
	orderDetailRepository repository.OrderDetailRepository
}

func (s *suiteOrderDetail) SetupSuite() {
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

	s.orderDetailRepository = NewOrderDetailRepository(dbGorm)
}

func (s *suiteOrderDetail) Create() {
	orderDetailUC := []model.OrderDetail{
		{
			ID:      "ODID-1",
			OrderId: "OID-1",
			BikeId:  "BID-1",
		},
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO order_details")).
		WithArgs("ODID-1", "OID-1", "BID-1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.orderDetailRepository.Create(orderDetailUC)

	s.Nil(err)
}

func (s *suiteOrderDetail) TestFindByIdOrder() {
	orderDetail := model.OrderDetail{
		ID:      "ODID-1",
		OrderId: "OID-1",
		BikeId:  "BID-1",
	}

	orderDetailRow := sqlmock.NewRows([]string{"id", "order_id", "bike_id"}).
		AddRow(orderDetail.ID, orderDetail.OrderId, orderDetail.BikeId)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `order_details` WHERE order_id = ?")).
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

	results, err := s.orderDetailRepository.FindByIdOrder("OID-1")

	s.Nil(err)
	s.NotNil(results)

	s.Equal(orderDetail.ID, (*results)[0].ID)
	s.Equal(orderDetail.OrderId, (*results)[0].OrderId)
	s.Equal(orderDetail.BikeId, (*results)[0].BikeId)
}

func TestOrderDetailRepository(t *testing.T) {
	suite.Run(t, new(suiteOrderDetail))
}
