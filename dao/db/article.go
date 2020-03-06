package db

import (
	"blogger/model"
	_ "github.com/go-sql-driver/mysql"
)

func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	query := "INSERT INTO `article` (`content`, `summary`, `title`, `username`, `category_id`, `view_count`, `comment_count`, `create_time`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := DB.Exec(query, article.Content, article.Summary, article.Title, article.Username,
		article.ArticleInfo.CategoryId, article.ArticleInfo.ViewCount, article.ArticleInfo.CommentCount,
		article.CreateTime.String()[:19])
	if err != nil {
		return
	}
	i, err := result.LastInsertId()
	if err != nil {
		return
	}
	articleId = i
	return
}

func GetArticleInfos(pageNum, pageSize int) (articleInfos []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize < 0 {
		return
	}
	query := "SELECT `id`, `summary`, `title`, `view_count`, `create_time`, `comment_count`, `username`, `category_id` FROM `article` WHERE `status` = 1 ORDER BY `create_time` DESC LIMIT ?, ?"
	err = DB.Select(&articleInfos, query, pageNum-1, pageSize)
	return
}

func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	if articleId < 0 {
		return
	}
	query := "SELECT `id`, `summary`, `title`, `view_count`, `content`, `create_time`, `comment_count`, `username`, `category_id` FROM `article` WHERE `id` = ? AND `status` = 1"
	err = DB.Get(&articleDetail, query, articleId)
	return
}

func GetArticleInfosByCategoryId(categoryId, int64, pageNum, pageSize int) (articleInfos []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize < 0 {
		return
	}
	query := "SELECT `id`, `summary`, `title`, `view_count`, `create_time`, `comment_count`, `username`, `category_id` FROM `article` WHERE `category_id` = ? AND `status` = 1 ORDER BY `create_time` DESC LIMIT ?, ?"
	err = DB.Select(&articleInfos, query, categoryId, pageNum-1, pageSize)
	return
}
