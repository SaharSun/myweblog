package gorm

import (
	"github.com/SaharSun/myweblog/internal/repository/article/gorm"
	"gorm.io/gorm"
)

type Article  struct {  
	db *gorm.DB
}

func NewArticle(db *gorm.DB) *Article {
	return &Article{db: db}
}