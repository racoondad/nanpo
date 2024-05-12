package dbconn

import (
	"log"
	"os"
	"time"

	"github.com/racoondad/nanpo/internal/configs"

	"gorm.io/gorm/schema"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormImpl = new(_gorm)

type _gorm struct{}

func (g *_gorm) Config(conf configs.SpecializedDB) *gorm.Config {
	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.Prefix,
			SingularTable: conf.Singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	_default := logger.New(newWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})

	switch conf.LogMode {
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
