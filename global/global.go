package global

import (
	"sync"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/wangrui19970405/wu-shi-admin/server/utils/timer"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/wangrui19970405/wu-shi-admin/server/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	WUSHI_DB     *gorm.DB
	WUSHI_DBList map[string]*gorm.DB
	WUSHI_REDIS  *redis.Client
	WUSHI_CONFIG config.Server
	WUSHI_VP     *viper.Viper
	// WUSHI_LOG    *oplogging.Logger
	WUSHI_LOG                 *zap.Logger
	WUSHI_Timer               timer.Timer = timer.NewTimerTask()
	WUSHI_Concurrency_Control             = &singleflight.Group{}

	BlackCache local_cache.Cache
	lock       sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return WUSHI_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := WUSHI_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
