package repository

import (
	articleEntity "github.com/SaharSun/myweblog/internal/entity/article"
	"github.com/stretchr/testify/mock"
)

type ArticleMock struct {
	mock.Mock
}

func (m ArticleMock) Create(article *articleEntity.Article) error {
	args := m.Called(article)
	return args.Error(0)
} 

func (m ArticleMock) Update(article *articleEntity.Article) error {
	args := m.Called(article)
	return args.Error(0)
}

func (m ArticleMock) Delete() *articleEntity.Article {
	args := m.Called()
	return args.Get(0).(*articleEntity.Article)
}

func (m ArticleMock) list() *articleEntity.Article {
	args := m.Called()
	return args.Get(0).(*articleEntity.Article) 
}
