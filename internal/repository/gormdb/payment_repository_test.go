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

type suitePayment struct {
	suite.Suite
	mock              sqlmock.Sqlmock
	paymentRepository repository.PaymentRepository
}

func (s *suitePayment) SetupSuite() {
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

	s.paymentRepository = NewPaymentRepository(dbGorm)
}

func (s *suitePayment) TestCreate() {
	paymentUC := model.Payment{
		ID:            "PID-1",
		PaymentStatus: "pending",
		PaymentType:   "bank_transfer",
		PaymentLink:   "https://app.sandbox.midtrans.com/snap/redirect/v3/...",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `payments` (`payment_status`,`payment_type`,`payment_link`,`created_at`,`updated_at`,`id`) VALUES (?,?,?,?,?,?)")).
		WithArgs("pending", "bank_transfer", "https://app.sandbox.midtrans.com/snap/redirect/v3/...", pkg.Anytime{}, pkg.Anytime{}, "PID-1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.paymentRepository.Create(paymentUC)

	s.Nil(err)
}

func (s *suitePayment) TestFindById() {
	payment := model.Payment{
		ID:            "PID-1",
		PaymentStatus: "pending",
		PaymentType:   "bank_transfer",
		PaymentLink:   "https://app.sandbox.midtrans.com/snap/redirect/v3/...",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	row := sqlmock.NewRows([]string{"id", "payment_status", "payment_type", "payment_link", "created_at", "updated_at"}).
		AddRow(payment.ID, payment.PaymentStatus, payment.PaymentType, payment.PaymentLink, payment.CreatedAt, payment.UpdatedAt)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `payments` WHERE id = ? LIMIT 1")).
		WithArgs("PID-1").
		WillReturnRows(row)

	result, err := s.paymentRepository.FindById("PID-1")

	s.Nil(err)
	s.NotNil(result)

	s.Equal(payment.ID, result.ID)
	s.Equal(payment.PaymentStatus, result.PaymentStatus)
	s.Equal(payment.PaymentLink, result.PaymentLink)
	s.Equal(payment.PaymentType, result.PaymentType)
}

func (s *suitePayment) TestUpdate() {
	paymentUC := model.Payment{
		PaymentStatus: "settlement",
		PaymentType:   "bank_transfer",
		UpdatedAt:     time.Now(),
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `payments` SET `payment_status`=?,`payment_type`=?,`updated_at`=? WHERE id = ?")).
		WithArgs("settlement", "bank_transfer", pkg.Anytime{}, "PID-1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.paymentRepository.Update("PID-1", paymentUC)

	s.Nil(err)
}

func TestPaymentRepository(t *testing.T) {
	suite.Run(t, new(suitePayment))
}
