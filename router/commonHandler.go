package router

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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

func UploadImage(c *gin.Context) {
	code := config.SUCCESS
	data := make(map[string]string)

	// 应该对文件名做具体限制 参照https://github.com/gin-gonic/gin/issues/1693
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, utils.GetResponseVo(code, config.GetMsg(code), data))
	}

	if image == nil {

	} else {
		imageName := utils.GetImageName(image.Filename)
		fullPath := utils.GetImageFullPath()
		savePath := utils.GetImagePath()

		src := fullPath + imageName
		if !utils.CheckImageExt(imageName) || !utils.CheckImageSize(file) {
			code = config.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			err := utils.CheckImage(fullPath)
			if err != nil {
				log.Warn(err)
				code = config.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				log.Warn(err)
				code = config.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				data["image_url"] = utils.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}
	}

	c.JSON(http.StatusOK, utils.GetResponseVo(code, config.GetMsg(code), data))

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
