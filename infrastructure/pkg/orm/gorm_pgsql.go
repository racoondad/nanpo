package orm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ensurePgsql(conf IGormConfig) error {
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

func gormPgsqlByConfig(conf IGormConfig) *gorm.DB {
	if conf.DbnameParam() == "" {
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
		sqlDB.SetMaxIdleConns(conf.MaxIdleConnsParam())
		sqlDB.SetMaxOpenConns(conf.MaxOpenConnsParam())
		return db
	}
}
