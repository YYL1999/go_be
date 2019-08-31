package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Link struct {
	ID      string `orm:"pk"`
	Content string
}
type Message struct {
	ID      string `orm:"pk"`
	Content string
}
type Tag struct {
	link    []Link
	message []Message
}

func init() {
	orm.RegisterModel(new(Link), new(Message))
}
func GetTag() (*Tag, error) {
	o := orm.NewOrm()
	var link []Link
	var message []Message
	qs := o.QueryTable("link")
	db := o.QueryTable("message")
	_, err := qs.All(&link)
	_, errs := db.All(&message)
	if errs != nil {
		return nil, nil
	}
	tag := &Tag{link, message}
	return tag, err
}
func SetTag(content string, link string) bool {
	o := orm.NewOrm()
	links := &Link{Content: link}
	message := &Message{Content: content}
	fmt.Println(links, message)
	// o.Insert(links)
	// o.Insert(message)
	id, err := o.Insert(links)
	ids, errs := o.Insert(message)
	fmt.Println(id, ids, err, errs)
	return true
}
