package router

import (
	"github.com/gin-gonic/gin"
	"goWeb/config"
	"goWeb/utils"
	"net/http"

	"time"
)

func GetTimeStamp(c *gin.Context) {
	unix := time.Now().Unix()
	defer c.Set("req", unix)
	c.JSON(http.StatusOK, utils.GetResponseVo(config.SUCCESS, config.GetMsg(config.SUCCESS), unix))
}
