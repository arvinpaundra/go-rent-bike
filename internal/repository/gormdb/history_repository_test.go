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

type suiteHistory struct {
	suite.Suite
	mock              sqlmock.Sqlmock
	historyRepository repository.HistoryRepository
}

func (s *suiteHistory) SetupSuite() {
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

	s.historyRepository = NewHistoryRepository(dbGorm)
}

func (s *suiteHistory) TestCreate() {
	historyUC := model.History{
		ID:         "HID-1",
		OrderId:    "OID-1",
		RentStatus: "pending payment",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `histories` (`id`,`order_id`,`rent_status`,`created_at`,`updated_at`) VALUES (?,?,?,?,?)")).
		WithArgs(historyUC.ID, historyUC.OrderId, historyUC.RentStatus, historyUC.CreatedAt, historyUC.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.historyRepository.Create(historyUC)

	s.Nil(err)
}

func (s *suiteHistory) TestFindAll() {
	history := model.History{
		ID:         "HID-1",
		OrderId:    "OID-1",
		RentStatus: "pending payment",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	historyRow := sqlmock.NewRows([]string{"id", "order_id", "rent_status", "created_at", "updated_at"}).
		AddRow(history.ID, history.OrderId, history.RentStatus, history.CreatedAt, history.UpdatedAt)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `histories`")).
		WillReturnRows(historyRow)

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

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `orders` WHERE `orders`.`id` = ? AND user_id = ?")).
		WithArgs("OID-1", "").
		WillReturnRows(orderRow)

	results, err := s.historyRepository.FindAll("")

	s.Nil(err)
	s.NotNil(results)

	s.Equal(history.ID, (*results)[0].ID)
	s.Equal(history.OrderId, (*results)[0].OrderId)
	s.Equal(history.RentStatus, (*results)[0].RentStatus)
}

func (s *suiteHistory) TestFindById() {
	history := model.History{
		ID:         "HID-1",
		OrderId:    "OID-1",
		RentStatus: "pending payment",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	historyRow := sqlmock.NewRows([]string{"id", "order_id", "rent_status", "created_at", "updated_at"}).
		AddRow(history.ID, history.OrderId, history.RentStatus, history.CreatedAt, history.UpdatedAt)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `histories` WHERE order_id = ? LIMIT 1")).
		WithArgs("OID-1").
		WillReturnRows(historyRow)

	result, err := s.historyRepository.FindByIdOrder("OID-1")

	s.Nil(err)
	s.NotNil(result)

	s.Equal(history.ID, result.ID)
	s.Equal(history.OrderId, result.OrderId)
	s.Equal(history.RentStatus, result.RentStatus)
}

func (s *suiteHistory) TestUpdate() {
	historyUC := model.History{
		RentStatus: "rented",
		UpdatedAt:  time.Now(),
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `histories` SET `rent_status`=?,`updated_at`=? WHERE order_id = ?")).
		WithArgs("rented", pkg.Anytime{}, "OID-1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.historyRepository.Update("OID-1", historyUC)

	s.Nil(err)
}

func TestHistoryRepository(t *testing.T) {
	suite.Run(t, new(suiteHistory))
}
