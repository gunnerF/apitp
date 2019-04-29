/**********************************************
** @Des: goTest
** @Author: jgn
** @Date:   2019/4/29 17:21
***********************************************/
package controllers

import (
	"apitp/utils"
	"github.com/gorilla/websocket"
	"net/http"
	//"time"
	//"apitp/models"
)

type WebSocketController struct {
	BaseController
}

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	}}

func (c *WebSocketController) URLMapping() {
	c.Mapping("GetWs", c.GetWs)
}

// @router / [get]
func (c *WebSocketController) GetWs() {
	ws, err := upgrade.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		c.jsonMsgResult(err, utils.ParamsError["code"].(int), 1, c.resultJsonArr)
	}
	clients[ws] = true
	//msg := models.Message{Message: "连接成功 " + time.Now().Format("2006-01-02 15:04:05")}
	//broadcast <- msg
	//不断的广播发送到页面上
	//for {
	//	//目前存在问题 定时效果不好 需要在业务代码替换时改为beego toolbox中的定时器
	//	time.Sleep(time.Second * 3)
	//	msg := models.Message{Message: "这是向页面发送的数据 " + time.Now().Format("2006-01-02 15:04:05")}
	//	broadcast <- msg
	//}
}
