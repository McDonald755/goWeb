package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"goWeb/config"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

func GetImageFullUrl(name string) string {
	return config.ImagePrefixUrl + "/" + GetImagePath() + name
}

func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = EncodeMD5(fileName)

	return fileName + ext
}

func GetImagePath() string {
	return config.ImageSavePath
}

func GetImageFullPath() string {
	return config.ImageSavePath
}

func CheckImageExt(fileName string) bool {
	ext := GetExt(fileName)
	for _, allowExit := range strings.Split(config.ImageAllowExits, ",") {
		if strings.ToUpper(allowExit) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

func CheckImageSize(f multipart.File) bool {
	size, err := GetSize(f)
	if err != nil {
		log.Error(err)
		return false
	}

	return size <= config.ImageMaxSize
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
