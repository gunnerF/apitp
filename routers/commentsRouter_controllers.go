package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["apitp/controllers:GroupController"] = append(beego.GlobalControllerRouter["apitp/controllers:GroupController"],
        beego.ControllerComments{
            Method: "GroupAdd",
            Router: `/groupAdd`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitp/controllers:GroupController"] = append(beego.GlobalControllerRouter["apitp/controllers:GroupController"],
        beego.ControllerComments{
            Method: "GroupDelete",
            Router: `/groupDelete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitp/controllers:GroupController"] = append(beego.GlobalControllerRouter["apitp/controllers:GroupController"],
        beego.ControllerComments{
            Method: "WordDown",
            Router: `/word-down`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitp/controllers:LoginController"] = append(beego.GlobalControllerRouter["apitp/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitp/controllers:NodeController"] = append(beego.GlobalControllerRouter["apitp/controllers:NodeController"],
        beego.ControllerComments{
            Method: "GetList",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitp/controllers:NodeController"] = append(beego.GlobalControllerRouter["apitp/controllers:NodeController"],
        beego.ControllerComments{
            Method: "GetNodeType",
            Router: `/getNodeType`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitp/controllers:NodeController"] = append(beego.GlobalControllerRouter["apitp/controllers:NodeController"],
        beego.ControllerComments{
            Method: "NodeAdd",
            Router: `/nodeAdd`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitp/controllers:NodeController"] = append(beego.GlobalControllerRouter["apitp/controllers:NodeController"],
        beego.ControllerComments{
            Method: "NodeDetail",
            Router: `/nodeDetail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitp/controllers:SmartAliPayController"] = append(beego.GlobalControllerRouter["apitp/controllers:SmartAliPayController"],
        beego.ControllerComments{
            Method: "HandleUserPay",
            Router: `/native`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitp/controllers:SmartAliPayController"] = append(beego.GlobalControllerRouter["apitp/controllers:SmartAliPayController"],
        beego.ControllerComments{
            Method: "Notify",
            Router: `/notify`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["apitp/controllers:SmartAliPayController"] = append(beego.GlobalControllerRouter["apitp/controllers:SmartAliPayController"],
        beego.ControllerComments{
            Method: "Return",
            Router: `/return`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
