package global

import (
	"gin/config"
	"github.com/go-redis/redis"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	HS_DB    *gorm.DB
	CONFIG   config.Ymal
	HS_LOG   *zap.Logger
	HS_REDIS *redis.Client
)
