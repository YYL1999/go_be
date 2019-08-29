package models

import (
	"github.com/astaxie/beego/orm"
)

type Tag struct {
	Link    string `orm:"pk"`
	Message string
}

func init() {
	orm.RegisterModel(new(Tag))
}
func GetTag() (Tag, error) {
	o := orm.NewOrm()
	tag := new(Tag)
	qs := o.QueryTable("tag")
	_, err := qs.All(tag)
	return *tag, err
}
func SetTag(content string, link string) Tag {
	o := orm.NewOrm()
	tag := Tag{link, content}
	o.Update(&tag)
	return tag
}
