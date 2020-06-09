package dao

import (
	"RemoteControlServer/model"
)

type UserDao struct {
	model.User
	model.LoginCode
}

func (u *UserDao) UserOnline() {
	var count int
	model.DB.Table("user").Where("user_id = ?", u.UserId).Count(&count)
	if count > 0 {
		return
	}
	user := model.User{
		UserId:   u.UserId,
		NickName: u.NickName,
	}
	if err := model.DB.Table("user").Create(&user).Error; err != nil {
		return
	}
}
