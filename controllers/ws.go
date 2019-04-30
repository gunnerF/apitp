/**********************************************
** @Des: webSocket控制器
** @Author: jgn
** @Date:   2019/4/29 17:21
***********************************************/
package controllers

import (
	"apitp/utils"
	"fmt"
	"github.com/astaxie/beego/toolbox"
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
	//实例化webSocket对象
	ws, err := upgrade.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		c.jsonMsgResult(err, utils.ParamsError["code"].(int), 1, c.resultJsonArr)
	}
	//将客户端对象放入管道map中
	clients[ws] = true

	//定时不断的广播发送到页面上
	//spec: 秒钟：0-59、分钟：0-59、小时：1-23、日期：1-31、月份：1-12、星期：0-6（0 表示周日）
	tk := toolbox.NewTask("wsTask", "0/10 * * * * *", func() error {
		select {
		case msg := <-broadcast:
			fmt.Println("客户端数量：", len(clients))
			for client := range clients {
				err := client.WriteJSON(msg)
				if err != nil {
					log.Printf("发送消息出错 error: %v", err)
					client.Close()
					delete(clients, client)
				}
			}
		default:
			fmt.Println("no data")
		}
		return nil
	})
	//tk.Run()
	toolbox.AddTask("wsTask", tk)
	//启动定时任务
	toolbox.StartTask()
}
