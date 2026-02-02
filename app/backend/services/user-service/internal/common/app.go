package common

import (
	"rmn-backend/services/user-service/pkg/config"
	"rmn-backend/services/user-service/pkg/redis"

	"rmn-backend/pkg/logger"
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
