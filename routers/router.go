// @APIVersion 0.1.0
// @Title School Service
// @Description School Service is responsible for...
package routers

import (
	"github.com/astaxie/beego"
	"github.com/school/controllers"
)

func init() {

	ns := beego.NewNamespace("/school",
		beego.NSInclude(
			&controllers.HealthController{Controller: controllers.Controller{}},
		),
	)
	beego.AddNamespace(ns)
}
