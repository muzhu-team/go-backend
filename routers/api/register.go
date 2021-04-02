package api

import (
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/service/auth_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterReq struct {
	Username string `json:"username" validate:"required" validate:"max=15,min=6"` //解析json username 必須輸入  再限制一下長度 根據業務要求來  比如你的賬號名不能超過10個字母
	Password string `json:"password" validate:"required" validate:"max=15,min=6"`
}

// @Summary Confirm registration information
// @Produce  application/json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /register [post]
func Register(c *gin.Context) {
	var req RegisterReq
	appG := app.Gin{C: c}
	//沒通過驗證，返回錯誤,驗證成功是200
	if err := c.ShouldBindJSON(&req); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	if _, er := auth_service.Reg(req.Username, req.Password); er != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}
