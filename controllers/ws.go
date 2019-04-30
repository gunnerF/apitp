/**********************************************
** @Des: webSocket控制器
** @Author: jgn
** @Date:   2019/4/29 17:21
***********************************************/
package controllers

import (
	"apitp/utils"
	"apitp/commands"
)

type WebSocketController struct {
	BaseController
}

func (c *WebSocketController) URLMapping() {
	c.Mapping("GetWs", c.GetWs)
}

// @router / [get]
func (c *WebSocketController) GetWs() {
	//实例化webSocket对象
	ws, err := upgrade.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		c.jsonMsgResult(err, utils.ParamsError["code"].(int), 1, c.resultJsonArr)
	}
	//将客户端对象放入管道map中
	commands.WsClients[ws] = true
}
