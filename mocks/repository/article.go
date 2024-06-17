package repository

import (
	articleEntity "github.com/SaharSun/myweblog/blob/master/internal/entity/article"
	"github.com/stretchr/testify/mock"
)

type ArticleMock struct {
	mock.Mock
}

func (m ArticleMock) Create (article *articleEntity.Article)error{
	args := m.Called(article)
	return args.Error(0)
}