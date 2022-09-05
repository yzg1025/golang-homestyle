package core

import (
	"fmt"
	"gin/global"
	"gin/initialization"
	"net/http"
	"time"
)

func RunServer() {
	var addr = fmt.Sprintf(":%d", global.CONFIG.Addr)
	hs := initialization.Routers()
	service := http.Server{
		Addr:         addr,
		Handler:      hs,
		ReadTimeout:  time.Second * 60,
		WriteTimeout: time.Second * 60,
	}
	service.ListenAndServe()
}
