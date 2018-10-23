package controllers

import (
	"log"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type MyWebSocketController struct {
	beego.Controller
}

var (
	Clients  = make(map[*websocket.Conn]bool)
	Upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	Broadcast = make(chan string)
)

func init() {
	go handleMessages()
}

//广播发送至页面
func handleMessages() {
	for {
		msg := <-Broadcast
		for client := range Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("client.WriteJSON error: %v", err)
				client.Close()
				delete(Clients, client)
			}
		}
	}
}
