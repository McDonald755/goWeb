package main

import (
	"goWeb/log"
	"goWeb/router"
	"time"
)

func main() {
	log.ConfigLocalFilesystemLogger("./Log", "log", time.Hour*24*7, time.Hour*24)
	r := router.APIRouter()
	//自定义端口号，若不制定则默认为8080
	r.Run(":9999")

	//如果为https域名下运行则使用RunTLS
	//r.RunTLS(":8090", config.PEM, config.KEY)
}
