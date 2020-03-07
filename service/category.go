package service

import (
	"blogger/dao/db"
	"blogger/model"
)

func GetAllCategories() (categories []*model.Category, err error) {
	allCategories, err := db.GetAllCategories()
	if err != nil {
		return
	}
	categories = allCategories
	return
}
