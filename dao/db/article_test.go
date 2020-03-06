package db

import (
	"blogger/model"
	"testing"
	"time"
)

func TestInsertArticle(t *testing.T) {
	article := &model.ArticleDetail{
		ArticleInfo: model.ArticleInfo{CategoryId: 1, CommentCount: 0, CreateTime: time.Now(), Title: "GoTest",
			Username: "tin", Summary: "sum_test", ViewCount: 1},
		Content: "GoTestContent",
	}
	t.Logf("article: %#v\n", article)
	articleId, err := InsertArticle(article)
	if err != nil {
		panic(err)
	}
	t.Logf("articleId: %#v\n", articleId)
}

func TestGetArticleInfos(t *testing.T) {
	articleInfos, err := GetArticleInfos(1, 15)
	if err != nil {
		panic(err)
	}
	t.Logf("articleInfos: %#v\n", articleInfos)
	for _, articleInfo := range articleInfos {
		t.Logf("articleInfo: %#v\n", articleInfo)
	}
}
