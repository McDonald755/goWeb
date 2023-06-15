package router

import (
	"github.com/gin-gonic/gin"
	"goWeb/config"
	"goWeb/middleware"
	"goWeb/utils"
	"net/http"
)

func APIRouter() *gin.Engine {
	r := gin.Default()
	//Cors跨域中间件,ReqLog请求日志记录中间件
	r.Use(middleware.Cors(), middleware.ReqLog())

	//https映射
	//r.Use(middleware.TlsHandler())

	//路由根目录
	CollectRoute(r)
	return r
}

func CollectRoute(r *gin.Engine) {
	apiVersion := r.Group(config.API_V1_GROUP)

	common := apiVersion.Group(config.COMMON)
	commonRouter(common)

	goWeb := apiVersion.Group(config.GOWEB)
	demo(goWeb)
}

func commonRouter(r *gin.RouterGroup) {
	r.StaticFS("/upload/images", http.Dir(utils.GetImageFullPath()))

	r.GET("/getTimeStamp", GetTimeStamp)
	r.GET("/uploadImage", UploadImage)
	r.GET("/uploadFiles", UploadFiles)
}

func demo(r *gin.RouterGroup) {
	r.POST("/saveOrUpdateGame", SaveOrUpdateGame)
	r.GET("/deleteGame", DeleteGame)
	r.GET("/getGameDetail", GetGameDetail)
	r.POST("/getGames", GetGames)
}
