package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

// CheckAuth checks if authentication information exists

type Device struct {
	ID           int      `json:"id" valid:"Required;Min(1)"`
	Type         int      `json:"type" valid:"Required;Min(1)"`
	Name         string   `json:"name" valid:"Required;MaxSize(100)"`
	Photo        string   `json:"photo" valid:"MaxSize(255)"`
	Model        string   `json:"model" valid:"Required;MaxSize(100)"`
	PurchaseDate int      `json:"purchase_date" valid:"Required;Min(1)"`
	Manufacturer string   `json:"manufacturer" valid:"Required;MaxSize(100)"`
	StatusNum    int      `json:"status_num" valid:"Min(0)"`
	UserId       int      `json:"user_id" valid:"Required,Min(1)"`
	Sensor       []Sensor `json:"sensors"`
}

type Sensor struct {
	ID           int    `json:"id" valid:"Required;Min(1)"`
	Type         int    `json:"type" valid:"Required;Min(1)"`
	Name         string `json:"name" valid:"Required;MaxSize(100)"`
	Photo        string `json:"photo" valid:"MaxSize(255)"`
	Model        string `json:"model" valid:"Required;MaxSize(100)"`
	PurchaseDate int    `json:"purchase_date" valid:"Required;Min(1)"`
	Manufacturer string `json:"manufacturer" valid:"Required;MaxSize(100)"`
	StatusNum    int    `json:"status_num" valid:"Min(0)"`
	EquipmentId  int    `json:"equipment_id" valid:"Required,Min(1)"`
}

func SelectDevice(id int) ([]Device, error) {

	var dev []Device

	deviceInfo := db.Select("*").Where(Device{ID: id}).First(&dev)

	if len(dev) == 0 {
		return []Device{}, deviceInfo.Error
	}

	sensorInfo := db.Where(Sensor{EquipmentId: dev[0].ID}).Find(&dev[0].Sensor)

	if (deviceInfo.Error != nil && deviceInfo.Error != gorm.ErrRecordNotFound) || (sensorInfo.Error != nil && sensorInfo.Error != gorm.ErrRecordNotFound) {
		return []Device{}, deviceInfo.Error
	}

	return dev, nil
}

func SelectDevices(limit int) ([]Device, error) {

	var dev []Device

	devicesInfo := db.Limit(limit).Find(&dev)

	//sensorInfo := db.Find(&dev[0].Sensor)

	for i := 0; i < len(dev); i++ {
		sensorInfo := db.Where(Sensor{EquipmentId: dev[i].ID}).Find(&dev[i].Sensor)
		if sensorInfo.Error != nil && sensorInfo.Error != gorm.ErrRecordNotFound {
			return []Device{}, sensorInfo.Error
		}
	}

	//if dev[0].ID == 0 || (deviceInfo.Error != nil && deviceInfo.Error != gorm.ErrRecordNotFound) || (sensorInfo.Error != nil && sensorInfo.Error != gorm.ErrRecordNotFound) {
	if dev[0].ID == 0 || (devicesInfo.Error != nil && devicesInfo.Error != gorm.ErrRecordNotFound) {
		return []Device{}, devicesInfo.Error
	}

	return dev, nil
}

func EditDevice(dev Device) error {

	//if dev.ID != dev.Sensor[0].ID{
	//	errors.New("更新设备失败")
	//}

	if err := db.Save(&dev).Error; err != nil {
		return err
	}

	if dev.Sensor[0].ID == 0 {
		return nil
	}

	for i := range dev.Sensor {
		if dev.ID != dev.Sensor[i].EquipmentId {
			return errors.New("参数错误")
		}
		if err := EditSensor(dev.Sensor[i]); err != nil {
			return err
		}
	}
	return nil
}

func DeleteDevice(id int) error {

	var dev Device
	var sen Sensor

	senResult := db.Where("equipment_id = ?", id).Delete(&sen)
	if senResult.Error != nil && senResult.Error != gorm.ErrRecordNotFound {
		return senResult.Error
	}
	devResult := db.Delete(&dev, id)
	if devResult.Error != nil && devResult.Error != gorm.ErrRecordNotFound {
		return devResult.Error
	}
	if devResult.RowsAffected == 0 {
		return errors.New("无此设备")
	}
	return nil
}

func SelectSensor(id int) (Sensor, error) {

	var sen Sensor
	sensorInfo := db.Select("*").Where(Sensor{ID: id}).First(&sen)
	if sen.ID == 0 || (sensorInfo.Error != nil && sensorInfo.Error != gorm.ErrRecordNotFound) {
		return Sensor{}, sensorInfo.Error
	}
	return sen, nil
}

func EditSensor(sen Sensor) error {

	if err := db.Save(&sen).Error; err != nil {
		return err
	}
	return nil
}

func DeleteSensor(id int) error {

	var sen Sensor

	result := db.Delete(&sen, id)
	//fmt.Println(result)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return result.Error
	}
	return nil
}
