package main

import (
	"RemoteControlServer/conf"
	_ "RemoteControlServer/docs"
	_ "RemoteControlServer/pkg/websocket"
	"RemoteControlServer/router"
	"os"
)

func main() {
	conf.Init()
	r := router.InitRouter()
	r.Run(os.Getenv("PORT"))
}
