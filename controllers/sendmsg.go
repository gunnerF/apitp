/**********************************************
** @Des: 广播消息
** @Author: jgn
** @Date:   2019/4/29 17:26
***********************************************/
package controllers

import (
	"apitp/models"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan models.Message)
)

func init() {
	go handleMessages()
}

//广播发送至页面
func handleMessages() {
	for {
		msg := <-broadcast
		fmt.Println("客户端数量：", len(clients))
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("发送消息出错 error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
