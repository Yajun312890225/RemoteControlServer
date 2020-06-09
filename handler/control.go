package handler

import (
	"RemoteControlServer/dao"
	"RemoteControlServer/model"
	"log"
	"net/http"

	. "RemoteControlServer/pkg/websocket"

	"github.com/gin-gonic/gin"
)

// Control 远程控制
// @Summary 远程控制
// @Description 获取JSON
// @Tags Control
// @Accept  application/json
// @Product application/json
// @Param data body model.Message true "message"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /control [post]
func Control(c *gin.Context) {
	data := model.Message{}
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error,
		})
		return
	}
	mangager := GetManager()
	if len(mangager.Clients) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "未连接",
		})
		return
	}
	mangager.Send(data, data.DeviceId)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}

// GetStatus 获取状态
// @Summary 获取状态
// @Description 获取JSON
// @Tags Control
// @Accept  application/json
// @Product application/json
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /getstatus [get]
func GetStatus(c *gin.Context) {
	mangager := GetManager()
	if len(mangager.Clients) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "未连接",
		})
		return
	} else {
		//这里去树莓派请求状态
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "已经连接",
		})
	}
}

// GetDevice 获取局域网设备
// @Summary 获取局域网设备
// @Description 获取JSON
// @Tags Control
// @Accept  application/json
// @Product application/json
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /getdevice [get]
func GetDevice(c *gin.Context) {
	deDao := dao.DeviceDao{}
	devices, err := deDao.GetDevice(c.ClientIP())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "获取错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": devices,
	})
}

// BindDevice 绑定设备
// @Summary 绑定设备
// @Description 获取JSON
// @Tags Control
// @Accept  application/json
// @Product application/json
// @Param data body model.UserDevice true "UserDevice"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /binddevice [post]
func BindDevice(c *gin.Context) {
	data := dao.UserDeviceDao{}
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error,
		})
		return
	}
	if err := data.Bind(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}
