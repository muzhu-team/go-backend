package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/EDDYCJY/go-gin-example/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/EDDYCJY/go-gin-example/middleware/jwt"
	"github.com/EDDYCJY/go-gin-example/pkg/export"
	"github.com/EDDYCJY/go-gin-example/pkg/qrcode"
	"github.com/EDDYCJY/go-gin-example/pkg/upload"
	"github.com/EDDYCJY/go-gin-example/routers/api"
	"github.com/EDDYCJY/go-gin-example/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	//后端跨域
	//r.Use(cors.Cors())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.POST("/api/auth", api.GetAuth)

	//用户注册
	r.POST("/api/register", api.Register)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	//test
	//获取设备信息
	r.GET("/api/device", v1.GetDevice)
	//更新设备信息
	r.POST("/api/device", v1.EditDevice)
	//删除设备
	r.DELETE("/api/device/:id", v1.DeleteDevice)

	//获取传感器信息
	r.GET("/api/sensor/:id", v1.GetSensor)
	//更新传感器信息
	r.POST("/api/sensor", v1.EditSensor)
	//删除传感器
	r.DELETE("/api/sensor/:id", v1.DeleteSensor)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		//导出标签
		r.POST("/tags/export", v1.ExportTag)
		//导入标签
		r.POST("/tags/import", v1.ImportTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		//生成文章海报
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)

		//获取设备信息
		apiv1.GET("/device/:id", v1.GetDevice)
		//更新设备信息
		apiv1.POST("/device", v1.EditDevice)
		//删除设备
		apiv1.DELETE("/device/:id", v1.DeleteDevice)

		//获取传感器信息
		apiv1.GET("/sensor/:id", v1.GetSensor)
		//更新传感器信息
		apiv1.POST("/sensor", v1.EditSensor)
		//删除传感器
		apiv1.DELETE("/sensor/:id", v1.DeleteSensor)
	}

	return r
}
