/**********************************************
** @Des: 定时任务
** @Author: jgn
** @Date:   2019/4/30 13:52
***********************************************/
package commands

import (
	"apitp/models"
	"fmt"
	"github.com/astaxie/beego/toolbox"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

var (
	MapMutex  = new(sync.Mutex)
	//客户端socket map
	WsClients = make(map[*websocket.Conn]bool)
	//缓冲2000条记录
	Broadcast = make(chan models.Message, 2000)
)

//创建webSocket定时任务
func NewWebSocketTask() {
	//定时不断的广播发送到页面上
	go func() {
		//spec: 秒钟：0-59、分钟：0-59、小时：1-23、日期：1-31、月份：1-12、星期：0-6（0 表示周日）
		tk := toolbox.NewTask("wsTask", "0/10 * * * * *", func() error {
			select {
			case msg := <-Broadcast:
				fmt.Println("客户端数量：", len(WsClients))
				for client := range WsClients {
					err := client.WriteJSON(msg)
					if err != nil {
						log.Printf("发送消息出错 error: %v", err)
						client.Close()
						delete(WsClients, client)
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
	}()
}
