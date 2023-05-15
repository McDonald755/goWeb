package main

import (
	"goWeb/config"
	"goWeb/log"
	"goWeb/router"
	"time"
)

func main() {
	log.ConfigLocalFilesystemLogger("./Log", "log", time.Hour*24*7, time.Hour*24)
	r := router.APIRouter()
	//自定义端口号，若不制定则默认为8080
	r.Run(config.PORT)

	//如果为https域名下运行则使用RunTLS
	//r.RunTLS(config.PORT, config.PEM, config.KEY)
}
