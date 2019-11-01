package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Link struct {
	ID      int `orm:"column(id);auto" description:"id"`
	Url     string
	Content string
}
type Message struct {
	ID      int `orm:"column(id);auto" description:"id"`
	Content string
}
type Tag struct {
	link    []Link
	message []Message
}

func init() {
	orm.RegisterModel(new(Link), new(Message))
}
func GetTag() (Tag, error) {
	o := orm.NewOrm()
	var link []Link
	var message []Message
	qs := o.QueryTable("link")
	db := o.QueryTable("message")
	_, err := qs.All(&link)
	_, errs := db.All(&message)
	var tag Tag
	if errs != nil {
		return tag, nil
	}
	tag.link = link
	tag.message = message
	return tag, err
}
func SetTag(content string, link string) bool {
	o := orm.NewOrm()
	var links Link
	var message Message
	links.Content = link
	message.Content = content
	//message := &Message{Content: content}
	// o.Insert(links)
	// o.Insert(message)
	id, err := o.Insert(&links)
	ids, errs := o.Insert(&message)
	fmt.Println(id, ids, err, errs)
	return true
}
