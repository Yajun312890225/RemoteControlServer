package handler

import (
	. "RemoteControlServer/pkg/websocket"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func Connect(c *gin.Context) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	client := &Client{Id: "", Socket: conn, Send: make(chan interface{}), IP: c.ClientIP()}

	GetManager().Register <- client

	go client.Read()
	go client.Write()
}
