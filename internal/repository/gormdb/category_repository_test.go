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

type suiteCategory struct {
	suite.Suite
	mock               sqlmock.Sqlmock
	categoryRepository repository.CategoryRepository
}

func (s *suiteCategory) SetupSuite() {
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

	s.categoryRepository = NewCategoryRepositoryGorm(dbGorm)
}

func (s *suiteCategory) TestCreate() {
	categoryUC := model.Category{
		ID:        "ID-1",
		Name:      "BMX",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `categories` (`id`,`name`,`created_at`,`updated_at`) VALUES (?,?,?,?)")).
		WithArgs(categoryUC.ID, categoryUC.Name, pkg.Anytime{}, pkg.Anytime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.categoryRepository.Create(categoryUC)

	s.Nil(err)
}

func (s *suiteCategory) TestFindAll() {
	category := model.Category{
		ID:        "ID-1",
		Name:      "BMX",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(category.ID, category.Name, category.CreatedAt, category.UpdatedAt)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `categories`")).
		WillReturnRows(rows)

	results, err := s.categoryRepository.FindAll()

	s.Nil(err)
	s.NotNil(results)

	s.Equal(category.ID, (*results)[0].ID)
	s.Equal(category.Name, (*results)[0].Name)
}

func (s *suiteCategory) TestFindById() {
	category := model.Category{
		ID:        "ID-1",
		Name:      "BMX",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	row := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(category.ID, category.Name, category.CreatedAt, category.UpdatedAt)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `categories` WHERE id = ? LIMIT 1")).
		WithArgs("ID-1").
		WillReturnRows(row)

	result, err := s.categoryRepository.FindById("ID-1")

	s.Nil(err)
	s.NotNil(result)

	s.Equal(category.ID, result.ID)
	s.Equal(category.Name, result.Name)
}

func (s *suiteCategory) TestUpdate() {
	categoryUC := model.Category{
		Name:      "BMX",
		UpdatedAt: time.Now(),
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `categories` SET `name`=?,`updated_at`=? WHERE id = ?")).
		WithArgs("BMX", pkg.Anytime{}, "ID-1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.categoryRepository.Update("ID-1", categoryUC)

	s.Nil(err)
}

func (s *suiteCategory) TestDelete() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `categories` WHERE id = ?")).
		WithArgs("ID-1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.categoryRepository.Delete("ID-1")

	s.Nil(err)
}

func TestCategoryRepository(t *testing.T) {
	suite.Run(t, new(suiteCategory))
}
