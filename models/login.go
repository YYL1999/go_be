package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type User struct {
	ID       int `orm:"column(id);pk"`
	Name     string
	Password string
	Tel      int
}

func init() {
	orm.RegisterModel(new(User))
}

//新建用户
func AddUser(user *User) (int64, error) {
	o := orm.NewOrm()         //数据库
	ID, err := o.Insert(user) //插入数据
	return ID, err
}

//查询账号
func GetUserById(ID int64) (*User, error) {
	o := orm.NewOrm()                    //数据库
	user := new(User)                    //User就是第9行struct的数据库，就是mysql的表
	qs := o.QueryTable("user")           //表名为t_user
	err := qs.Filter("id", ID).One(user) //One是指查询一条数据，One(user)是查询mysql表中一条数据
	return user, err
}

//查询用户名
func GetUserByName(name string) (*User, error) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable("user")
	err := qs.One(user)
	return user, err
}

//手机号查询账号
func GetUserByTel(tel string) (*User, error) {
	o := orm.NewOrm()
	user := new(User)                      //User就是第9行struct的数据库，就是mysql的表
	qs := o.QueryTable("user")             //表名为t_user
	err := qs.Filter("tel", tel).One(user) //One是指查询一条数据，One(user)是查询mysql表中一条数据
	return user, err
}

//插入新用户
func InsertUser(tel string, name string, password string, id string) (*User, error) {
	b, error := strconv.Atoi(tel)
	c, error := strconv.Atoi(id)
	if error != nil {
		fmt.Println("字符串转换成整数失败")
	}
	o := orm.NewOrm()
	var user User
	user.Name = name
	user.Password = password
	user.ID = c
	user.Tel = b
	_, err := o.Insert(&user)
	return nil, err
}
