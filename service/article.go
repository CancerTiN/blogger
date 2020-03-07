package service

import (
	"blogger/dao/db"
	"blogger/model"
)

func GetArticleRecords(pageNum, pageSize int) (articleRecords []*model.ArticleRecord, err error) {
	infos, err := db.GetArticleInfos(pageNum, pageSize)
	if err != nil {
		return
	}
	ids := getCategoryIds(infos)
	categories, err := db.GetCategories(ids)
	for _, info := range infos {
		articleRecord := &model.ArticleRecord{ArticleInfo: *info}
		for _, category := range categories {
			if category.CategoryId == info.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecords = append(articleRecords, articleRecord)
	}
	return
}

func getCategoryIds(articleInfos []*model.ArticleInfo) (categoryIds []int64) {
	for _, articleInfo := range articleInfos {
		exist := false
		for _, categoryId := range categoryIds {
			if categoryId == articleInfo.CategoryId {
				exist = true
				break
			}
		}
		if !exist {
			categoryIds = append(categoryIds, articleInfo.CategoryId)
		}
	}
	return
}

func GetArticleRecordsById(categoryId int64, pageNum, pageSize int) (articleRecords []*model.ArticleRecord, err error) {
	infos, err := db.GetArticleInfosByCategoryId(categoryId, pageNum, pageSize)
	if err != nil {
		return
	}
	ids := getCategoryIds(infos)
	categories, err := db.GetCategories(ids)
	for _, info := range infos {
		articleRecord := &model.ArticleRecord{ArticleInfo: *info}
		for _, category := range categories {
			if category.CategoryId == info.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecords = append(articleRecords, articleRecord)
	}
	return
}
