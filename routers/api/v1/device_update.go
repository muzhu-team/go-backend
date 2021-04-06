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
	Type         int    `form:"tag_id" valid:"Required;Min(1)"`
	Name         string `form:"title" valid:"Required;MaxSize(100)"`
	Photo        string `form:"desc" valid:"Required;MaxSize(255)"`
	Model        string `form:"content" valid:"Required;MaxSize(65535)"`
	PurchaseDate int    `form:"modified_by" valid:"Required;MaxSize(100)"`
	Manufacturer string `form:"cover_image_url" valid:"Required;MaxSize(255)"`
	StatusNum    int    `form:"state" valid:"Range(0,1)"`
	UserId       int    `form:"state" valid:"Range(0,1)"`
}

type Sensor struct {
	ID           int    `form:"id" valid:"Required;Min(1)"`
	Type         int    `form:"tag_id" valid:"Required;Min(1)"`
	Name         string `form:"title" valid:"Required;MaxSize(100)"`
	Photo        string `form:"desc" valid:"Required;MaxSize(255)"`
	Model        string `form:"content" valid:"Required;MaxSize(65535)"`
	PurchaseDate int    `form:"modified_by" valid:"Required;MaxSize(100)"`
	Manufacturer string `form:"cover_image_url" valid:"Required;MaxSize(255)"`
	StatusNum    int    `form:"state" valid:"Range(0,1)"`
	EquipmentId  int    `form:"state" valid:"Range(0,1)"`
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
