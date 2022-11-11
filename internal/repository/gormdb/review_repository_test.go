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

type suiteReview struct {
	suite.Suite
	mock             sqlmock.Sqlmock
	reviewRepository repository.ReviewRepository
}

func (s *suiteReview) SetupSuite() {
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

	s.reviewRepository = NewReviewRepositoryGorm(dbGorm)
}

func (s *suiteReview) TestCreate() {
	reviewUC := model.Review{
		ID:          "RID-1",
		BikeId:      "BID-1",
		UserId:      "UID-1",
		Rating:      5,
		Description: "Review description section.",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `reviews` (`id`,`bike_id`,`user_id`,`rating`,`description`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?)")).
		WithArgs("RID-1", "BID-1", "UID-1", 5, "Review description section.", pkg.Anytime{}, pkg.Anytime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.reviewRepository.Create(reviewUC)

	s.Nil(err)
}

func TestReviewRepository(t *testing.T) {
	suite.Run(t, new(suiteReview))
}
