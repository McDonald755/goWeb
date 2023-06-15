package utils

import (
	"crypto/md5"
	"encoding/hex"
	"goWeb/vo"
	"time"
)

func GetResponsePageVo(code int, message string, data *vo.ResponsePageVo) *vo.ResponsePageVo {
	data.Code = code
	data.Message = message
	return data
}

func GetResponseVo(code int, message string, data interface{}) *vo.ResponseVo {
	responseVo := vo.ResponseVo{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return &responseVo
}

func TimeStampToTime(tm int64) time.Time {
	tm = tm / 1000
	timeFormat := time.Unix(tm, 0).Format("2006-01-02 15:04:05")
	duration, _ := time.ParseInLocation("2006-01-02 15:04:05", timeFormat, time.Local)
	return duration
}

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
