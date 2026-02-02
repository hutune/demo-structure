package common

import (
	"mtsgn-access-user-svc/pkg/config"
	"mtsgn-access-user-svc/pkg/redis"

	"gitea.solu-m.io/smart-pos/sp-system-common-svc/common/logger"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type App struct {
	Cfg    *config.Config
	DB     *gorm.DB
	Redis  *redis.RedisClient
	Minio  *minio.Client
	Logger *logger.ZeroLogger
}
