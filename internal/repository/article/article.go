package article

import (
	"errors"
	"github.com/SaharSun/myweblog/internal/entity/article"

)

 var (
	ErrArticleNotFound  = errors.New("Article Not found")
	ErrArticleAlreadyExist = errors.New("Article Already Exist")
 )



type Repository interface {
	Create(article *article.Article)error
	Update(article *article.Article)error
	Detial()*article.Article
	Delete(id int) error
	List()[]*article.Article
}