package main

import (
	_ "github.com/linzon/beego_blog/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/linzon/beego_blog/models"
)


func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}


func main() {
	beego.Run()
}

