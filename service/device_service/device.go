package device_service

import (
	"github.com/EDDYCJY/go-gin-example/models"
	"strconv"
)

func SelectDevice(id string) (models.Device, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return models.Device{}, err
	}
	return models.SelectDevice(intID)
}
