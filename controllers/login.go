package controllers

import (
	"apitp/services"
	"apitp/utils"
	"github.com/patrickmn/go-cache"
	"strconv"
	"strings"
	"time"
)

type LoginController struct {
	BaseController
}

func (c *LoginController) URLMapping() {
	c.Mapping("Login", c.Login)
}

//@router / [post]
func (c *LoginController) Login() {
	if c.userId > 0 {
		c.jsonMsgResult(utils.AlreadyLogin["message"], utils.AlreadyLogin["message"].(int), 0, c.resultJsonArr)
	}
	username := strings.TrimSpace(c.GetString("username"))
	password := strings.TrimSpace(c.GetString("password"))
	if username != "" && password != "" {
		user, err := new(services.AdminService).AdminGetByName(username)
		//flash := beego.NewFlash()
		if err != nil || user.Password != utils.Md5([]byte(password+user.Salt)) {
			c.jsonMsgResult(utils.UserOrPassword["message"], utils.UserOrPassword["code"].(int), 0, c.resultJsonArr)
		} else if user.Status == 0 {
			c.jsonMsgResult(utils.AccountDisabled["message"], utils.AccountDisabled["code"].(int), 0, c.resultJsonArr)
		} else {
			user.LastIp = c.getClientIP()
			user.LastLogin = time.Now().Unix()
			userService := new(services.AdminService)
			userService.Update(user)
			token := utils.Md5([]byte(c.getClientIP() + "|" + user.Salt + strconv.FormatInt(time.Now().Unix(), 10)))
			utils.Che.Set(token, user, cache.DefaultExpiration)
			data := make(utils.ResultJsonArr, 0)
			data = append(data, map[string]interface{}{"token": token})
			c.jsonMsgResult(utils.LoginSuccess["message"], utils.LoginSuccess["code"].(int), 1, data)
		}
	} else {
		c.jsonMsgResult(utils.UserPwdEmpty["message"], utils.UserPwdEmpty["code"].(int), 1, c.resultJsonArr)
	}
}
