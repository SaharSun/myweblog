package article  

import (  
	
	// "github.com/SaharSun/myweblog/internal/repository" 
	"github.com/SaharSun/myweblog/internal/infrastracture/logger"  
	"github.com/stretchr/testify/assert"  
	"github.com/SaharSun/myweblog/internal/entity/article"  
	"github.com/SaharSun/myweblog/mocks/repository" 
	"testing" 
	
	
)  

func TestArticle_Create_Success(t *testing.T){  
	// articleMock := new(repository.ArticleMock) 
	articleMock := &repository.ArticleMock{} 
	articleEnt := article.NewArticle("article", "article-1",  
	"asddffgh xdfgggfggggg", []string{"tag1", "tag2","tag3"})  
	  
	// بررسی صدازنی شده متدهای مورد نیاز  
	articleMock.On("Create", articleEnt).Return(nil)  
	loggerMock := new(logger.Logger)  
	articleService := NewArticle(articleMock, loggerMock)   
	err := articleService.Create(articleEnt)  
	assert.Nil(t, err)  
	articleMock.AssertCalled(t, "Create", articleEnt)  
}