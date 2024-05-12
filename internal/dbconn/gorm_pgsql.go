package dbconn

import (
	"github.com/racoondad/nanpo/internal/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ensurePgsql(conf configs.SpecializedDB) error {
	pgsqlConfig := postgres.Config{
		DSN:                  conf.MainDsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig)); err != nil {
		return err
	} else {
		return db.Exec(conf.CreateProvider()).Error
	}
}

func gormPgsqlByConfig(conf configs.SpecializedDB) *gorm.DB {
	if conf.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  conf.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig), gormImpl.Config(conf)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
		sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
		return db
	}
}
