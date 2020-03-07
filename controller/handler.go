package controller

import (
	"blogger/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func IndexHandler(c *gin.Context) {
	records, err := service.GetArticleRecords(0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	categories, err := service.GetAllCategories()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  records,
		"category_list": categories,
	})
}

func CategoryHandler(c *gin.Context) {
	query := c.Query("category_id")
	categoryId, err := strconv.ParseInt(query, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	records, err := service.GetArticleRecordsById(categoryId, 0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	categories, err := service.GetAllCategories()
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  records,
		"category_list": categories,
	})
}
