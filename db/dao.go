package db

import (
	log "github.com/sirupsen/logrus"
	"goWeb/config"
)

func GetPictures(t, status string) *[]PICTURE {
	var result []PICTURE
	res := config.DB.Table("PICTURE").Where("type = ? and status = ? and deleted = 0", t, status).Find(&result)
	if res.Error != nil {
		log.Error(res.Error)
	}
	return &result
}
