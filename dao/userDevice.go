package dao

import "RemoteControlServer/model"

type UserDeviceDao model.UserDevice

func (ud *UserDeviceDao) Bind() error {
	userDevice := model.UserDevice{
		UserId:   ud.UserId,
		DeviceId: ud.DeviceId,
	}
	if err := model.DB.Table("user_device").Create(&userDevice).Error; err != nil {
		return err
	}
	device := model.Device{}
	model.DB.Table("device").Model(&device).Update(model.Device{IsBind: "1"})
	return nil
}
