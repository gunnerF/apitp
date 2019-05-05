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
	//web socket管道
	WsChan = make(chan *websocket.Conn, 2000)
	//客户端socket map
	WsClients = make(map[*websocket.Conn]bool)
	//缓冲2000条记录
	Broadcast = make(chan models.Message, 2000)
	//数据交换管道
	localWsChan = make(chan *websocket.Conn, 2000)
)

//创建webSocket定时任务
func NewWebSocketTask() {
	//定时不断的广播发送到页面上
	go func() {
		//spec: 秒钟：0-59、分钟：0-59、小时：1-23、日期：1-31、月份：1-12、星期：0-6（0 表示周日）
		tk := toolbox.NewTask("wsTask", "0/2 * * * * *", func() error {
			//从管道中读出客户端写入map中
			func() {
				select {
				case ws := <-WsChan:
					WsClients[ws] = true
				default:
					return
				}
			}()
			select {
			//读取管道中消息
			case msg := <-Broadcast:
				//主线程等待
				var wg sync.WaitGroup
				fmt.Println("客户端数量：", len(WsClients))
				//开启协程将客户端写入交换管道
				for client := range WsClients {
					localWsChan<- client
				}
				//开启多个协程并发向客户端推送消息
				for i := 0; i < 4; i++ {
					wg.Add(1)
					go func(msg models.Message) {
						for {
							select {
							case  client := <-localWsChan:
								err := client.WriteJSON(msg)
								if err != nil {//出错后移除客户端对象
									log.Printf("发送消息出错 error: %v", err)
									client.Close()
									delete(WsClients, client)
								}
							default:
								wg.Done()
								return
							}

						}
					}(msg)
				}
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
