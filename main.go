package main

import (
	"gin/core"
	"gin/global"
	"gin/utils"
)
func main() {
	utils.Init()
	db,_ := global.HS_DB.DB()
	defer db.Close()
	core.RunServer()
}
