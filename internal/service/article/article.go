package article

import (
	"errors"

	"github.com/SaharSun/myweblog/internal/repository/article"
	// "github.com/rs/zerolog/log"

	entityArticle "github.com/SaharSun/myweblog/internal/entity/article"
)

type Article struct {
	articleRepository article.Repository
	// logger            loggerInfra.Logger
}

func NewArticle(articleRepository article.Repository) *Article {
	return &Article{
		articleRepository: articleRepository,
		// logger:            logger,
	}
}

func (a Article) Create(artEnt *entityArticle.Article) error {
	id, err := a.articleRepository.Create(artEnt)
	if err != nil {
		if errors.Is(err, article.ErrArticleAlreadyExist) {
			return id, err
		}
		a.logger.Error(err.Error())
		return id, err
	}
	return id, nil
}
