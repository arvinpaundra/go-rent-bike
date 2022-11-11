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
)

type suiteUser struct {
	suite.Suite
	mock           sqlmock.Sqlmock
	userRepository repository.UserRepository
}

func (s *suiteUser) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()

	s.NoError(err)

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}))

	s.userRepository = NewUserRepositoryGorm(dbGorm)
}

func (s *suiteUser) TestCreate() {
	user := model.User{
		ID:       "ID-1",
		Fullname: "Arvin Paundra",
		Phone:    "087654321",
		Address:  "Jl Rinjani",
		Role:     "customer",
		Email:    "arvin@mail.com",
		Password: "$2a$10$zWF20TmjfYLz/0ovS2Na9uLLvQbXwN4LAt5TDYCVvFsfgS3TXbR6e",
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`id`,`fullname`,`phone`,`address`,`role`,`email`,`password`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?,?)")).
		WithArgs(user.ID, user.Fullname, user.Phone, user.Address, user.Role, user.Email, user.Password, pkg.Anytime{}, pkg.Anytime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.userRepository.Create(user)

	s.NoError(err)
}

func (s *suiteUser) TestFindByEmail() {
	user := model.User{
		ID:       "ID-1",
		Fullname: "Arvin Paundra",
		Phone:    "087654321",
		Address:  "Jl Rinjani",
		Role:     "customer",
		Email:    "arvin@mail.com",
		Password: "$2a$10$zWF20TmjfYLz/0ovS2Na9uLLvQbXwN4LAt5TDYCVvFsfgS3TXbR6e",
	}

	row := sqlmock.NewRows([]string{"id", "fullname", "phone", "address", "role", "email", "password", "created_at", "updated_at"}).
		AddRow(user.ID, user.Fullname, user.Phone, user.Address, user.Role, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(row)

	result, err := s.userRepository.FindByEmail("arvin@mail.com")

	s.NoError(err)

	s.Equal(user.ID, result.ID)
	s.Equal(user.Fullname, result.Fullname)
	s.Equal(user.Phone, result.Phone)
	s.Equal(user.Role, result.Role)
	s.Equal(user.Email, result.Email)
	s.Equal(user.Password, result.Password)
}

func (s *suiteUser) TestFindAll() {
	user := model.User{
		ID:       "ID-1",
		Fullname: "Arvin Paundra",
		Phone:    "087654321",
		Address:  "Jl Rinjani",
		Role:     "customer",
		Email:    "arvin@mail.com",
		Password: "$2a$10$zWF20TmjfYLz/0ovS2Na9uLLvQbXwN4LAt5TDYCVvFsfgS3TXbR6e",
	}

	row := sqlmock.NewRows([]string{"id", "fullname", "phone", "address", "role", "email", "password", "created_at", "updated_at"}).
		AddRow(user.ID, user.Fullname, user.Phone, user.Address, user.Role, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `users`.`id`,`users`.`fullname`,`users`.`phone`,`users`.`address`,`users`.`role`,`users`.`email`,`users`.`created_at`,`users`.`updated_at` FROM `users`")).
		WillReturnRows(row)

	results, err := s.userRepository.FindAll()

	s.Nil(err)
	s.NotNil(results)

	s.Equal(user.ID, (*results)[0].ID)
	s.Equal(user.Fullname, (*results)[0].Fullname)
	s.Equal(user.Phone, (*results)[0].Phone)
	s.Equal(user.Address, (*results)[0].Address)
	s.Equal(user.Role, (*results)[0].Role)
	s.Equal(user.Email, (*results)[0].Email)
	s.Equal(user.Password, (*results)[0].Password)
}

func (s *suiteUser) TestFindById() {
	user := model.User{
		ID:       "ID-1",
		Fullname: "Arvin Paundra",
		Phone:    "087654321",
		Address:  "Jl Rinjani",
		Role:     "customer",
		Email:    "arvin@mail.com",
		Password: "$2a$10$zWF20TmjfYLz/0ovS2Na9uLLvQbXwN4LAt5TDYCVvFsfgS3TXbR6e",
	}

	row := sqlmock.NewRows([]string{"id", "fullname", "phone", "address", "role", "email", "password", "created_at", "updated_at"}).
		AddRow(user.ID, user.Fullname, user.Phone, user.Address, user.Role, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `users`.`id`,`users`.`fullname`,`users`.`phone`,`users`.`address`,`users`.`role`,`users`.`email`,`users`.`created_at`,`users`.`updated_at` FROM `users`")).
		WithArgs("ID-1").
		WillReturnRows(row)

	result, err := s.userRepository.FindById("ID-1")

	s.Nil(err)
	s.NotNil(result)

	s.Equal(user.ID, result.ID)
	s.Equal(user.Fullname, result.Fullname)
	s.Equal(user.Phone, result.Phone)
	s.Equal(user.Address, result.Address)
	s.Equal(user.Role, result.Role)
	s.Equal(user.Email, result.Email)
	s.Equal(user.Password, result.Password)
}

func (s *suiteUser) TestUpdate() {
	userUC := model.User{
		Fullname: "Arvin Paundra",
		Phone:    "087654321",
		Address:  "Jl Rinjani",
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `fullname`=?,`phone`=?,`address`=?,`updated_at`=? WHERE id = ?")).
		WithArgs(userUC.Fullname, userUC.Phone, userUC.Address, pkg.Anytime{}, "ID-1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.userRepository.Update("ID-1", userUC)

	s.Nil(err)
}

func (s *suiteUser) TestDelete() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `users` WHERE id = ?")).
		WithArgs("ID-1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.userRepository.Delete("ID-1")

	s.Nil(err)
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(suiteUser))
}
