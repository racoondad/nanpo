package nanpo

import (
	"github.com/racoondad/nanpo/infrastructure/pkg/configs"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	VP *viper.Viper
	// 配置
	CONFIG configs.Server
	// 数据库
	DB *gorm.DB
	// 缓存
	RD *redis.Client
	// 日志
	LOG *zap.Logger
)

type Application interface {
	InstallConfig()     // 初始化配置
	InstallLogging()    // 初始化日志
	InstallDatabase()   // 初始化数据库
	InstallRouter()     // 初始化路由
	InstallEvent()      // 初始化消息中间件
	InstallMqttServer() // 初始化MQTT服务
	InstallServer()     // 初始化主服务
}
