package router

import (
	"RemoteControlServer/handler"
	"os"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter 配置路由
func InitRouter() *gin.Engine {
	gin.SetMode("debug")
	r := gin.Default()
	// 配置swagger
	swagURL := ginSwagger.URL(os.Getenv("SWAGGER_URL"))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swagURL))

	r.GET("/ws", handler.Connect)
	r.POST("/control", handler.Control)
	r.GET("/getstatus", handler.GetStatus)
	r.GET("/getdevice", handler.GetDevice)
	r.POST("/login", handler.Login)
	r.POST("/binddevice", handler.BindDevice)
	return r
}
