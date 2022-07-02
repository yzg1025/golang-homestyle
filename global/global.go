package global

import (
	"gin/config"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	HS_DB  *gorm.DB
	CONFIG config.Ymal
	HS_LOG *zap.Logger
)
