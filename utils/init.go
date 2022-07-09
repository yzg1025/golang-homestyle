package utils

import (
	"gin/global"
)

func init() {
	ViperTool()
	global.HS_DB = GormMySql()
	global.HS_LOG = ZAP()
	global.HS_REDIS = InitRedis()
}
