package global

import (
	"gin/config"
	"gorm.io/gorm"
)

var (
	HS_DB   *gorm.DB
	GCONFIG config.Ymal
)
