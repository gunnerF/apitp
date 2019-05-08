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

type TaskClient struct {
	//客户端socket map
	WsClients map[*websocket.Conn]bool
	TaskMutex sync.Mutex
}

func initClient() *TaskClient {
	return &TaskClient{
		WsClients: make(map[*websocket.Conn]bool),
		TaskMutex: sync.Mutex{},
	}
}

var (
	//缓冲2000条记录
	Broadcast = make(chan models.Message, 2000)
	ClientA   *TaskClient
	ClientB   *TaskClient
	ClientC   *TaskClient
	ClientD   *TaskClient
)

//创建webSocket定时任务
func NewWebSocketTask() {
	ClientA = initClient()
	ClientB = initClient()
	ClientC = initClient()
	ClientD = initClient()
	//定时不断的广播发送到页面上
	go func() {
		//spec: 秒钟：0-59、分钟：0-59、小时：1-23、日期：1-31、月份：1-12、星期：0-6（0 表示周日）
		tk := toolbox.NewTask("wsTask", "0/5 * * * * *", func() error {
			select {
			//读取管道中消息
			case msg := <-Broadcast:
				wg := sync.WaitGroup{}
				wg.Add(4)
				go sendMessage(ClientA, msg, wg)
				go sendMessage(ClientB, msg, wg)
				go sendMessage(ClientC, msg, wg)
				go sendMessage(ClientD, msg, wg)
				wg.Wait()
			default:
				fmt.Println("no data")
				return nil
			}
			return nil
		})
		//tk.Run()
		toolbox.AddTask("wsTask", tk)
		//启动定时任务
		toolbox.StartTask()
	}()
}

//向客户端推送消息
func sendMessage(taskClient *TaskClient, msg models.Message, wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("客户端数量：", len(taskClient.WsClients))
	taskClient.TaskMutex.Lock()
	for client := range taskClient.WsClients {
		err := client.WriteJSON(msg)
		if err != nil { //出错后移除客户端对象
			log.Printf("发送消息出错 error: %v", err)
			client.Close()
			delete(taskClient.WsClients, client)
		}
	}
	taskClient.TaskMutex.Unlock()
}
