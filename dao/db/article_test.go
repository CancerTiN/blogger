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
	id, err := InsertArticle(article)
	if err != nil {
		panic(err)
	}
	t.Logf("id: %#v\n", id)
}

func TestGetArticleInfos(t *testing.T) {
	infos, err := GetArticleInfos(0, 15)
	if err != nil {
		panic(err)
	}
	t.Logf("infos: %#v\n", infos)
	for _, info := range infos {
		t.Logf("info: %#v\n", info)
	}
}

func TestGetArticleDetail(t *testing.T) {
	detail, err := GetArticleDetail(1)
	if err != nil {
		panic(err)
	}
	t.Logf("detail: %#v\n", detail)
}
