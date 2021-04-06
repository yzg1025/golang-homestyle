package utils

import "gin/global"

func Init()  {
	ViperTool()
	global.HS_DB = GormMySql()
	global.HS_LOG = ZAP()
}
