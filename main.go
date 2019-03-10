package main

import (
	"github.com/astaxie/beego"
	"github.com/school/controllers"
	"github.com/school/database"
	_ "github.com/school/routers"
)

func init() {
	// Configure ORM DB Driver and Register DB
	database.RegisterDriverAndDatabase("mysql")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/school/swagger"] = "swagger"
	}
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
