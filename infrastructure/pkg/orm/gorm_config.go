package orm

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/schema"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type IGormConfig interface {
	MainDsn() string
	Dsn() string
	CreateProvider() string
	DbTypeParam() string
	HostParam() string
	PortParam() uint
	ConfigParam() string
	DbnameParam() string
	UsernameParam() string
	PasswordParam() string
	PrefixParam() string
	SingularParam() bool
	EngineParam() string
	MaxIdleConnsParam() int
	MaxOpenConnsParam() int
	LogModeParam() string
	LogZapParam() bool
}

var gormImpl = new(_gorm)

type _gorm struct{}

func (g *_gorm) Config(conf IGormConfig) *gorm.Config {
	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.PrefixParam(),
			SingularTable: conf.SingularParam(),
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	_default := logger.New(newWriter(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
	), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Warn,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	})

	switch conf.LogModeParam() {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}
	return config
}

type writer struct {
	logger.Writer
}

func newWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}
