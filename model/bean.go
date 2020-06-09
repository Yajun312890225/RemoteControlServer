package model

import "time"

// Message
type Message struct {
	DeviceId string `json:"deviceId" binding:"required"`
	Type     int    `json:"type" binding:"required"` //1打开 2 关闭
	During   int    `json:"during"`                  //持续时间

}

// RptMesaage
type RptMesaage struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// DeviceLogin
type DeviceLogin struct {
	MacId string `json:"macid"`
}

type LoginCode struct {
	Code string `json:"code"`
}

type CheckLogin struct {
	SessionKey string `json:"session_key"`
	OpenId     string `json:"openid"`
}

// User 用户
type User struct {
	UserId    string     `gorm:"primary_key;"  json:"userId"`
	NickName  string     `gorm:"type:varchar(128)" json:"nickName"`
	Phone     string     `gorm:"type:varchar(11)" json:"phone" `
	Avatar    string     `gorm:"type:varchar(255)" json:"avatar"`
	Sex       string     `gorm:"type:varchar(255)" json:"sex"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// Device
type Device struct {
	DeviceId  string     `gorm:"primary_key;"  json:"deviceId"`
	IP        string     `gorm:"type:varchar(128)"  json:"ip"`
	BindUser  string     `gorm:"type:varchar(128);" json:"bindUser"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// Res 请求返回统一格式
type Res struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}
