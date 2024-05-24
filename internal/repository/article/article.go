package article

import (
	"errors"
	"github.com/SaharSun/myweblog/internal/entity/article"
)

var (
	ErrArticleNotFound = errors.New("article Not found")
	ErrArticleAlreadyExist = errors.New("article Already Exist")
)

type Repository interface {
	Create(article *article.Article) (int64, error)
	Update(article *article.Article) error
	Detail(article *article.Article) (*article.Article, error)
	Delete(article *article.Article) error
	List() ([]*article.Article, error)
}
