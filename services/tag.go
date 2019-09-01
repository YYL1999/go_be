package services

import (
	"go_demo/models"
)

func GetTag() (models.Tag, error) {
	content, err := models.GetTag()
	return content, err
}
func SetTag(content string, link string) bool {
	ok := models.SetTag(content, link)

	return ok
}
