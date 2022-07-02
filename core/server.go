package core

import (
	"fmt"
	"gin/global"
	"gin/initialization"
)

func RunServer() {
	var addr = fmt.Sprintf(":%d", global.CONFIG.Addr)
	hs := initialization.Routers()
	hs.Run(addr)
}
