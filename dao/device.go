package dao

import (
	"RemoteControlServer/model"
	"strings"
)

type DeviceDao model.Device

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
		IsBind:   "0",
	}
	if err := model.DB.Table("device").Create(&device).Error; err != nil {
		return
	}
}
func (d *DeviceDao) GetDevice(IP string) (devices []DeviceDao, err error) {
	table := model.DB.Table("device").Where("ip = ? AND is_bind=0", IP)
	if err = table.Find(&devices).Error; err != nil {
		return
	}
	return
}
