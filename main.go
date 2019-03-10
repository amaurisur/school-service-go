package main

import (
	"github.com/astaxie/beego"
	"github.com/school/database"
	_ "github.com/school/routers"
)

func init() {
	// Configure ORM DB Driver and Register DB
	database.RegisterDriverAndDatabase()
}

func main() {
	beego.Run()
}
