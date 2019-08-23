package main

import (
	"go_demo/logs"
	_ "go_demo/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql",
		"root:211314lpf@@tcp(192.168.10.229:3306)/golang?charset=utf8", 30)
	orm.RunSyncdb("default", false, true)

}
func main() {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	logs.StartLog()
	beego.Run()
}
