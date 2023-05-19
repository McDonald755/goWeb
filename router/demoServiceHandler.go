package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"goWeb/config"
	"goWeb/service"
	"goWeb/utils"
	"goWeb/vo"
	"net/http"
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
	validate := validator.New()
	gameVo := vo.GameVo{}
	//参数校验
	c.ShouldBind(&gameVo)
	if err := validate.Struct(gameVo); err != nil {
		c.JSON(http.StatusOK, utils.GetResponseVo(config.ERROR_DB, config.GetMsg(config.ERROR_DB), nil))
		return
	}
	err, res := service.SaveOrUpdateGame(gameVo)

	if err == nil {
		c.JSON(http.StatusOK, utils.GetResponseVo(config.ERROR_DB, config.GetMsg(config.ERROR_DB), nil))
		return
	}

	c.JSON(http.StatusOK, utils.GetResponseVo(config.SUCCESS, config.GetMsg(config.SUCCESS), res))

}

func DeleteGame(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, utils.GetResponseVo(config.INVALID_PARAMS, config.GetMsg(config.INVALID_PARAMS), nil))
	}

	res := service.DeleteGame(id)
	if res == nil {
		c.JSON(http.StatusOK, utils.GetResponseVo(config.ERROR_DB, config.GetMsg(config.ERROR_DB), nil))
		return
	}

	c.JSON(http.StatusOK, utils.GetResponseVo(config.SUCCESS, config.GetMsg(config.SUCCESS), res))
}

func GetGameDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, utils.GetResponseVo(config.INVALID_PARAMS, config.GetMsg(config.INVALID_PARAMS), nil))
		return
	}

	res := service.GetGameDetail(id)
	if res == nil {
		c.JSON(http.StatusOK, utils.GetResponseVo(config.ERROR_DB, config.GetMsg(config.ERROR_DB), nil))
		return
	}

	c.JSON(http.StatusOK, utils.GetResponseVo(config.SUCCESS, config.GetMsg(config.SUCCESS), res))
}

func GetGames(c *gin.Context) {
	validate := validator.New()
	gameVo := vo.GameVo{}
	//参数校验
	c.ShouldBind(&gameVo)
	if err := validate.Struct(gameVo); err != nil {
		c.JSON(http.StatusOK, utils.GetResponseVo(config.ERROR_DB, config.GetMsg(config.ERROR_DB), nil))
		return
	}
	res := service.GetGames(gameVo)

	if res == nil {
		c.JSON(http.StatusOK, utils.GetResponseVo(config.ERROR_DB, config.GetMsg(config.ERROR_DB), nil))
		return
	}

	c.JSON(http.StatusOK, utils.GetResponseVo(config.SUCCESS, config.GetMsg(config.SUCCESS), res))
}
