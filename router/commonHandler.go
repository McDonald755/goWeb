package router

import (
	"github.com/gin-gonic/gin"
	"goWeb/config"
	"goWeb/utils"
	"net/http"
	"path/filepath"

	"time"
)

func GetTimeStamp(c *gin.Context) {
	unix := time.Now().Unix()
	defer c.Set("req", unix)
	c.JSON(http.StatusOK, utils.GetResponseVo(config.SUCCESS, config.GetMsg(config.SUCCESS), unix))
}

func UploadFile(c *gin.Context) {
	// 应该对文件名做具体限制 参照https://github.com/gin-gonic/gin/issues/1693
	file, _ := c.FormFile("file")
	filename := filepath.Base(file.Filename)
	c.SaveUploadedFile(file, filename)

	c.JSON(http.StatusOK, utils.GetResponseVo(config.SUCCESS, config.GetMsg(config.SUCCESS), filename))

}
func UploadFiles(c *gin.Context) {
	// 应该对文件名做具体限制 参照https://github.com/gin-gonic/gin/issues/1693
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	for _, file := range files {
		filename := filepath.Base(file.Filename)
		c.SaveUploadedFile(file, filename)
	}

	c.JSON(http.StatusOK, utils.GetResponseVo(config.SUCCESS, config.GetMsg(config.SUCCESS), nil))

}
