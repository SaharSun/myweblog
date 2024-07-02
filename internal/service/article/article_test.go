package article

import (
	"testing"

	"github.com/SaharSun/myweblog/internal/entity/article"
	"github.com/SaharSun/myweblog/mocks/repository"
	"github.com/stretchr/testify/assert"
)

func TestArticle_Create_Success(t *testing.T){
	articleMock := new(repository.ArticleMock)
	articleEnt := article.NewArticle("article", "article-1",
	"asddffgh xdfgggfggggg", []string{"tag1", "tag2","tag3"})
	articleMock.On("Create", articleEnt).Return(nil)

	articleService := NewArticle(articleMock) 
	err := articleService.Create(articleEnt)
	assert.Nil(t , err)
	articleMock.AssertExpectations(t)
}
