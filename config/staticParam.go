package config

const (
	// API root Group
	API = "/api"

	// V1 v1 Group
	V1           = "/v1"
	API_V1_GROUP = API + V1

	// goWeb Group
	GOWEB = "/goWeb"

	// COMMON common
	COMMON = "/common"

	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 10001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 10002
	ERROR_AUTH_TOKEN               = 10003
	ERROR_AUTH                     = 10004
	ERROR_UPLOAD_FAIL              = 10005
	ERROR_DB                       = 10006

	// 保存图片失败
	ERROR_UPLOAD_SAVE_IMAGE_FAIL = 30001
	// 检查图片失败
	ERROR_UPLOAD_CHECK_IMAGE_FAIL = 30002
	// 校验图片错误，图片格式或大小有问题
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003
)

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	ERROR_UPLOAD_FAIL:              "上传文件失败",
	ERROR_DB:                       "数据库错误",
	// 保存图片失败
	ERROR_UPLOAD_SAVE_IMAGE_FAIL: "保存图片失败",
	// 检查图片失败
	ERROR_UPLOAD_CHECK_IMAGE_FAIL: "检查图片失败",
	// 校验图片错误，图片格式或大小有问题
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
