package router

import (
	"github.com/gin-gonic/gin"
)

//func GetPictures(c *gin.Context) {
//	status := c.Query("status")
//	t := c.Query("type")
//	if t == "" || status == "" {
//		c.JSON(http.StatusOK, utils.GetResponseVo(config.INVALID_PARAMS, config.GetMsg(config.INVALID_PARAMS), nil))
//		return
//	}
//	defer c.Set("req", map[string]interface{}{"type": t, "status": status})
//	result := db.GetPictures(t, status)
//	c.JSON(http.StatusOK, utils.GetResponseVo(config.SUCCESS, config.GetMsg(config.SUCCESS), result))
//}

func SaveOrUpdateGame(c *gin.Context) {

	/*
		获取参数
		query 不带默认值
		defaultQuery 带默认值
	*/

}

func DeleteGame(c *gin.Context) {

}

func GetGameDetail(c *gin.Context) {

}

func GetGames(c *gin.Context) {

}
