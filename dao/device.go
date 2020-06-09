package dao

import (
	"RemoteControlServer/model"
	"strings"
)

type DeviceDao model.Device

type BindDeviceDao struct {
	UserId   string `json:"userId"`
	DeviceId string `json:"deviceId"`
}

func (d *DeviceDao) DeviceOnline(deviceId, IP string) {
	var count int
	devicestr := strings.ToUpper(strings.Replace(deviceId, ":", "", -1))
	model.DB.Table("device").Where("device_id = ?", devicestr).Count(&count)
	if count > 0 {
		return
	}
	device := DeviceDao{
		DeviceId: devicestr,
		IP:       IP,
		BindUser: "",
	}
	if err := model.DB.Table("device").Create(&device).Error; err != nil {
		return
	}
}
func (d *DeviceDao) GetDeviceByIP(IP string) (devices []DeviceDao, err error) {
	table := model.DB.Table("device").Where("ip = ? AND bind_user=''", IP)
	if err = table.Find(&devices).Error; err != nil {
		return
	}
	return
}
func (d *DeviceDao) Bind(bindDevice BindDeviceDao, IP string) error {
	//根据ip做判断是否可以绑定设备
	device := model.Device{}
	if err := model.DB.Table("device").Model(&device).Update(model.Device{DeviceId: bindDevice.DeviceId, BindUser: bindDevice.UserId}).Error; err != nil {
		return err
	}
	return nil
}
func (d *DeviceDao) GetDevicesByUser(user string) (devices []model.Device, err error) {
	table := model.DB.Table("device").Where("bind_user= ?", user)
	if err = table.Find(&devices).Error; err != nil {
		return
	}
	return
}

func (d *DeviceDao) GetDevice(deviceId string) (device DeviceDao, err error) {
	table := model.DB.Table("device").Where("device_id = ?", strings.ToUpper(strings.Replace(deviceId, ":", "", -1)))
	if err = table.First(&device).Error; err != nil {
		return
	}
	return
}
