package repomock

import (
	"github.com/arvinpaundra/go-rent-bike/internal/model"
	"github.com/stretchr/testify/mock"
)

type ReviewRepositoryMock struct {
	Mock mock.Mock
}

func (r *ReviewRepositoryMock) Create(reviewUC model.Review) error {
	ret := r.Mock.Called(reviewUC)

	return ret.Error(0)
}
