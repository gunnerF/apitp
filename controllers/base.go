/**********************************************
** @Des: 基础控制器
** @Author: jgn
** @Date:   2019-03-11 14:17:37
** @Last Modified by:   jgn
** @Last Modified time: 2019-03-11 14:17:37
***********************************************/
package controllers

import (
	"apitp/models"
	"apitp/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"strings"
)

type BaseController struct {
	beego.Controller
	controllerName string //控制器名称
	actionName     string
	user           *models.Admin       //用户实例
	userId         int                 //用户id
	userName       string              //用户名称
	loginName      string              //登录名称
	page           int                 //页码
	pageSize       int                 //分页值
	allowUrl       string              //权限控制url
	resultJsonArr  utils.ResultJsonArr //接口统一返回值类型
}

//初始化基础信息
func (c *BaseController) Prepare() {
	controllerName, actionName := c.GetControllerAndAction()
	c.controllerName = strings.ToLower(controllerName[:len(controllerName)-10])
	c.actionName = strings.ToLower(actionName)
	c.page = 1
	c.pageSize = 10
	c.Data["version"] = beego.AppConfig.String("version")
	c.Data["siteName"] = beego.AppConfig.String("site.name")
	c.Data["curRoute"] = c.controllerName + "." + c.actionName
	c.Data["curController"] = c.controllerName
	c.Data["curAction"] = c.actionName
	//初始化json返回值对象
	c.resultJsonArr = make(utils.ResultJsonArr, 0)
	//login validate
	if strings.Compare(c.controllerName, "login") != 0 && strings.Compare(c.actionName, "login") != 0 {
		//login validate
		c.userId = 1
		c.userName = "admin"
		//c.auth()
	}
}

//login validate
func (c *BaseController) auth() {
	token := c.GetString("token")
	if token == "" {
		c.jsonMsgResult(utils.MissToken["message"], utils.MissToken["code"].(int), 0, c.resultJsonArr)
	}
	var user = &models.Admin{}
	cheUser, found := utils.Che.Get(token)
	if found && cheUser != nil {
		user = cheUser.(*models.Admin)
	} else {
		c.jsonMsgResult(utils.OverdueToken["message"], utils.OverdueToken["code"].(int), 0, c.resultJsonArr)
	}
	c.userId = user.Id
	c.loginName = user.LoginName
	c.userName = user.RealName
	c.user = user

	//menu
	c.AdminAuth()
	//权限验证
	isHasAuth := strings.Contains(c.allowUrl, c.controllerName+"/"+c.actionName)
	isHasAuth = true
	noAuth := "uploadFile"
	isNoAuth := strings.Contains(noAuth, c.actionName)
	if isHasAuth == false && isNoAuth == false {
		c.jsonMsgResult(utils.NoPermission["message"], utils.NoPermission["code"].(int), 0, c.resultJsonArr)
	}
}

//permission menu
func (c *BaseController) AdminAuth() {

}

//isPost
func (c *BaseController) isPost() bool {
	return c.Ctx.Request.Method == "POST"
}

//获取客户端ip
func (c *BaseController) getClientIP() string {
	str := c.Ctx.Request.RemoteAddr
	last := strings.LastIndex(str, ":")
	return str[:last]
}

//重定向
func (c *BaseController) redirect(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}

//json类型数据返回
//@params message interface{} 返回消息
//@params code int 返回状态码
//@params count int64 数据数量
//@params data interface{} 数据
func (c *BaseController) jsonMsgResult(message interface{}, code int, count int64, data interface{}) {
	result := make(map[string]interface{})
	result["status"] = code
	result["message"] = message
	result["count"] = count
	result["data"] = data
	c.Data["json"] = result
	c.ServeJSON()
	c.StopRun()
}

//请求参数验证
//@params obj interface{} model类型对象
func (c *BaseController) ParamsValidate(obj interface{}) {
	valid := validation.Validation{}
	b, err := valid.Valid(obj)
	if err != nil {
		c.jsonMsgResult(utils.SystemError["message"], utils.SystemError["code"].(int), 1, err)
	}
	if !b {
		for _, err := range valid.Errors {
			c.jsonMsgResult(err.Field+err.Message, utils.ParamsError["code"].(int), 1, c.resultJsonArr)
		}
	}
}
