package v1

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/service/device_service"
	"net/http"

	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/gin-gonic/gin"
)

// @Summary Device Update
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/articles/{id} [get]

func GetDevice(c *gin.Context) {
	appG := app.Gin{C: c}
	deviceInfo, err := device_service.SelectDevice(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_DEVICE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, deviceInfo)
}

func EditSensor(c *gin.Context) {
	appG := app.Gin{C: c}
	var dev models.Device
	//fmt.Println(c.Param("id"))
	//err := device_service.EditDevice(models.Device{})
	//if err != nil {
	//	appG.Response(http.StatusInternalServerError, e.ERROR_GET_DEVICE_FAIL, nil)
	//	return
	//}

	if err := c.ShouldBindJSON(&dev); err != nil {
		fmt.Print(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	fmt.Println("Sensor = ", dev.Sensor[0])

	fmt.Println("id = ", dev.ID)

	err := device_service.EditDevice(dev)
	if err != nil {
		println(err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_DEVICE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, c.Param("id"))
}

func EditDevice(c *gin.Context) {
	appG := app.Gin{C: c}
	var dev models.Device
	//fmt.Println(c.Param("id"))
	//err := device_service.EditDevice(models.Device{})
	//if err != nil {
	//	appG.Response(http.StatusInternalServerError, e.ERROR_GET_DEVICE_FAIL, nil)
	//	return
	//}

	if err := c.ShouldBindJSON(&dev); err != nil {
		fmt.Print(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	fmt.Println("Sensor = ", dev.Sensor[0])

	fmt.Println("id = ", dev.ID)

	err := device_service.EditDevice(dev)
	if err != nil {
		println(err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_DEVICE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, c.Param("id"))
}

func DeleteDevice(c *gin.Context) {
	appG := app.Gin{C: c}
	err := device_service.DeleteDevice(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_DEVICE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, "删除成功")
}
