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

//id           int    `json:"id" valid:"Required;Min(1)"`
//type         int    `json:"type" valid:"Required;Min(1)"`
//name         string `json:"name" valid:"Required;MaxSize(100)"`
//photo        string `json:"photo" valid:"MaxSize(255)"`
//model        string `json:"model" valid:"Required;MaxSize(100)"`
//purchase_date int    `json:"purchase_date" valid:"Required;Min(1)"`
//manufacturer string `json:"manufacturer" valid:"Required;MaxSize(100)"`
//status_num    int    `json:"status_num" valid:"Min(0)"`
//userId       int    `json:"equipment_id" valid:"Required,Min(1)"`

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

func SelectDevice(id int) (Device, error) {

	var dev Device
	deviceInfo := db.Select("*").Where(Device{ID: id}).First(&dev)
	if dev.ID == 0 || (deviceInfo.Error != nil && deviceInfo.Error != gorm.ErrRecordNotFound) {
		return Device{}, deviceInfo.Error
	}
	return dev, nil
}

func EditDevice(dev Device) error {

	if err := db.Save(&dev).Error; err != nil {
		return err
	}
	return nil
}

func DeleteDevice(id int) error {

	var dev Device

	result := db.Delete(&dev, id)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("无此设备")
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
