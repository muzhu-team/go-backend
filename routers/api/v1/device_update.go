package v1

import (
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

type Device struct {
	ID           int    `form:"id" valid:"Required;Min(1)"`
	Type         int    `form:"type" valid:"Required;Min(1)"`
	Name         string `form:"name" valid:"Required;MaxSize(100)"`
	Photo        string `form:"photo" valid:"MaxSize(255)"`
	Model        string `form:"model" valid:"Required;MaxSize(100)"`
	PurchaseDate int    `form:"purchase_date" valid:"Required;Min(1)"`
	Manufacturer string `form:"manufacturer" valid:"Required;MaxSize(100)"`
	StatusNum    int    `form:"status_num" valid:"Min(0)"`
	UserId       int    `form:"equipment_id" valid:"Required,Min(1)"`
}

type Sensor struct {
	ID           int    `form:"id" valid:"Required;Min(1)"`
	Type         int    `form:"type" valid:"Required;Min(1)"`
	Name         string `form:"name" valid:"Required;MaxSize(100)"`
	Photo        string `form:"photo" valid:"MaxSize(255)"`
	Model        string `form:"model" valid:"Required;MaxSize(100)"`
	PurchaseDate int    `form:"purchase_date" valid:"Required;Min(1)"`
	Manufacturer string `form:"manufacturer" valid:"Required;MaxSize(100)"`
	StatusNum    int    `form:"status_num" valid:"Min(0)"`
	EquipmentId  int    `form:"equipment_id" valid:"Required,Min(1)"`
}

func EditDevice(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, c.Param("id"))
	//id := com.StrTo(c.Param("id")).MustInt()
	//valid := validation.Validation{}
	//valid.Min(id, 1, "id")
	//
	//if valid.HasErrors() {
	//	app.MarkErrors(valid.Errors)
	//	appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	//	return
	//}
	//
	//articleService := article_service.Article{ID: id}
	//exists, err := articleService.ExistByID()
	//if err != nil {
	//	appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
	//	return
	//}
	//if !exists {
	//	appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
	//	return
	//}
	//
	//article, err := articleService.Get()
	//if err != nil {
	//	appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLE_FAIL, nil)
	//	return
	//}
	//
	//appG.Response(http.StatusOK, e.SUCCESS, article)
}

func DeleteDevice(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, c.Params)
}
