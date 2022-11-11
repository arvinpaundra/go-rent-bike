package pkg

import (
	"database/sql/driver"
	repomock "github.com/arvinpaundra/go-rent-bike/internal/repository/gormdb/mock"
	"github.com/stretchr/testify/mock"
	"time"
)

// repository
type Anytime struct{}

func (a Anytime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)

	return ok
}

// usecase tests
var (
	UserRepository        = repomock.UserRepositoryMock{Mock: mock.Mock{}}
	ReportRepository      = repomock.ReportRepositoryMock{Mock: mock.Mock{}}
	HistoryRepository     = repomock.HistoryRepositoryMock{Mock: mock.Mock{}}
	OrderRepository       = repomock.OrderRepositoryMock{Mock: mock.Mock{}}
	OrderDetailRepository = repomock.OrderDetailRepositoryMock{Mock: mock.Mock{}}
	PaymentRepository     = repomock.PaymentRepositoryMock{Mock: mock.Mock{}}
	RenterRepository      = repomock.RenterRepositoryMock{Mock: mock.Mock{}}
	CategoryRepository    = repomock.CategoryRepositoryMock{Mock: mock.Mock{}}
	ReviewRepository      = repomock.ReviewRepositoryMock{Mock: mock.Mock{}}
	BikeRepository        = repomock.BikeRepositoryMock{Mock: mock.Mock{}}
)
