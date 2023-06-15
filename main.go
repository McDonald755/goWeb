package main

import (
	"goWeb/config"
	"goWeb/log"
	"goWeb/router"
	"net/http"
	"time"
)

func main() {
	log.ConfigLocalFilesystemLogger("./Log", "log", time.Hour*24*7, time.Hour*24)
	r := router.APIRouter()

	//自定义端口号，默认为8080
	r.Run(config.PORT)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           config.PORT,
		Handler:        r,
		ReadTimeout:    config.ReadTimeout,
		WriteTimeout:   config.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	if config.ENV == "debug" {
		server.ListenAndServe()
	} else if config.ENV == "prod" {
		//如果为https域名下运行则使用RunTLS
		r.RunTLS(config.PORT, config.PEM, config.KEY)
	}

}
