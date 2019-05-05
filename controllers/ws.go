/**********************************************
** @Des: webSocket控制器
** @Author: jgn
** @Date:   2019/4/29 17:21
***********************************************/
package controllers

import (
	"apitp/commands"
	"fmt"
	"log"
)

type WebSocketController struct {
	BaseController
}

func (c *WebSocketController) URLMapping() {
	c.Mapping("GetWs", c.GetWs)
}

// @router / [get]
func (c *WebSocketController) GetWs() {
	fmt.Println("ws init start")
	//实例化webSocket对象
	ws, err := upgrade.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Fatal("ws init error:", err)
	}
	//将客户端对象放入管道map中
	commands.WsChan <- ws
	fmt.Println("ws init end")
}
