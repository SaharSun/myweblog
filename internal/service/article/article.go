package article

import (
	"myweblog/internal/repository/article"

	"github.com/SaharSun/myweblog/internal/repository/article"

	entityArticle "github.com/SaharSun/myweblog/internal/entity/article"
)

type Article struct {
	articleRepository article.Repository 
}

func NewArticle(article article.Repository)*Article {
	return &Article{
		articleRepository: articleRepository,
	}
}


func (a Article)Create(artEnt *entityArticle.Article)error {
	a.articleRepository.Creat(artEnt)
}
