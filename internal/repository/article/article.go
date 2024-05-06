package article

import (
	"error"
	"errors"

	"github.com/SaharSun/myweblog/internal/entity/article"
)

 var (
	ErrArticleNotFound  = errors.New("Article Not found")
	ErrArticleAlreadyExist = errors.New("Article Already Exist")
 )



type Repository interface {
	Creat(article *article.Article)error
	Update(article *article.Article)error
	Detial()*article.Article
	Delet() error
	List()[]*article.Article
}