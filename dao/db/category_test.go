package db

import "testing"

func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(1)
	if err != nil {
		panic(err)
	}
	t.Logf("category: %#v\n", category)
}

func TestGetCategories(t *testing.T) {
	categories, err := GetCategories([]int64{1, 2, 3})
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		t.Logf("id: %#v; category: %#v\n", category.CategoryId, category)
	}
}

func TestGetAllCategories(t *testing.T) {
	categories, err := GetAllCategories()
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		t.Logf("id: %#v; category: %#v\n", category.CategoryId, category)
	}
}
