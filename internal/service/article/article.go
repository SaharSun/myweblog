package article

import (
	"errors"

	"github.com/SaharSun/myweblog/internal/repository/article"
	loggerInfra "github.com/SaharSun/myweblog/internal/infrastracture/logger"

	entityArticle "github.com/SaharSun/myweblog/internal/entity/article"
)

type Article struct {
	articleRepository article.Repository
	logger            loggerInfra.Logger
}

func NewArticle(articleRepository article.Repository , logger loggerInfra.Logger) *Article {
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

package repository  

type ArticleMock struct {}  

// پیاده‌سازی متد Create  
func (m *ArticleMock) Create(article Article) error {  
    return nil // شبیه‌سازی موفق  
}  

// پیاده‌سازی متد Delete  
func (m *ArticleMock) Delete(id int) error {  
    return nil // شبیه‌سازی موفق  
}