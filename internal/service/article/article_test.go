package article

import(
	"github.com/SaharSun/myweblog/blob/master/internal/entity/article"
	"testing"

)

func TestArticle_create_success(t *testing.T){
	articleMock := repository.ArticleMock{}
	articleService := NewArticle(articleMock)
	
}