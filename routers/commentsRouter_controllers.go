package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["apitp/controllers:GroupController"] = append(beego.GlobalControllerRouter["apitp/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "GroupAdd",
			Router:           `/groupAdd`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["apitp/controllers:GroupController"] = append(beego.GlobalControllerRouter["apitp/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "GroupDelete",
			Router:           `/groupDelete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["apitp/controllers:NodeController"] = append(beego.GlobalControllerRouter["apitp/controllers:NodeController"],
		beego.ControllerComments{
			Method:           "GetList",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["apitp/controllers:NodeController"] = append(beego.GlobalControllerRouter["apitp/controllers:NodeController"],
		beego.ControllerComments{
			Method:           "GetNodeType",
			Router:           `/getNodeType`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["apitp/controllers:NodeController"] = append(beego.GlobalControllerRouter["apitp/controllers:NodeController"],
		beego.ControllerComments{
			Method:           "NodeAdd",
			Router:           `/nodeAdd`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["apitp/controllers:NodeController"] = append(beego.GlobalControllerRouter["apitp/controllers:NodeController"],
		beego.ControllerComments{
			Method:           "NodeDetail",
			Router:           `/nodeDetail`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
