package services

import (
	"fmt"
	"go_demo/models"
)

func AddArticle(content string, title string) (bool, int64) {
	id, err := models.Add(content, title)
	fmt.Println(err)
	if err != nil {
		return false, 0
	}
	return true, id
}
func DeleteArticleByID(id int64) models.Article {
	article := models.DeleteOne(id)
	return article
}
func GetAllArticle() (article []models.Article) {
	article, err := models.GetAll()
	if err != nil {
		return
	}
	return
}

func GetArticleById(id int64) (article *models.Article) {
	article, err := models.GetOne(id)
	if err != nil {
		return
	}
	return article
}
