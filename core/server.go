package core

import (
	"fmt"
	"gin/global"
	"gin/initialization"
)

func RunServer()  {
	var addr = fmt.Sprintf(":%d",global.GCONFIG.Addr)
	hs := initialization.Routers()
	hs.Run(addr)
}
