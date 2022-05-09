package router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goWeb/config"
	"net/http"

	"github.com/unrolled/secure"
)

func APIRouter() *gin.Engine {
	r := gin.Default()
	//跨域中间件
	r.Use(Cors())

	//https映射
	//r.Use(TlsHandler())

	//请求日志记录中间件
	r.Use(ReqLog())

	//路由根目录
	CollectRoute(r)
	return r
}

func CollectRoute(r *gin.Engine) {
	apiVersion := r.Group(config.API_V1_GROUP)

	common := apiVersion.Group(config.COMMON)
	commonRouter(common)

	pool := apiVersion.Group(config.POOL)
	demoRouter(pool)
}

func commonRouter(r *gin.RouterGroup) {
	r.GET("/getTimeStamp", GetTimeStamp)
}

func demoRouter(r *gin.RouterGroup) {
	r.POST("/getPicture", GetPictures)
}

/**
解决跨域
*/
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

/**
https映射
*/
func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:8090",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			return
		}
		c.Next()
	}
}

func ReqLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		//请求方式
		method := c.Request.Method
		//请求路由
		reqUrl := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()
		//请求参数
		get, _ := c.Get("req")
		marshal, _ := json.Marshal(get)
		// 打印日志
		logrus.WithFields(logrus.Fields{
			"status_code": statusCode,
			"client_ip":   clientIP,
			"req_method":  method,
			"req_uri":     reqUrl,
			"body":        string(marshal),
		}).Info()
	}
}
