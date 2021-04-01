package api

import (
	"github.com/gin-gonic/gin"
)

// @Summary Get Auth
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /register [post]
func Register(c *gin.Context) {

}
