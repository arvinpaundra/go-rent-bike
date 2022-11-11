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

type suiteRenter struct {
	suite.Suite
	mock             sqlmock.Sqlmock
	renterRepository repository.RenterRepository
}

func (s *suiteRenter) SetupSuite() {
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

	s.renterRepository = NewRenterRepositoryGorm(dbGorm)
}

func (s *suiteRenter) TestCreate() {
	renterUC := model.Renter{
		ID:          "RID-1",
		UserId:      "UID-1",
		RentName:    "Twins' Brother Bike Rental",
		RentAddress: "Jl Morioh",
		Description: "Full with description texts",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `renters` (`id`,`user_id`,`rent_name`,`rent_address`,`description`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?)")).
		WithArgs("RID-1", "UID-1", "Twins' Brother Bike Rental", "Jl Morioh", "Full with description texts", pkg.Anytime{}, pkg.Anytime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.renterRepository.Create(renterUC)

	s.Nil(err)
}

func (s *suiteRenter) TestFindAll() {
	renter := model.Renter{
		ID:          "RID-1",
		UserId:      "UID-1",
		RentName:    "Twins' Brother Bike Rental",
		RentAddress: "Jl Morioh",
		Description: "Full with description texts",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	renterRow := sqlmock.NewRows([]string{"id", "user_id", "rent_name", "rent_address", "description", "created_at", "updated_at"}).
		AddRow(renter.ID, renter.UserId, renter.RentName, renter.RentAddress, renter.Description, renter.CreatedAt, renter.UpdatedAt)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `renters` WHERE rent_name LIKE ?")).
		WithArgs("%%").
		WillReturnRows(renterRow)

	user := model.User{
		ID:       "UID-1",
		Fullname: "Arvin Paundra",
		Phone:    "087654321",
		Address:  "Jl Rinjani",
		Role:     "customer",
		Email:    "arvin@mail.com",
	}

	row := sqlmock.NewRows([]string{"id", "fullname", "phone", "address", "role", "email", "created_at", "updated_at"}).
		AddRow(user.ID, user.Fullname, user.Phone, user.Address, user.Role, user.Email, user.CreatedAt, user.UpdatedAt)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `users`.`id`,`users`.`fullname`,`users`.`phone`,`users`.`address`,`users`.`role`,`users`.`email`,`users`.`created_at`,`users`.`updated_at` FROM `users` WHERE `users`.`id` = ?")).
		WithArgs("UID-1").
		WillReturnRows(row)

	results, err := s.renterRepository.FindAll("")

	s.Nil(err)
	s.NotNil(results)

	s.Equal(renter.ID, (*results)[0].ID)
	s.Equal(renter.UserId, (*results)[0].UserId)
	s.Equal(renter.RentName, (*results)[0].RentName)
	s.Equal(renter.RentAddress, (*results)[0].RentAddress)
	s.Equal(renter.Description, (*results)[0].Description)
}

func (s *suiteRenter) TestFindById() {
	renter := model.Renter{
		ID:          "RID-1",
		UserId:      "UID-1",
		RentName:    "Twins' Brother Bike Rental",
		RentAddress: "Jl Morioh",
		Description: "Full with description texts",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	renterRow := sqlmock.NewRows([]string{"id", "user_id", "rent_name", "rent_address", "description", "created_at", "updated_at"}).
		AddRow(renter.ID, renter.UserId, renter.RentName, renter.RentAddress, renter.Description, renter.CreatedAt, renter.UpdatedAt)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `renters` WHERE id = ? LIMIT 1")).
		WithArgs("RID-1").
		WillReturnRows(renterRow)

	user := model.User{
		ID:       "UID-1",
		Fullname: "Arvin Paundra",
		Phone:    "087654321",
		Address:  "Jl Rinjani",
		Role:     "customer",
		Email:    "arvin@mail.com",
	}

	row := sqlmock.NewRows([]string{"id", "fullname", "phone", "address", "role", "email", "created_at", "updated_at"}).
		AddRow(user.ID, user.Fullname, user.Phone, user.Address, user.Role, user.Email, user.CreatedAt, user.UpdatedAt)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT `users`.`id`,`users`.`fullname`,`users`.`phone`,`users`.`address`,`users`.`role`,`users`.`email`,`users`.`created_at`,`users`.`updated_at` FROM `users` WHERE `users`.`id` = ?")).
		WithArgs("UID-1").
		WillReturnRows(row)

	results, err := s.renterRepository.FindById("RID-1")

	s.Nil(err)
	s.NotNil(results)

	s.Equal(renter.ID, results.ID)
	s.Equal(renter.UserId, results.UserId)
	s.Equal(renter.RentName, results.RentName)
	s.Equal(renter.RentAddress, results.RentAddress)
	s.Equal(renter.Description, results.Description)
}

func (s *suiteRenter) TestFindByIdUser() {
	renter := model.Renter{
		ID:          "RID-1",
		UserId:      "UID-1",
		RentName:    "Twins' Brother Bike Rental",
		RentAddress: "Jl Morioh",
		Description: "Full with description texts",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	renterRow := sqlmock.NewRows([]string{"id", "user_id", "rent_name", "rent_address", "description", "created_at", "updated_at"}).
		AddRow(renter.ID, renter.UserId, renter.RentName, renter.RentAddress, renter.Description, renter.CreatedAt, renter.UpdatedAt)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `renters` WHERE `user_id` = ? LIMIT 1")).
		WithArgs("UID-1").
		WillReturnRows(renterRow)

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

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `bikes` WHERE `bikes`.`renter_id` = ?")).
		WithArgs("RID-1").
		WillReturnRows(bikeRow)

	user := model.User{
		ID:       "UID-1",
		Fullname: "Arvin Paundra",
		Phone:    "087654321",
		Address:  "Jl Rinjani",
		Role:     "customer",
		Email:    "arvin@mail.com",
	}

	row := sqlmock.NewRows([]string{"id", "fullname", "phone", "address", "role", "email", "created_at", "updated_at"}).
		AddRow(user.ID, user.Fullname, user.Phone, user.Address, user.Role, user.Email, user.CreatedAt, user.UpdatedAt)

	s.mock.ExpectQuery(regexp.QuoteMeta("")).
		WithArgs("UID-1").
		WillReturnRows(row)

	results, err := s.renterRepository.FindByIdUser("UID-1")

	s.Nil(err)
	s.NotNil(results)

	s.Equal(renter.ID, results.ID)
	s.Equal(renter.UserId, results.UserId)
	s.Equal(renter.RentName, results.RentName)
	s.Equal(renter.RentAddress, results.RentAddress)
	s.Equal(renter.Description, results.Description)
}

func (s *suiteRenter) TestUpdate() {
	renterUC := model.Renter{
		RentName:    "Twins' Brother Bike Rental",
		RentAddress: "Jl Morioh",
		Description: "Full with description texts",
		UpdatedAt:   time.Now(),
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `renters` SET `rent_name`=?,`rent_address`=?,`description`=?,`updated_at`=? WHERE id = ?")).
		WithArgs("Twins' Brother Bike Rental", "Jl Morioh", "Full with description texts", pkg.Anytime{}, "RID-1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.renterRepository.Update("RID-1", renterUC)

	s.Nil(err)
}

func (s *suiteRenter) TestDelete() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `renters` WHERE id = ?")).
		WithArgs("RID-1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.renterRepository.Delete("RID-1")

	s.Nil(err)
}

func TestRenterRepository(t *testing.T) {
	suite.Run(t, new(suiteRenter))
}
