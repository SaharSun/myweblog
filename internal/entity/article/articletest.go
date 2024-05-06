package article

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestArticle_SetTag(t *testing.T) {
	tags := []string{"tag1", "tag2", "tag3"}
	article := NewArticle("title", "slug", "asdfg", []string{})
	article.SetTags(tags)
	expected := "tag1,tag2,tag3"

	assert.Equal(t, expected, article.Tags)

}
func TestNewArticle(t *testing.T) {
	tags := []string{"tag1", "tag2", "tag3"}
	article := NewArticle("title", "slug", "asdfg", tags)
	// article.SetTags(tags)
	expected := "tag1,tag2,tag3"

	assert.Equal(t, expected, article.Tags)
	assert.Equal(t, time.Now().Unix(), article.CreatedAt)

}
