package article

import (
	"errors"

	loggerInfra "github.com/SaharSun/myweblog/internal/infrastracture/logger"
	"github.com/SaharSun/myweblog/internal/repository/article"

	entityArticle "github.com/SaharSun/myweblog/internal/entity/article"
)

type Article struct {
	articleRepository article.Repository
	logger            loggerInfra.Logger
}

func NewArticle(articleRepository article.Repository, logger loggerInfra.Logger) *Article {
	return &Article{
		articleRepository: articleRepository,
		logger:            logger,
	}
}

func (a *Article) Create(artEnt *entityArticle.Article) error {
	err := a.articleRepository.Create(artEnt)
	if err != nil {
		if errors.Is(err, article.ErrArticleAlreadyExist) {
			return err
		}
		a.logger.Error(err.Error())
		return err
	}
	return nil
}

// پیاده‌سازی متد Delete
func (a Article) Delete(artEnt *entityArticle.Article) error {
	return nil // شبیه‌سازی موفق
}