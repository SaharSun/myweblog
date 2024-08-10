package article

import (
	// "errors"
	"github.com/SaharSun/myweblog/internal/entity/article"
	// articleRepository "github.com/SaharSun/myweblog/internal/repository/article"
	
	"github.com/SaharSun/myweblog/mocks/mocks/infrastructure"
	"github.com/SaharSun/myweblog/mocks/mocks/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	// "github.com/SaharSun/myweblog/internal/infrastracture/logger"

	"testing"
)

func TestArticle_Create_Success(t *testing.T) {
	logger := new(infrastructure.LoggerMock)
	articleMock := new(repository.ArticleMock)
	articleEnt := article.NewArticle("article1", "article-1",
		"dsggdsgds lkadsjgdajk dlsagjlgd jalgjdka lkadgs",
		[]string{"tag1", "tag2", "tag3"})
	articleMock.On("Create", articleEnt).Return(1, nil)

	articleService := NewArticle(articleMock, logger)
	id:= articleService.Create(articleEnt)
	// assert.Nil(t, err)
	assert.Equal(t, id, int64(1))
	articleMock.AssertExpectations(t)
}
