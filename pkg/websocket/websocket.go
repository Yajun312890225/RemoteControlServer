package websocket

import (
	. "RemoteControlServer/model"
	"log"

	. "RemoteControlServer/dao"

	"github.com/gorilla/websocket"
)

type ClientManager struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
}

type Client struct {
	Id     string
	Socket *websocket.Conn
	Send   chan interface{}
	IP     string
}

func (m *ClientManager) Init() {
	for {
		select {
		case conn := <-m.Register:
			m.Clients[conn] = true
			log.Println("A new socket has connected.")
		case conn := <-m.Unregister:
			if _, ok := m.Clients[conn]; ok {
				close(conn.Send)
				delete(m.Clients, conn)
				log.Println("A socket has disconnected.")
			}
		}
	}
}

func (m *ClientManager) Send(message interface{}, deviceId string) {
	for conn := range m.Clients {
		if conn.Id == deviceId {
			conn.Send <- message
		}
	}
}

func (c *Client) Read() {
	defer func() {
		manager.Unregister <- c
		c.Socket.Close()
	}()

	for {
		msg := DeviceLogin{}
		err := c.Socket.ReadJSON(&msg)
		if err != nil {
			manager.Unregister <- c
			c.Socket.Close()
			break
		}
		//此处处理客户端来的数据
		c.Id = msg.MacId
		device := DeviceDao{}
		device.DeviceOnline(c.Id, c.IP)
	}
}

func (c *Client) Write() {
	defer func() {
		c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.Socket.WriteJSON(message)
		}
	}
}

var manager *ClientManager

func GetManager() *ClientManager {
	return manager
}

func init() {
	manager = &ClientManager{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
	go manager.Init()
}
