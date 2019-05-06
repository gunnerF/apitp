/**********************************************
** @Des: webSocket控制器
** @Author: jgn
** @Date:   2019/4/29 17:21
***********************************************/
package controllers

import (
	"apitp/commands"
	"log"
	"math/rand"
	"time"
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
		log.Fatal("ws init error:", err)
	}
	//随机将客户端对象放入map中
	rand.Seed(time.Now().Unix())
	num := rand.Intn(4)
	//fmt.Println("num:", num)
	switch num {
	case 0:
		commands.ClientA.TaskMutex.Lock()
		commands.ClientA.WsClients[ws] = true
		commands.ClientA.TaskMutex.Unlock()
	case 1:
		commands.ClientB.TaskMutex.Lock()
		commands.ClientB.WsClients[ws] = true
		commands.ClientB.TaskMutex.Unlock()
	case 2:
		commands.ClientC.TaskMutex.Lock()
		commands.ClientC.WsClients[ws] = true
		commands.ClientC.TaskMutex.Unlock()
	case 3:
		commands.ClientD.TaskMutex.Lock()
		commands.ClientD.WsClients[ws] = true
		commands.ClientD.TaskMutex.Unlock()
	}
}
