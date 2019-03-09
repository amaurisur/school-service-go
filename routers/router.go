// @APIVersion 0.1.0
// @Title School Service
// @Description School Service is responsible for...
package routers

import (
	"github.com/astaxie/beego"
	"github.com/school/controllers"
	"github.com/school/database"
)

func Init(db database.Database) {

	ns := beego.NewNamespace("/school",
		beego.NSInclude(
			&controllers.HealthController{Controller: controllers.Controller{}},
		),
	)
	beego.AddNamespace(ns)
}
