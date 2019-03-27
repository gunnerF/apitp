/**********************************************
** @Des: 基础模型
** @Author: jgn
** @Date:   2019-03-11 14:17:37
** @Last Modified by:   jgn
** @Last Modified time: 2019-03-11 14:17:37
***********************************************/
package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
)

type BaseModel struct {
}

//初始化模型
func Init() {
	//设置validate返回消息
	setDefaultMessage()
	host := beego.AppConfig.String("db.host")
	port := beego.AppConfig.String("db.port")
	user := beego.AppConfig.String("db.user")
	password := beego.AppConfig.String("db.password")
	dbName := beego.AppConfig.String("db.name")
	//prefix := beego.AppConfig.String("db.prefix")
	timezone := beego.AppConfig.String("db.timezone")

	if port == "" {
		port = "3306"
	}
	dbConn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8"

	if timezone != "" {
		dbConn += "&loc=" + url.QueryEscape(timezone)
	}
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	//注册连接数据库(mysql)
	orm.RegisterDataBase("default", "mysql", dbConn)
	//注册model
	orm.RegisterModel(new(Node), new(SubGroup), new(Group))

}

//重写表名称
func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}

//汉化验证返回信息
func setDefaultMessage() {
	msg := map[string]string{
		"Required":     "不能为空",
		"Min":          "不能小于%d",
		"Max":          "不能大于%d",
		"Range":        "必须在 %d 到 %d之间",
		"MinSize":      "最小值是 %d",
		"MaxSize":      "最大值是 %d",
		"Length":       "长度是 %d",
		"Alpha":        "Must be valid alpha characters",
		"Numeric":      "Must be valid numeric characters",
		"AlphaNumeric": "Must be valid alpha or numeric characters",
		"Match":        "Must match %s",
		"NoMatch":      "Must not match %s",
		"AlphaDash":    "Must be valid alpha or numeric or dash(-_) characters",
		"Email":        "Must be a valid email address",
		"IP":           "Must be a valid ip address",
		"Base64":       "Must be valid base64 characters",
		"Mobile":       "Must be valid mobile number",
		"Tel":          "Must be valid telephone number",
		"Phone":        "Must be valid telephone or mobile phone number",
		"ZipCode":      "Must be valid zipcode",
	}
	validation.SetDefaultMessage(msg)
}
