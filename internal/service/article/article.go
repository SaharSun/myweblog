package article

import (
	"errors"
	"myweblog/internal/repository/article"

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
	// if err := a.articleRepository.Creat(artEnt); err != nil {
	// 	if errors.Is(err, article.ErrArticleAlreadyExist) {
	// 		return err
	// 	}
	// 	log.Error().Err(err).Msg("error in creating article")
	// 	return err
	// }
	// return nil
}
