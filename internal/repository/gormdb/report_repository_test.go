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

type suiteReport struct {
	suite.Suite
	mock             sqlmock.Sqlmock
	reportRepository repository.ReportRepository
}

func (s *suiteReport) SetupSuite() {
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

	s.reportRepository = NewReportRepository(dbGorm)
}

func (s *suiteReport) TestCreate() {
	reportUC := model.Report{
		ID:         "ID-1",
		RenterId:   "RID-1",
		UserId:     "UID-1",
		TitleIssue: "Title Issue",
		BodyIssue:  "Body issue.",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `reports` (`renter_id`,`user_id`,`title_issue`,`body_issue`,`created_at`,`updated_at`,`id`) VALUES (?,?,?,?,?,?,?)")).
		WithArgs("RID-1", "UID-1", "Title Issue", "Body issue.", pkg.Anytime{}, pkg.Anytime{}, "ID-1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.reportRepository.Create(reportUC)

	s.Nil(err)
}

func (s *suiteReport) TestFindAll() {
	report := model.Report{
		ID:         "ID-1",
		RenterId:   "RID-1",
		UserId:     "UID-1",
		TitleIssue: "Title Issue",
		BodyIssue:  "Body issue.",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	reportRow := sqlmock.NewRows([]string{"id", "renter_id", "user_id", "title_issue", "body_issue", "created_at", "updated_at"}).
		AddRow(report.ID, report.RenterId, report.UserId, report.TitleIssue, report.BodyIssue, report.CreatedAt, report.UpdatedAt)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `reports` WHERE renter_id = ?")).
		WithArgs("RID-1").
		WillReturnRows(reportRow)

	user := model.User{
		ID:        "UID-1",
		Fullname:  "Arvin Paundra",
		Phone:     "087654321",
		Address:   "Jl Rinjani",
		Role:      "customer",
		Email:     "arvin@mail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userRow := sqlmock.NewRows([]string{"id", "fullname", "phone", "address", "role", "email", "created_at", "updated_at"}).
		AddRow(user.ID, user.Fullname, user.Phone, user.Address, user.Role, user.Email, user.CreatedAt, user.UpdatedAt)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `users`.`id`,`users`.`fullname`,`users`.`phone`,`users`.`address`,`users`.`role`,`users`.`email`,`users`.`created_at`,`users`.`updated_at` FROM `users` WHERE `users`.`id` = ?")).
		WillReturnRows(userRow)

	results, err := s.reportRepository.FindAll("RID-1")

	s.Nil(err)
	s.NotNil(results)

	s.Equal(report.ID, (*results)[0].ID)
	s.Equal(report.UserId, (*results)[0].UserId)
	s.Equal(report.RenterId, (*results)[0].RenterId)
	s.Equal(report.TitleIssue, (*results)[0].TitleIssue)
	s.Equal(report.BodyIssue, (*results)[0].BodyIssue)
}

func TestReportRepository(t *testing.T) {
	suite.Run(t, new(suiteReport))
}
