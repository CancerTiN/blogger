package db

import (
	"blogger/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func InsertCategory(category *model.Category) (categoryId int64, err error) {
	query := "INSERT INTO `category` (`category_name`, `category_no`) VALUE (?, ?)"
	result, err := DB.Exec(query, category.CategoryName, category.CategoryNo)
	if err != nil {
		fmt.Printf("DB.Exec(%#v) return: (%#v, %#v)\n", query, result, err)
	}
	i, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("result.LastInsertId() return: (%#v, %#v)\n", i, err)
	}
	categoryId = i
	return
}

func GetCategoryById(id int64) (category *model.Category, err error) {
	category = &model.Category{}
	query := "SELECT `id`, `category_name`, `category_no` FROM `category` WHERE `id` = ?"
	err = DB.Get(category, query, id)
	if err != nil {
		fmt.Printf("DB.Get(%#v, %#v, %#v) return: (%#v)\n", category, query, id, err)
	}
	return
}

func GetCategories(ids []int64) (categories []*model.Category, err error) {
	query := "SELECT `id`, `category_name`, `category_no` FROM `category` WHERE `id` IN (?)"
	s, i, err := sqlx.In(query, ids)
	if err != nil {
		fmt.Printf("sqlx.In(%#v, %#v) return: (%#v, %#v, %#v)\n", query, ids, s, i, err)
	}
	err = DB.Select(&categories, s, i...)
	if err != nil {
		fmt.Printf("DB.Select(%#v, %#v, %#v...) return: (%#v)\n", &categories, s, i, err)
	}
	return
}

func GetAllCategories() (categories []*model.Category, err error) {
	query := "SELECT `id`, `category_name`, `category_no` FROM `category` ORDER BY `category_no` ASC"
	err = DB.Select(&categories, query)
	if err != nil {
		fmt.Printf("DB.Get(%#v, %#v) return: (%#v)\n", &categories, query, err)
	}
	return
}
