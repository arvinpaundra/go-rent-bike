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

type suiteBike struct {
	suite.Suite
	mock           sqlmock.Sqlmock
	bikeRepository repository.BikeRepository
}

func (s *suiteBike) SetupSuite() {
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

	s.bikeRepository = NewBikeRepositoryGorm(dbGorm)
}

func (s *suiteBike) TestCreate() {
	bikeUC := model.Bike{
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

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `bikes` (`id`,`renter_id`,`category_id`,`name`,`price_per_hour`,`condition`,`description`,`is_available`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?)")).
		WithArgs("BID-1", "RID-1", "CID-1", "Sample Mountain Bike", float64(15000), "Perfect", "Bike descriptions.", "1", pkg.Anytime{}, pkg.Anytime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.bikeRepository.Create(bikeUC)

	s.Nil(err)
}

func (s *suiteBike) TestFindAll() {
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

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `bikes` WHERE name LIKE ?")).
		WithArgs("%%").
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

	results, err := s.bikeRepository.FindAll("")

	s.Nil(err)
	s.NotNil(results)

	s.Equal(bike.ID, (*results)[0].ID)
	s.Equal(bike.RenterId, (*results)[0].RenterId)
	s.Equal(bike.CategoryId, (*results)[0].CategoryId)
	s.Equal(bike.Name, (*results)[0].Name)
	s.Equal(bike.PricePerHour, (*results)[0].PricePerHour)
	s.Equal(bike.Condition, (*results)[0].Condition)
	s.Equal(bike.Description, (*results)[0].Description)
	s.Equal(bike.IsAvailable, (*results)[0].IsAvailable)
}

func (s *suiteBike) TestFindById() {
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

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `bikes` WHERE id = ? LIMIT 1")).
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

	review := model.Review{
		ID:          "RID-1",
		BikeId:      "BID-1",
		UserId:      "UID-1",
		Rating:      5,
		Description: "Review description section.",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	reviewRow := sqlmock.NewRows([]string{"id", "bike_id", "user_id", "rating", "description"}).
		AddRow(review.ID, review.BikeId, review.UserId, review.Rating, review.Description)

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `reviews` WHERE `reviews`.`bike_id` = ?")).
		WithArgs("BID-1").
		WillReturnRows(reviewRow)

	result, err := s.bikeRepository.FindById("BID-1")

	s.Nil(err)
	s.NotNil(result)

	s.Equal(bike.ID, result.ID)
	s.Equal(bike.RenterId, result.RenterId)
	s.Equal(bike.CategoryId, result.CategoryId)
	s.Equal(bike.Name, result.Name)
	s.Equal(bike.PricePerHour, result.PricePerHour)
	s.Equal(bike.Condition, result.Condition)
	s.Equal(bike.Description, result.Description)
	s.Equal(bike.IsAvailable, result.IsAvailable)
}

func (s *suiteBike) TestFindByIdRenter() {
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

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `bikes` WHERE renter_id = ?")).
		WithArgs("RID-1").
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

	results, err := s.bikeRepository.FindByIdRenter("RID-1")

	s.Nil(err)
	s.NotNil(results)

	s.Equal(bike.ID, (*results)[0].ID)
	s.Equal(bike.RenterId, (*results)[0].RenterId)
	s.Equal(bike.CategoryId, (*results)[0].CategoryId)
	s.Equal(bike.Name, (*results)[0].Name)
	s.Equal(bike.PricePerHour, (*results)[0].PricePerHour)
	s.Equal(bike.Condition, (*results)[0].Condition)
	s.Equal(bike.Description, (*results)[0].Description)
	s.Equal(bike.IsAvailable, (*results)[0].IsAvailable)
}

func (s *suiteBike) TestFindByIdCategory() {
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

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `bikes` WHERE category_id = ?")).
		WithArgs("CID-1").
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

	results, err := s.bikeRepository.FindByIdCategory("CID-1")

	s.Nil(err)
	s.NotNil(results)

	s.Equal(bike.ID, (*results)[0].ID)
	s.Equal(bike.RenterId, (*results)[0].RenterId)
	s.Equal(bike.CategoryId, (*results)[0].CategoryId)
	s.Equal(bike.Name, (*results)[0].Name)
	s.Equal(bike.PricePerHour, (*results)[0].PricePerHour)
	s.Equal(bike.Condition, (*results)[0].Condition)
	s.Equal(bike.Description, (*results)[0].Description)
	s.Equal(bike.IsAvailable, (*results)[0].IsAvailable)
}

func (s *suiteBike) TestUpdate() {
	bikeUC := model.Bike{
		CategoryId:   "CID-1",
		Name:         "Sample Mountain Bike",
		PricePerHour: 15000,
		Condition:    "Perfect",
		Description:  "Bike descriptions.",
		IsAvailable:  "1",
		UpdatedAt:    time.Now(),
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `bikes` SET `category_id`=?,`name`=?,`price_per_hour`=?,`condition`=?,`description`=?,`is_available`=?,`updated_at`=? WHERE id = ?")).
		WithArgs("CID-1", "Sample Mountain Bike", float64(15000), "Perfect", "Bike descriptions.", "1", pkg.Anytime{}, "BID-1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.bikeRepository.Update("BID-1", bikeUC)

	s.Nil(err)
}

func (s *suiteBike) TestDelete() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `bikes` WHERE id = ?")).
		WithArgs("BID-1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	err := s.bikeRepository.Delete("BID-1")

	s.Nil(err)
}

func TestBikeRepository(t *testing.T) {
	suite.Run(t, new(suiteBike))
}
