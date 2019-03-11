package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/school/controllers:HealthController"] = append(beego.GlobalControllerRouter["github.com/school/controllers:HealthController"],
        beego.ControllerComments{
            Method: "Health",
            Router: `/health`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/school/controllers:SchoolController"] = append(beego.GlobalControllerRouter["github.com/school/controllers:SchoolController"],
        beego.ControllerComments{
            Method: "PostPerson",
            Router: `/person`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/school/controllers:SchoolController"] = append(beego.GlobalControllerRouter["github.com/school/controllers:SchoolController"],
        beego.ControllerComments{
            Method: "PostStudent",
            Router: `/student`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/school/controllers:SchoolController"] = append(beego.GlobalControllerRouter["github.com/school/controllers:SchoolController"],
        beego.ControllerComments{
            Method: "GetAllStudent",
            Router: `/student`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
