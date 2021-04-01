package api

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/service/auth_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterReq struct {
	Username string `json:"username" binding:"required` //解析json username 必須輸入  再限制一下長度 根據業務要求來  比如你的賬號名不能超過10個字母
	Password string `json:"password" binding:"required`
}

// @Summary Confirm registration information
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /register [post]
func Register(c *gin.Context) {
	var req RegisterReq

	appG := app.Gin{C: c}

	//綁定驗證其
	err := c.ShouldBindJSON(&req)
	fmt.Println(appG, err, "c")
	//沒通過驗證，返回錯誤,驗證成功是200
	if err != nil {
		//app.MarkErrors( err.Error())
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	_, e := auth_service.Reg(req.Username, req.Password)

	fmt.Println(e, "errors")

	//正確以後開始存表

	//if !ok {
	//	app.MarkErrors(valid.Errors)
	//	appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	//	return
	//}

	//if ; err != nil {
	//
	//
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

}
