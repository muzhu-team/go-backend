package device_service

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"strconv"
)

func SelectDevice(id string, limit string) ([]models.Device, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return []models.Device{}, err
	}
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		return []models.Device{}, err
	}

	if intID == 0 {
		return models.SelectDevices(intLimit)
	}
	return models.SelectDevice(intID)
}

func EditDevice(dev models.Device) error {
	//intID, err := strconv.Atoi(id)
	//if err != nil {
	//	return err
	//}
	return models.EditDevice(dev)
}

func DeleteDevice(id string) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return models.DeleteDevice(intID)
}

func SelectSensor(id string) (models.Sensor, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return models.Sensor{}, err
	}
	return models.SelectSensor(intID)
}

func EditSensor(sen models.Sensor) error {
	//intID, err := strconv.Atoi(id)
	//if err != nil {
	//	return err
	//}
	return models.EditSensor(sen)
}

func DeleteSensor(id string) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return models.DeleteSensor(intID)
}
